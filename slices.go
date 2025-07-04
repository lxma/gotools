package gotools

import (
	"cmp"
	"github.com/lxma/golist/v2"
	"iter"
	"math/big"
	"math/rand"
	"sort"
)

// TakeWhile returns the first elements of a slice for which a given function
// returns true. Example:
//
//	TakeWhile([]int{1, 2, 3, 4}, func(n int) bool { return n < 3 })
//
// returns
//
//	[]int{1, 2}
func TakeWhile[A any](slc []A, f func(A) bool) []A {
	for i, elt := range slc {
		if !f(elt) {
			return slc[0:i]
		}
	}
	return slc
}

// DropWhile returns the given slice except for the first elements of a slice
// for which a given function returns true. Example:
//
//	DropWhile([]int{1, 2, 3, 4}, func(n int) bool { return n < 3 })
//
// returns
//
//	[]int{3, 4}
func DropWhile[A any](slc []A, f func(A) bool) []A {
	for i, elt := range slc {
		if !f(elt) {
			return slc[i:]
		}
	}
	return []A{}
}

// IntSequence returns a sequence of consecutive integers of a given length
// starting with "from" (default: 0 )
// Examples:
//
//	IntSequence(3)
//
// returns []int{0, 1, 2};
//
//	IntSequence(3, 1)
//
// returns []int{1, 2, 3}
func IntSequence(length int, from ...int) []int {
	start := 0
	if len(from) > 0 {
		start = from[0]
	}
	seq := make([]int, length)
	for i := 0; i < length; i++ {
		seq[i] = start + i
	}
	return seq
}

// Repeat returns a sequence that repeats a given element n times.
// Example:
//
//	Repeat("s", 3)
//
// returns
//
//	[]string{"s", "s", "s"}
//
// Note: This works well only with value types (or constants – if this is intended).
// To produce individual reference values, use [Repeatedly]
func Repeat[A any](elt A, n int) []A {
	result := make([]A, n)
	for i := 0; i < n; i++ {
		result[i] = elt
	}
	return result
}

// Repeatedly (similar to [Repeat]) executes a function n times returning a slice
// of n times its result. Purpose is to permit the repeated
// construction of a value. Example: After
//
//	r1 := Repeatedly(2, func() []int {
//	  return []int{1}
//	})
//	r2 := Repeat([]int{1}, 2)
//
// r1 and r2 both look like [][]int{{1}, {1}}. But after
//
//	r1[0][0] = 2
//	r2[0][0] = 2
//
// r1[1][0] is 1, while r2[1][0] is 2, because r2[0] and r2[1] point to the same slice.
func Repeatedly[A any](n int, f func() A) []A {
	result := make([]A, n)
	for i := 0; i < n; i++ {
		result[i] = f()
	}
	return result
}

// Transpose does a matrix transpose
// Example:
//
//	Transpose([][]int{
//		{1, 2, 3},
//		{4, 5, 6},
//	})
//
// returns
//
//	[][]int{
//		{1, 4},
//		{2, 5},
//		{3, 6},
//	}
//
// In case either dimension is 0, an empty slice is returned.
// (A matrix with zero columns – transposed – will have zero
// rows. Thus, the result is an empty slice. A matrix with zero
// rows cannot be said to have non-zero columns. So, also in this
// case an empty slice is returned. This is due to the fact that
// the dimensions of the matrix are not held separately.)
func Transpose[C any](data [][]C) [][]C {
	if len(data) == 0 || len(data[0]) == 0 {
		return [][]C{}
	}
	newWidth := len(data)
	newHeight := len(data[0])
	result := make([][]C, newHeight)
	for newY := 0; newY < newHeight; newY++ {
		thisLine := make([]C, newWidth)
		for newX := 0; newX < newWidth; newX++ {
			thisLine[newX] = data[newX][newY]
		}
		result[newY] = thisLine
	}
	return result
}

// Map maps all values of a slice with a function f.
// Example:
//
//	Map([]string{"1", "2", "3"}, StringToInt)
//
// returns
//
//	[]int{1, 2, 3}
func Map[A any, B any](slc []A, f func(A) B) []B {
	result := make([]B, len(slc))
	for i, value := range slc {
		result[i] = f(value)
	}
	return result
}

