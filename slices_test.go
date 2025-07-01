package gotools

import (
    "cmp"
    "github.com/stretchr/testify/assert"
    "reflect"
    "slices"
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

func TestRange(t *testing.T) {
    assert.Equal(t, []int{0, 1, 2}, Range(3))
    assert.Equal(t, []int{}, Range(0))
    assert.Panics(t, func() { Range(-1) })
}

func TestRangeFromTo(t *testing.T) {
    assert.Equal(t, []int{2, 3, 4}, RangeFromTo(2, 5))
    assert.Equal(t, []int{}, RangeFromTo(0, 0))
    assert.Panics(t, func() { RangeFromTo(2, 1) }, "To must be greater or equal than from")
    assert.Equal(t, []int{-2, -1, 0, 1, 2}, RangeFromTo(-2, 3))
}

func CheckRandomized[T cmp.Ordered](t *testing.T, slice []T, msg string) {
    changedOrderOrSmall := len(slice) < 2
    for i := 0; i < 10000; i++ {
        randomized := Randomize(slice)
        if !reflect.DeepEqual(randomized, slice) {
            changedOrderOrSmall = true
        }
        slices.Sort(randomized)
        assert.Equal(t, randomized, slice, msg)
        if changedOrderOrSmall {
            break
        }
    }
    assert.True(t, changedOrderOrSmall, "At some point, the order must be changed")
}

func TestRandomize(t *testing.T) {
    originalSlice1 := []int{1, 2}
    matchOriginal := false
    matchOpposite := false
    for i := 0; i < 10000; i++ {
        randomSlice := Randomize(originalSlice1)
        if reflect.DeepEqual(originalSlice1, randomSlice) {
            matchOriginal = true
        } else if reflect.DeepEqual([]int{2, 1}, randomSlice) {
            matchOpposite = true
        } else {
            assert.Fail(t, "Randomize should retain the original values in /some/ order")
        }
        if matchOriginal && matchOpposite {
            break
        }
    }
    assert.True(t, matchOriginal, "Sometimes the original order should be retained")
    assert.True(t, matchOpposite, "Sometimes the order should be changed")

    CheckRandomized(t, []int{}, "Empty slice must be randomizable")
    CheckRandomized(t, []string{"x"}, "Empty slice must be randomizable")
    CheckRandomized(t, []int{1, 2, 2, 3, 4}, "Multiple identical values should be possible")
    CheckRandomized(t, []string{"a", "b", "b", "cd", "e"}, "Strings should be possible")

    counts := [6]int{0, 0, 0, 0, 0, 0}
    expectedCounts := 5000
    for i := 0; i < expectedCounts*6; i++ {
        rnd := Randomize([]int{1, 2, 3})
        if reflect.DeepEqual(rnd, []int{1, 2, 3}) {
            counts[0]++
        } else if reflect.DeepEqual(rnd, []int{1, 3, 2}) {
            counts[1]++
        } else if reflect.DeepEqual(rnd, []int{2, 1, 3}) {
            counts[2]++
        } else if reflect.DeepEqual(rnd, []int{2, 3, 1}) {
            counts[3]++
        } else if reflect.DeepEqual(rnd, []int{3, 1, 2}) {
            counts[4]++
        } else if reflect.DeepEqual(rnd, []int{3, 2, 1}) {
            counts[5]++
        } else {
            assert.Fail(t, "Randomize should retain the original values in /some/ order")
        }
    }
    for i := 0; i < 6; i++ {
        if Abs(expectedCounts-counts[i]) > expectedCounts/20 {
            assert.Fail(t, "", "Randomize should give all alternatives about equal probability (but alternative %d got %.1f%% instead of 16.7%%).", i, float64(counts[i]*100)/float64(expectedCounts*6))
        }
    }
}

func TestSortLexicographic(t *testing.T) {
    assert.Equal(t, [][]int{{1, 2}, {2, 1}}, SortedLex([][]int{{2, 1}, {1, 2}}))
    assert.Equal(t, [][]int{}, SortedLex([][]int{}), "Sort empty slice")
    assert.Equal(t, [][]int{{1}, {1, 2}}, SortedLex([][]int{{1, 2}, {1}}), "Sort slices of different length")
    assert.Equal(t, [][]int{{1}, {1, 2}, {2}}, SortedLex([][]int{{1}, {2}, {1, 2}}), "Sort slices of different length 2")
    assert.Equal(t, [][]string{{"a"}, {"a", "b"}}, SortedLex([][]string{{"a", "b"}, {"a"}}), "Sorting slices of slices of strings")
}

func TestPermutations(t *testing.T) {
    perms := Permutations([]int{1, 2, 3})
    SortLex(perms)
    assert.Equal(t, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}, perms)

    assert.Equal(t, [][]int{}, Permutations([]int{}), "Empty slice")
    assert.Equal(t, [][]int{{1}}, Permutations([]int{1}), "Empty slice")
}

func collectPermutations[T cmp.Ordered](slc []T) [][]T {
    allPerms := make([][]T, 0)
    for perm := range PermutationsIter(slc) {
        allPerms = append(allPerms, perm)
    }
    SortLex(allPerms)
    return allPerms
}

func TestPermutationsIter(t *testing.T) {
    assert.Equal(t, [][]int{{1, 2}, {2, 1}}, collectPermutations([]int{1, 2}))
    assert.Equal(t, [][]int{}, collectPermutations([]int{}))
    assert.Equal(t, [][]int{{1}}, collectPermutations([]int{1}))
}
