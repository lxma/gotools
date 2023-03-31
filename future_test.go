package gotools

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFuture(t *testing.T) {
	start := time.Now()
	fut1 := Future(func() int {
		time.Sleep(100 * time.Millisecond)
		return 1
	})
	fut2 := Future(func() int {
		time.Sleep(100 * time.Millisecond)
		return 2
	})
	value1 := fut1()
	value2 := fut2()
	duration := time.Since(start)
	assert.Equal(t, 1, value1, "value 1 is correct")
	assert.Equal(t, 2, value2, "value 2 is correct")
	assert.Less(t, 90*time.Millisecond, duration, "needed to wait one execution time")
	assert.Greater(t, 190*time.Millisecond, duration, "did not need to wait full execution time")
}

func TestFutureWithError(t *testing.T) {
	var v1 int
	start := time.Now()
	fut1 := FutureWithError(func() (int, error) {
		time.Sleep(100 * time.Millisecond)
		return 2, nil
	})
	fut2 := FutureWithError(func() (string, error) {
		time.Sleep(100 * time.Millisecond)
		return "test", nil
	})
	fut3 := FutureWithError(func() ([]int, error) {
		time.Sleep(100 * time.Millisecond)
		return []int{2, 3}, nil
	})
	fut4 := FutureWithError(func() ([]int, error) {
		time.Sleep(100 * time.Millisecond)
		return nil, fmt.Errorf("testerror")
	})
	v1, err1 := fut1()
	v2, err2 := fut2()
	v3, err3 := fut3()
	_, err4 := fut4()
	duration := time.Since(start)
	assert.Less(t, 90*time.Millisecond, duration, "needed to wait one execution time")
	assert.Greater(t, 190*time.Millisecond, duration, "did not need to wait full execution time")

	assert.Equal(t, 2, v1)
	assert.Equal(t, nil, err1)
	assert.Equal(t, "test", v2)
	assert.Equal(t, nil, err2)
	assert.Equal(t, []int{2, 3}, v3)
	assert.Equal(t, nil, err3)
	assert.Equal(t, "testerror", err4.Error())
}

func TestFutureWithOk(t *testing.T) {
	var v1 int
	start := time.Now()
	fut1 := FutureWithOk(func() (int, bool) {
		time.Sleep(100 * time.Millisecond)
		return 2, true
	})
	fut2 := FutureWithOk(func() (string, bool) {
		time.Sleep(100 * time.Millisecond)
		return "test", true
	})
	fut3 := FutureWithOk(func() ([]int, bool) {
		time.Sleep(100 * time.Millisecond)
		return []int{2, 3}, true
	})
	fut4 := FutureWithOk(func() ([]int, bool) {
		time.Sleep(100 * time.Millisecond)
		return nil, false
	})
	v1, ok1 := fut1()
	v2, ok2 := fut2()
	v3, ok3 := fut3()
	_, ok4 := fut4()
	duration := time.Since(start)
	assert.Less(t, 90*time.Millisecond, duration, "needed to wait one execution time")
	assert.Greater(t, 190*time.Millisecond, duration, "did not need to wait full execution time")

	assert.Equal(t, 2, v1)
	assert.True(t, ok1)
	assert.Equal(t, "test", v2)
	assert.True(t, ok2)
	assert.Equal(t, []int{2, 3}, v3)
	assert.True(t, ok3)
	assert.False(t, ok4)
}
