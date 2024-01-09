package gotools

import (
	"sync"
)

// Future executes a function f in background. It returns
// a function that will wait for the result and returns it.
//
// The return value will be buffered, so the function can be
// called multiple times. I.e. it can be treated as a regular
// result and be passed around without worrying who calls it
// first.
// Example:
//
//	futureValue1 := Future(func() int {
//	    time.Sleep(100 * time.Millisecond)
//	    return 1
//	})
//	futureValue2 := Future(func() int {
//	    time.Sleep(100 * time.Millisecond)
//	    return 2
//	})
//	fmt.Printf("Values: %d, %d\n", futureValue1(), futureValue2())
//
// will print "Values: 1, 2" and it will run for little more than 100ms.
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

// FutureWithError (just like [Future] – but with additional error output)
// executes a function f in background. It returns a function that will
// wait for the result and returns it.
//
// The return value will be buffered, so the function can be
// called multiple times. I.e. it can be treated as a regular
// result and be passed around without worrying who calls it
// first.
func FutureWithError[A any](f func() (A, error)) func() (A, error) {
	c := make(chan futureBufferWithError)
	m := sync.Mutex{}
	haveResult := false
	var bufferedResult futureBufferWithError
	go func() {
		val, err := f()
		c <- futureBufferWithError{val, err}
		close(c)
	}()
	return func() (A, error) {
		m.Lock()
		defer m.Unlock()
		if !haveResult {
			bufferedResult = <-c
			haveResult = true
		}
		return bufferedResult.value.(A), bufferedResult.error
	}
}

// FutureWithOk (just like [Future] – but with additional ok output)
// executes a function f in background. It returns a function that will
// wait for the result and returns it.
//
// The return value will be buffered, so the function can be
// called multiple times. I.e. it can be treated as a regular
// result and be passed around without worrying who calls it
// first.
func FutureWithOk[A any](f func() (A, bool)) func() (A, bool) {
	c := make(chan futureBufferWithOk)
	m := sync.Mutex{}
	haveResult := false
	var bufferedResult futureBufferWithOk
	go func() {
		val, ok := f()
		c <- futureBufferWithOk{val, ok}
		close(c)
	}()
	return func() (A, bool) {
		m.Lock()
		defer m.Unlock()
		if !haveResult {
			bufferedResult = <-c
			haveResult = true
		}
		return bufferedResult.value.(A), bufferedResult.ok
	}
}
