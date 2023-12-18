package gotools

import (
	"sync"
)

// Future executes a function f in background. It returns
// a function that will wait for the result and returns it.
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

// FutureWithOk (just like [Future] – but with additional ok output)
// executes a function f in background. It returns a function that will
// wait for the result and returns it.
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
