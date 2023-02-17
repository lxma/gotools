package gotools

import "strconv"

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

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

func StringToInt(st string) int {
	result, err := strconv.Atoi(st)
	if err != nil {
		panic(err)
	}
	return result
}

func Abs[N Number](n N) N {
	if n < 0 {
		return -n
	} else {
		return n
	}
}
