package gotools

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestAbs(t *testing.T) {
	assert.Equal(t, 3.5, Abs(3.5), "abs float")
	assert.Equal(t, 3.5, Abs(-3.5), "abs negative float")
	assert.Equal(t, 2, Abs(2), "abs int")
	assert.Equal(t, 2, Abs(-2), "abs negative int")
}

func TestStringToInt(t *testing.T) {
	assert.Equal(t, 14, StringToInt("14"))
	assert.Equal(t, -8, StringToInt("-8"))
	assert.Panics(t, func() {
		StringToInt("not an int")
	})
}

func TestBigInt(t *testing.T) {
	var bi *big.Int = BigInt(4)
	assert.Equal(t, "4", bi.String())
	bi = BigInt(int16(3))
	assert.Equal(t, "3", bi.String())
}

func TestGCD(t *testing.T) {
	assert.Equal(t, 2, GCD(10, 12))
	assert.Equal(t, 10, GCD(30, 20))
	assert.Equal(t, 1, GCD(1, 20))
	assert.Equal(t, "2", BigGCD(BigInt(10), BigInt(12)).String())
	assert.Equal(t, "10", BigGCD(BigInt(30), BigInt(20)).String())
	assert.Equal(t, "1", BigGCD(BigInt(1), BigInt(20)).String())
}

func TestLCM(t *testing.T) {
	assert.Equal(t, 60, LCM(10, 12))
	assert.Equal(t, 60, LCM(30, 20))
	assert.Equal(t, "60", BigLCM(BigInt(10), BigInt(12)).String())
	assert.Equal(t, "60", BigLCM(BigInt(30), BigInt(20)).String())
	assert.Equal(t, "20", BigLCM(BigInt(1), BigInt(20)).String())
}