// Filter takes a slice slc and a function f. It returns a new slice
// containing all values of that slice for which f returns true.
// Example:
//
//	Filter([]int{1, 2, 3, 4}, func(n int) bool { return n%2 == 0 })
//
// returns
//
//	[]int{2, 4}
func Filter[A any](slc []A, f func(A) bool) []A {
	keep := make([]bool, len(slc))
	count := 0
	for i := 0; i < len(slc); i++ {
		keep[i] = f(slc[i])
		if keep[i] {
			count++
		}
	}
	result := make([]A, count)
	newIdx := 0
	for i, elt := range slc {
		if keep[i] {
			result[newIdx] = elt
			newIdx++
		}
	}
	return result
}

// Remove takes a slice slc and a function f. It returns a new slice
// containing all values of that slice for which f returns true.
// Example:
//
//	Remove([]int{1, 2, 3, 4}, func(n int) bool { return n%2 == 0 })
//
// returns
//
//	[]int{1, 3}
func Remove[A any](slc []A, f func(A) bool) []A {
	keep := make([]bool, len(slc))
	count := 0
	for i := 0; i < len(slc); i++ {
		keep[i] = !f(slc[i])
		if keep[i] {
			count++
		}
	}
	result := make([]A, count)
	newIdx := 0
	for i, elt := range slc {
		if keep[i] {
			result[newIdx] = elt
			newIdx++
		}
	}
	return result
}

// Sum takes a slice of numbers. It sums all values of that slice. If the slice is empty,
// it returns 0.
// Examples:
//
//	Sum([]int{1,2,3})
//
// returns 6
//
//	Sum([]float64{1.5, 2.5, 3.5})
//
// returns 7.5
func Sum[N Number](slc []N) N {
	result := N(0)
	for _, val := range slc {
		result += val
	}
	return result
}

// Prod takes a slice of numbers. It returns the product of all values of that slice. If the slice is empty,
// it returns 1.
// Examples:
//
//	Prod([]int{1,2,3})
//
// returns 6
//
//	Prod([]float64{1.5, 2, 4})
//
// returns 12.0 (a float value)
func Prod[N Number](slc []N) N {
	result := N(1)
	for _, val := range slc {
		result *= val
	}
	return result
}

// Reverse reverses a slice
// Example:
//
//	Prod([]int{1,2,3})
//
// returns []int{3,2,1}
func Reverse[A any](slc []A) []A {
	reversed := make([]A, len(slc))
	for i, elt := range slc {
		reversed[len(slc)-i-1] = elt
	}
	return reversed
}

// CopySlice copies a slice (it does not copy the values, so it's not a deep copy)
func CopySlice[C any](slc []C) []C {
	newSlice := make([]C, len(slc))
	copy(newSlice, slc)
	return newSlice
}

// Reduce is a classic reduce as used in functional programming. It takes
// a binary function, a start value and a slice of values
// and calculates f(f(f(f(startValue, slc[0]), slc[1]), slc[2]) ...)
// Example:
//
//	func add(a int, b int) int { return a + b }
//	Reduce(add, 1, []int{2, 3, 4})
//
// returns 10
func Reduce[A any, B any](f func(A, B) A, startValue A, slc []B) A {
	result := startValue
	for _, v := range slc {
		result = f(result, v)
	}
	return result
}

// Concat concatenates slices
// Example:
//
//	Concat([]int{1,2,3}, []int{2,3}, []int{8})
//
// returns
//
//	[]int{1,2,3,2,3,8}
func Concat[A any](inputSlices ...[]A) []A {
	totalLength := Sum(Map(inputSlices, func(slc []A) int {
		return len(slc)
	}))
	result := make([]A, totalLength)
	idx := 0
	for _, slc := range inputSlices {
		for _, value := range slc {
			result[idx] = value
			idx++
		}
	}
	return result
}

// SortBy (similar to [sort.Slice]) sorts elements in slice using a comparator (less) function.
// The corresponding slice is altered and returned.
func SortBy[A any](slc []A, less func(A, A) bool) []A {
	sort.Slice(slc, func(i, j int) bool {
		return less(slc[i], slc[j])
	})
	return slc
}

