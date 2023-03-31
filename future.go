package gotools

import "sync"

func Future[A any](f func() A) func() A {
	ch := make(chan A)
	m := sync.Mutex{}
	haveResult := false
	var bufferedResult A
	go func() {
		ch <- f()
		close(ch)
	}()
	return func() A {
		m.Lock()
		defer m.Unlock()
		if !haveResult {
			bufferedResult = <-ch
			haveResult = true
		}
		return bufferedResult
	}
}

type futureBufferWithError struct {
	value any
	error error
}

type futureBufferWithOk struct {
	value any
	ok    bool
}

func FutureWithError[A any](f func() (A, error)) func() (A, error) {
	ch := make(chan futureBufferWithError)
	m := sync.Mutex{}
	haveResult := false
	var bufferedResult futureBufferWithError
	go func() {
		val, err := f()
		ch <- futureBufferWithError{val, err}
		close(ch)
	}()
	return func() (A, error) {
		m.Lock()
		defer m.Unlock()
		if !haveResult {
			bufferedResult = <-ch
			haveResult = true
		}
		return bufferedResult.value.(A), bufferedResult.error
	}
}

func FutureWithOk[A any](f func() (A, bool)) func() (A, bool) {
	ch := make(chan futureBufferWithOk)
	m := sync.Mutex{}
	haveResult := false
	var bufferedResult futureBufferWithOk
	go func() {
		val, ok := f()
		ch <- futureBufferWithOk{val, ok}
		close(ch)
	}()
	return func() (A, bool) {
		m.Lock()
		defer m.Unlock()
		if !haveResult {
			bufferedResult = <-ch
			haveResult = true
		}
		return bufferedResult.value.(A), bufferedResult.ok
	}
}
