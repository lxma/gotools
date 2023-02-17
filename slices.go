package gotools

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

func Map[A any, B any](arr []A, f func(A) B) []B {
	result := make([]B, len(arr))
	for i, value := range arr {
		result[i] = f(value)
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