// SortStable (similar to [sort.SliceStable]) sorts elements in slice using a comparator (less) function.
// The corresponding slice is altered and returned.
func SortStable[A any](slc []A, less func(A, A) bool) []A {
	sort.SliceStable(slc, func(i, j int) bool {
		return less(slc[i], slc[j])
	})
	return slc
}

// Every returns true if `pred` returns true for every value of a sequence. Example:
//
//	Every([]int{2, 3, 4}, func(n int) bool { return n > 1 }) // returns true
//	Every([]int{2, 0, 4}, func(n int) bool { return n > 1 }) // returns false
//	Every([]int{}, func(n int) bool { return n > 1 })        // returns true
func Every[A any](slc []A, pred func(A) bool) bool {
	for _, value := range slc {
		if !pred(value) {
			return false
		}
	}
	return true
}

// Some returns true if `pred` returns true for at least one value of a sequence.
//
//	Some([]int{2, 3, 4}, func(n int) bool { return n < 1 }) // returns false
//	Some([]int{2, 0, 4}, func(n int) bool { return n < 1 }) // returns true
//	Some([]int{}, func(n int) bool { return n < 1 })        // returns false
func Some[A any](slc []A, pred func(A) bool) bool {
	for _, value := range slc {
		if pred(value) {
			return true
		}
	}
	return false
}

// PartitionBy imagines a slice to be devided into blocks for which `f` returns a
// different value. It returns the blocks as slice of slices. No entries are removed.
// Example:
//
//	PartitionBy([]int{1, 2, 3, 4, 5}, func(n int) int {
//	  return n / 2
//	})
//
// returns `[][]int{{1}, {2, 3}, {4, 5}}`
func PartitionBy[A any, B comparable](slc []A, f func(A) B) [][]A {
	result := make([][]A, 0)
	var previousPredValue B
	chunkStart := 0
	for i, value := range slc {
		thisPredValue := f(value)
		if i != 0 && previousPredValue != thisPredValue {
			result = append(result, slc[chunkStart:i])
			chunkStart = i
		}
		previousPredValue = thisPredValue
	}
	if chunkStart < len(slc) {
		result = append(result, slc[chunkStart:])
	}
	return result
}

// PartitionAt imagines a slice to be devided into blocks by `compareValue`. It
// returns the blocks as slice of slices with entries of `compareValue` removed.
// Example:
//
//	PartitionAt([]int{1, 2, 0, 3, 4, 0, 5}, 0)
//
// returns `[][]int{{1, 2}, {3, 4}, {5}}`
func PartitionAt[C comparable](slc []C, compareValue C) [][]C {
	result := make([][]C, 0)
	chunkStart := 0
	i := 0
	for i < len(slc) {
		if slc[i] == compareValue {
			result = append(result, slc[chunkStart:i])
			chunkStart = i + 1
			i++
		}
		i++
	}
	if chunkStart < len(slc) {
		result = append(result, slc[chunkStart:i])
	}
	return result
}

// Identity plainly returns it's input value. It can be used e.g. for PartitionBy,
// Example:
//
//	PartitionBy([]int{1,1,2,1,3,3}, Identity[int])
//
// returns `[][]int{{1,1}, {2}, {1}, {3,3}}`
func Identity[A any](v A) A {
	return v
}

// IndexOf returns the index of elt in slice slc
func IndexOf[T comparable](slc []T, elt T) (int, bool) {
	for i, e := range slc {
		if e == elt {
			return i, true
		}
	}
	return -1, false
}

// Range returns a slice containing values 0,...,n-1 in increasing order
func Range(n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = i
	}
	return result
}

// RangeFromTo returns a slice containing values from,...,to-1 in increasing order
func RangeFromTo(from, to int) []int {
	result := make([]int, to-from)
	for i := 0; i < to-from; i++ {
		result[i] = from + i
	}
	return result
}

// Randomize shuffles the elements in a slice in random order. Each order
// is returned with equal probability.
func Randomize[T any](slice []T) {
	for i := 0; i < len(slice)-1; i++ {
		j := rand.Intn(len(slice) - i)
		if j != 0 {
			slice[i], slice[i+j] = slice[i+j], slice[i]
		}
	}
}

