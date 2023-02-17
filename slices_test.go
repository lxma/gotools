package gotools

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, Map([]string{"1", "2", "3"}, StringToInt), "straight forward test")
	assert.Equal(t, []int{}, Map([]string{}, StringToInt), "empty slice")
}

func TestReverse(t *testing.T) {
	assert.Equal(t, []int{3, 2, 1}, Reverse([]int{1, 2, 3}), "straight forward reverse")
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
	assert.Equal(t, 14, Sum([]int{2, 2, 5, 5}), "simple sum")
	assert.Equal(t, 14.0, Sum([]float64{2, 2, 5, 5}), "simple sum")
	assert.Equal(t, 0, Sum([]int{}), "empty sum")
}

func TestProd(t *testing.T) {
	assert.Equal(t, 100, Prod([]int{2, 2, 5, 5}), "simple sum")
	assert.Equal(t, 100.0, Prod([]float64{2, 2, 5, 5}), "simple sum")
	assert.Equal(t, 1, Prod([]int{}), "empty sum")
}

func TestTranspose(t *testing.T) {
	assert.Equal(t,
		[][]int{
			[]int{1, 4},
			[]int{2, 5},
			[]int{3, 6},
		},
		Transpose([][]int{
			[]int{1, 2, 3},
			[]int{4, 5, 6},
		}),
		"simple transpose",
	)
}
