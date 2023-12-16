package gotools

import "strconv"

// Number should cover all types for which Go offers basic arithmetic operations
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

// Min returns the minimum of multiple given numbers. It returns 0 if no number is provided.
func Min[N Number](nums ...N) N {
	if len(nums) == 0 {
		return 0
	} else {
		min := nums[0]
		for i := 1; i < len(nums); i++ {
			if nums[i] < min {
				min = nums[i]
			}
		}
		return min
	}
}

// Max returns the maximum of multiple given numbers. It returns 0 if no number is provided.
func Max[N Number](nums ...N) N {
	if len(nums) == 0 {
		return 0
	} else {
		max := nums[0]
		for i := 1; i < len(nums); i++ {
			if nums[i] > max {
				max = nums[i]
			}
		}
		return max
	}
}

// StringToInt converts the given string to an integer. It panics if this is not possible.
func StringToInt(st string) int {
	result, err := strconv.Atoi(st)
	if err != nil {
		panic(err)
	}
	return result
}

// Abs returns the absolute value of a number
func Abs[N Number](n N) N {
	if n < 0 {
		return -n
	} else {
		return n
	}
}