// Randomized returns a slice containing the values of the original slice
// in random order. Each order is returned with equal probability.
func Randomized[T any](slice []T) []T {
	result := make([]T, len(slice))
	copy(result, slice)
	Randomize(result)
	return result
}

// SortLex takes a slice of slices and sorts it in lexicographic
// order.
func SortLex[T cmp.Ordered](slice [][]T) {
	sort.Slice(slice, func(i, k int) bool {
		return lessSlices(slice[i], slice[k])
	})
}

func lessSlices[T cmp.Ordered](a, b []T) bool {
	for n := 0; n < len(a); n++ {
		if n >= len(b) {
			return false
		} else if cmp.Less(a[n], b[n]) {
			return true
		} else if cmp.Less(b[n], a[n]) {
			return false
		}
	}
	return len(a) < len(b)
}

// SortedLex copies a slice of slices, sorts it in lexicographic
// order and returns it.
func SortedLex[T cmp.Ordered](slice [][]T) [][]T {
	slc := CopySlice(slice)
	SortLex(slc)
	return slc
}

// bigFactorial returns the nth factorial using big ints.
func bigFactorial(n *big.Int) *big.Int {
	if n.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0)
	}
	result := big.NewInt(1)
	for i := big.NewInt(1); i.Cmp(n) <= 0; i.Add(i, big.NewInt(1)) {
		result.Mul(result, i)
	}
	return result
}

// nthPermutationOf returns the nth permutation of a slice.
//
// Generally, There are len(slice)! different permutations of the slice.
// nthPermutationOf can be called with indices from 0 to len(slice)!-1
// each returning a different permutation of slice.
func nthPermutationOf[T any](slice []T, n *big.Int) []T {
	idxList := golist.MakeList(Range(len(slice))...)
	rest := new(big.Int)
	rest.Set(n)
	result := make([]T, len(slice))
	var idx *big.Int
	for n := len(slice); n > 0; n-- {
		rest, idx = rest.DivMod(rest, big.NewInt(int64(n)), big.NewInt(0))
		idxElt := idxList.EltAtIndex(int(idx.Int64()))
		result[n-1] = slice[idxElt.Value()]
		idxList.Remove(idxElt)
	}
	return result
}

// Permutations takes a slice of values and returns a slice of slices
// containing each possible permutation.
//
// The original slice needs to be quite small. (The function itself imposes
// a hard boundary of 20, but practically also lower values will likely exceed
// the memory.) Use PermutationsIter() if you don't want to use up memory.
//
// Example:
//
//	Permutations([]int{1,2,3})
//
// returns [][]int{{1,2,3}, {1,3,2}, {2,1,3}, {2,3,1}, {3,1,2}, {3,2,1}}
// (in arbitrary order).
func Permutations[T any](slice []T) [][]T {
	if len(slice) == 0 {
		return [][]T{}
	}
	numPermutationsBig := bigFactorial(big.NewInt(int64(len(slice))))
	Assert(numPermutationsBig.IsInt64(), "Cannot make permutations with huge slices (as length %d)", len(slice))
	numPermutations := numPermutationsBig.Int64()
	result := make([][]T, numPermutations)
	for i := int64(0); i < numPermutations; i++ {
		result[i] = nthPermutationOf(slice, big.NewInt(int64(i)))
	}
	return result
}

// PermutationsIter takes a slice of values and returns an iterator to loop across
// all iterations of elements of that slice.
//
// Different from Permutations(), the concrete permutations are not calculated, so
// the size of the slice can be large without using space.
//
// Example:
//
//	for perm := range PermutationsIter([]int{1,2}) {
//	    fmt.Println(perm)
//	}
//
// prints []int{1,2} and []int{2,1} (in arbitrary order).
func PermutationsIter[T any](slice []T) iter.Seq[[]T] {
	numPermutations := bigFactorial(big.NewInt(int64(len(slice))))
	return func(yield func(perm []T) bool) {
		for i := big.NewInt(0); i.Cmp(numPermutations) < 0; i.Add(i, big.NewInt(1)) {
			if !yield(nthPermutationOf(slice, i)) {
				return
			}
		}
	}
}
