package gotools

func TakeWhile[A any](slc []A, f func(A) bool) []A {
	for i, elt := range slc {
		if !f(elt) {
			return slc[0:i]
		}
	}
	return slc
}

func DropWhile[A any](slc []A, f func(A) bool) []A {
	for i, elt := range slc {
		if !f(elt) {
			return slc[i:]
		}
	}
	return []A{}
}

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

func Repeat[A any](elt A, n int) []A {
	result := make([]A, n)
	for i := 0; i < n; i++ {
		result[i] = elt
	}
	return result
}

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

func Map[A any, B any](slc []A, f func(A) B) []B {
	result := make([]B, len(slc))
	for i, value := range slc {
		result[i] = f(value)
	}
	return result
}

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

func Sum[N Number](slc []N) N {
	var result N
	for _, val := range slc {
		result += val
	}
	return result
}

func Prod[N Number](slc []N) N {
	var result N = N(1)
	for _, val := range slc {
		result *= val
	}
	return result
}

func Reverse[A any](slc []A) []A {
	reversed := make([]A, len(slc))
	for i, elt := range slc {
		reversed[len(slc)-i-1] = elt
	}
	return reversed
}

func CopySlice[C any](slc []C) []C {
	newSlice := make([]C, len(slc))
	copy(newSlice, slc)
	return newSlice
}
