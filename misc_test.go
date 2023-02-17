package gotools

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetIntegersInString(t *testing.T) {
	assert.Equal(t, []int{1, -2, 3}, GetIntegersInString("Und 1 und -2, ,+3"))
}

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
	value1 := <-fut1
	value2 := <-fut2
	duration := time.Since(start)
	assert.Equal(t, 1, value1, "value 1 is correct")
	assert.Equal(t, 2, value2, "value 2 is correct")
	assert.Less(t, 90*time.Millisecond, duration, "needed to wait one execution time")
	assert.Greater(t, 190*time.Millisecond, duration, "did not need to wait full execution time")
}
