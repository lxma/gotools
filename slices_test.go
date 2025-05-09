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

func TestRepeatedly(t *testing.T) {
	r1 := Repeatedly(2, func() []int {
		return []int{1}
	})
	r2 := Repeat([]int{1}, 2)
	assert.Equal(t, [][]int{{1}, {1}}, r1)
	assert.Equal(t, [][]int{{1}, {1}}, r2)
	r1[0][0] = 2
	r2[0][0] = 2
	assert.Equal(t, [][]int{{2}, {1}}, r1)
	assert.Equal(t, [][]int{{2}, {2}}, r2)
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
	assert.Equal(t, [][]int{}, Transpose([][]int{}))
	assert.Equal(t, [][]int{}, Transpose([][]int{{}, {}}))
}

func TestReduce(t *testing.T) {
	assert.Equal(t, 10, Reduce(func(a int, b int) int { return a + b }, 1, []int{2, 3, 4}))
}

func TestConcat(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 2, 3, 8}, Concat([]int{1, 2, 3}, []int{}, []int{2, 3}, []int{8}))
	emptySlcOfSlces := [][]int{}
	assert.Equal(t, []int{}, Concat(emptySlcOfSlces...))
}

func less(a, b int) bool { return a < b }

func TestSort(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, SortBy([]int{3, 1, 2}, less))
	assert.Equal(t, []int{}, SortBy([]int{}, less))
	assert.Equal(t, []int{1, 2, 3}, SortStable([]int{3, 1, 2}, less))
	assert.Equal(t, []int{}, SortStable([]int{}, less))
}

func TestEvery(t *testing.T) {
	assert.Equal(t, true, Every([]int{2, 3, 4}, func(n int) bool { return n > 1 }))
	assert.Equal(t, false, Every([]int{2, 0, 4}, func(n int) bool { return n > 1 }))
	assert.Equal(t, true, Every([]int{}, func(n int) bool { return n > 1 }))
}

func TestSome(t *testing.T) {
	assert.Equal(t, false, Some([]int{2, 3, 4}, func(n int) bool { return n < 1 }))
	assert.Equal(t, true, Some([]int{2, 0, 4}, func(n int) bool { return n < 1 }))
	assert.Equal(t, false, Some([]int{}, func(n int) bool { return n < 1 }))
}

func TestPartitionAt(t *testing.T) {
	assert.Equal(t, [][]int{{1, 2}, {3, 4}, {5}}, PartitionAt([]int{1, 2, 0, 3, 4, 0, 5}, 0))
	assert.Equal(t, [][]int{{1, 2, 3}}, PartitionAt([]int{1, 2, 3}, 0))
	assert.Equal(t, [][]int{}, PartitionAt([]int{}, 0))
}

func TestPartitionBy(t *testing.T) {
	assert.Equal(t, [][]int{{1}, {2, 3}, {4, 5}},
		PartitionBy([]int{1, 2, 3, 4, 5}, func(n int) int {
			return n / 2
		}))
	assert.Equal(t, [][]int{{1, 2, 3}},
		PartitionBy([]int{1, 2, 3}, func(n int) bool {
			return true
		}))
	assert.Equal(t, [][]int{},
		PartitionBy([]int{}, func(n int) bool {
			return true
		}))
	assert.Equal(t, [][]int{{1, 1}, {2}, {1}, {3, 3, 3}},
		PartitionBy([]int{1, 1, 2, 1, 3, 3, 3}, Identity[int]))
}

func TestIndexOf(t *testing.T) {
	val, ok := IndexOf([]int{1, 2, 3}, 1)
	assert.Equal(t, 0, val)
	assert.Equal(t, true, ok)
	val, ok = IndexOf([]int{1, 2, 3}, 3)
	assert.Equal(t, 2, val)
	assert.Equal(t, true, ok)
	val, ok = IndexOf([]int{1, 2, 3}, 4)
	assert.Equal(t, -1, val)
	assert.Equal(t, false, ok)
	val, ok = IndexOf([]int{}, 4)
	assert.Equal(t, -1, val)
	assert.Equal(t, false, ok)
}
