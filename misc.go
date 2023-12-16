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

// Partial takes a function with two arguments and one (the second) argument.
// It returns a function with only one argument. Example
//
//	func add(a int, b int) int { return a + b }
//	addTwo := Partial(add, 2)
//
// now, addTwo(3) returns 5
// This is handy for mapping, e.g.:
//
//	Map([]int{1, 2, 3}, Partial(add,10))
//
// returns
//
//	[]int{11, 12, 13}
func Partial[A any, B any, ResultType any](f func(A, B) ResultType, b B) func(A) ResultType {
	return func(a A) ResultType {
		return f(a, b)
	}
}

// PartialR takes a function with two arguments and one (the first) argument.
// It returns a function with only one argument. Example
//
//	func subtract(a int, b int) int { return a - b }
//	subtractFromTen := Partial(add, 10)
//
// now, subtractFromTen(3) returns 7
// This is handy for mapping, e.g.:
//
//	Map([]int{1, 2, 3}, Partial(subtract, 10))
//
// returns
//
//	[]int{9, 8, 7}
func PartialR[A any, B any, ResultType any](f func(A, B) ResultType, a A) func(B) ResultType {
	return func(b B) ResultType {
		return f(a, b)
	}
}

// Partial2 takes a function with three arguments and two arguments.
// It returns a function with only one argument. Example
//
//	func add(a int, b int, c int) int { return a + b + c }
//	addTwo := Partial(add, 1, 1)
//
// now, addTwo(3) returns 5
func Partial2[A any, B any, C any, ResultType any](f func(A, B, C) ResultType, b B, c C) func(A) ResultType {
	return func(a A) ResultType {
		return f(a, b, c)
	}
}
