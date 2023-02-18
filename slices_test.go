package gotools

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntSequence(t *testing.T) {
	assert.Equal(t, []int{0, 1, 2}, IntSequence(3))
	assert.Equal(t, []int{1, 2, 3, 4}, IntSequence(4, 1))
}

func TestRepeat(t *testing.T) {
	assert.Equal(t, []int{4, 4}, Repeat(4, 2))
	assert.Equal(t, []string{"a", "a", "a"}, Repeat("a", 3))
}

func TestFilter(t *testing.T) {
	assert.Equal(t, []int{2, 4}, Filter([]int{1, 2, 3, 4}, func(n int) bool { return n%2 == 0 }))
}

func TestRemove(t *testing.T) {
	assert.Equal(t, []int{1, 3}, Remove([]int{1, 2, 3, 4}, func(n int) bool { return n%2 == 0 }))
}

func TestTakeWhile(t *testing.T) {
	assert.Equal(t, []int{1, 2}, TakeWhile([]int{1, 2, 3, 4}, func(n int) bool { return n < 3 }))
	assert.Equal(t, []int{}, TakeWhile([]int{1, 2, 3, 4}, func(n int) bool { return n > 10 }))
	assert.Equal(t, []int{}, TakeWhile([]int{}, func(n int) bool { return n < 3 }))
}

func TestDropWhile(t *testing.T) {
	assert.Equal(t, []int{3, 4}, DropWhile([]int{1, 2, 3, 4}, func(n int) bool { return n < 3 }))
	assert.Equal(t, []int{}, DropWhile([]int{}, func(n int) bool { return n < 3 }))
	assert.Equal(t, []int{}, DropWhile([]int{1, 2, 3, 4}, func(n int) bool { return n < 10 }))
}

func TestMap(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, Map([]string{"1", "2", "3"}, StringToInt))
	assert.Equal(t, []int{}, Map([]string{}, StringToInt), "empty slice")
}

func TestReverse(t *testing.T) {
	assert.Equal(t, []int{3, 2, 1}, Reverse([]int{1, 2, 3}))
	assert.Equal(t, []int{}, Reverse([]int{}), "empty reverse")
}

func TestCopySlice(t *testing.T) {
	arr := [3]int{1, 2, 3}
	slc1 := arr[0:2]
	slc2 := CopySlice(slc1)
	slc1 = append(slc1, 4)
	slc2 = append(slc2, 5)
	assert.Equal(t, []int{1, 2, 4}, slc1, "straight forward copy – changes at slc2 don't affect slc1")
	assert.Equal(t, []int{1, 2, 5}, slc2, "straight forward copy – changes at slc1 don't affect slc2")
	assert.Equal(t, []int{}, CopySlice([]int{}), "empty copy")
}

func TestSum(t *testing.T) {
	assert.Equal(t, 14, Sum([]int{2, 2, 5, 5}), "sum of ints")
	assert.InDelta(t, 14.3, Sum([]float64{2, 2, 5.2, 5.1}), 0.000001, "sum of floats")
	assert.Equal(t, 0, Sum([]int{}), "empty sum")
}

func TestProd(t *testing.T) {
	assert.Equal(t, 100, Prod([]int{2, 2, 5, 5}), "product of ints")
	assert.InDelta(t, 106.08, Prod([]float64{2, 2, 5.1, 5.2}), 0.000001, "product of floats")
	assert.Equal(t, 1, Prod([]int{}), "empty sum")
}

func TestTranspose(t *testing.T) {
	assert.Equal(t,
		[][]int{
			{1, 4},
			{2, 5},
			{3, 6},
		},
		Transpose([][]int{
			{1, 2, 3},
			{4, 5, 6},
		}),
		"simple transpose",
	)
}
