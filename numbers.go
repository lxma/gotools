package gotools

import (
	"math/big"
	"strconv"
)

// Number should cover all types for which Go offers basic arithmetic operations
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
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

// BigInt creates a big integer from any given integer type. That's how
// [big.NewInt] ought to be typed.
func BigInt[I Integer](n I) *big.Int {
	return big.NewInt(int64(n))
}

// GCD returns the greatest common divisor of two integers.
func GCD[I Integer](a, b I) I {
	if a < b {
		return GCD(b, a)
	} else {
		c := a % b
		if c == 0 {
			return b
		} else {
			return GCD(b, c)
		}
	}
}

// LCM returns the least common multiple of two integers
func LCM[I Integer](a, b I) I {
	return a / GCD(a, b) * b
}

// BigGCD returns the greatest common divisor of two [big.Int] variables.
func BigGCD(a, b *big.Int) *big.Int {
	if a.Cmp(b) == -1 {
		return BigGCD(b, a)
	} else {
		c := big.NewInt(0)
		c.Mod(a, b)
		if c.Cmp(big.NewInt(0)) == 0 {
			return b
		} else {
			return BigGCD(b, c)
		}
	}
}

// BigLCM returns the least common multiple of two [big.Int] variables
func BigLCM(a, b *big.Int) *big.Int {
	result := new(big.Int)
	result.Set(a)
	result.Div(result, BigGCD(a, b))
	result.Mul(result, b)
	return result
}
