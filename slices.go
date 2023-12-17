package gotools

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
func Repeat[A any](elt A, n int) []A {
	result := make([]A, n)
	for i := 0; i < n; i++ {
		result[i] = elt
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
func Transpose[C any](data [][]C) [][]C {
	if len(data) == 0 {
		return [][]C{}
	}
	if len(data[0]) == 0 {
		return [][]C{{}}
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
		for _, val := range slc {
			result[idx] = val
			idx++
		}
	}
	return result
}
