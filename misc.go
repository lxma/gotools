package gotools

import (
	"fmt"
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

// Ternary is a weak substitute for the missing ternary operator. I consider it as an act of self-defense
// from strange language lawyers that think that
//
//	var x int
//	if useOne {
//	   x = 1
//	} else {
//	   x = 0
//	}
//
// is better readable than
//
//	x = useOne ? 1 : 0
func Ternary[A any](condition bool, trueVal A, falseVal A) A {
	if condition {
		return trueVal
	} else {
		return falseVal
	}
}

// Assert panics if a condition is not met. It should be used to signal a situation
// that cannot occur. (So, if the situation occurs, there is a bug in the program
// logic itself.)
func Assert(condition bool, parameters ...any) {
	if !condition {
		message := fmt.Errorf("unexpected assert failure")
		if len(parameters) > 0 {
			message = fmt.Errorf(parameters[0].(string), parameters[1:]...)
		}
		panic(message)
	}
}

// Assume behaves technically like assert. The difference is the semantic. Assume should be
// used when you – during dirty programming – assume a situation should not appear and you
// avoid proper error handling. In case this is violated, you search for different types of
// errors. Also, you should replace assumes with proper error handling before going productive.
func Assume(condition bool, parameters ...any) {
	if !condition {
		message := fmt.Errorf("unexpected wrong assumtion")
		if len(parameters) > 0 {
			message = fmt.Errorf(parameters[0].(string), parameters[1:]...)
		}
		panic(message)
	}
}
