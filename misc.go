package gotools

import (
	"regexp"
	"strconv"
)

// GetIntegersInString returns all integers in a given string
func GetIntegersInString(line string) []int {
	re := regexp.MustCompile("(-?\\d+)")
	matches := re.FindAllStringSubmatch(line, -1)
	result := make([]int, len(matches))
	for i, match := range matches {
		number, _ := strconv.Atoi(match[1])
		result[i] = number
	}
	return result
}

// PartialR (similar to [PartialL] takes a function with two arguments and one (the first)
// argument. It returns a function with only one argument. Example
//
//	func subtract(a int, b int) int { return a - b }
//	subtractFromTen := PartialR(subtract, 10)
//
// Now, subtractFromTen(3) returns 7. This is handy for mapping, e.g.:
//
//	Map([]int{1, 2, 3}, PartialR(subtract, 10))
//
// returns
//
//	[]int{9, 8, 7}
func PartialR[A any, B any, ResultType any](f func(A, B) ResultType, a A) func(B) ResultType {
	return func(b B) ResultType {
		return f(a, b)
	}
}

// PartialL (similar to [PartialR] takes a function with two arguments and one (the second)
// argument. It returns a function with only one argument. Example
//
//	func add(a int, b int) int { return a + b }
//	addTwo := PartialL(add, 2)
//
// now, addTwo(3) returns 5
// This is handy for mapping, e.g.:
//
//	Map([]int{1, 2, 3}, PartialL(add,10))
//
// returns
//
//	[]int{11, 12, 13}
func PartialL[A any, B any, ResultType any](f func(A, B) ResultType, b B) func(A) ResultType {
	return func(a A) ResultType {
		return f(a, b)
	}
}

// Partial2R (similar to [Partial2L] takes a function with three arguments and the first two arguments.
// It returns a function with only one argument. Example
//
//	func subtract2(a int, b int, c int) int { return a - b - c }
//	f := Partial2R(subtract2, 10, 2)
//
// now, f(3) returns 5
func Partial2R[A any, B any, C any, ResultType any](f func(A, B, C) ResultType, a A, b B) func(C) ResultType {
	return func(c C) ResultType {
		return f(a, b, c)
	}
}

// Partial2L (similar to [Partial2R] takes a function with three arguments and the last two arguments.
// It returns a function with only one argument. Example
//
//	func subtract2(a int, b int, c int) int { return a - b - c }
//	subtractThree := Partial2L(subtract2, 1, 2)
//
// now, subtractThree(10) returns 7
func Partial2L[A any, B any, C any, ResultType any](f func(A, B, C) ResultType, b B, c C) func(A) ResultType {
	return func(a A) ResultType {
		return f(a, b, c)
	}
}
