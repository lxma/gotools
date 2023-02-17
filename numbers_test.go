package gotools

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbs(t *testing.T) {
	assert.Equal(t, 3.5, Abs(3.5), "abs float")
	assert.Equal(t, 3.5, Abs(-3.5), "abs negative float")
	assert.Equal(t, 2, Abs(2), "abs int")
	assert.Equal(t, 2, Abs(-2), "abs negative int")
}

func TestMax(t *testing.T) {
	assert.Equal(t, 2, Max(2, 1))
	assert.Equal(t, 0, Max[int]())
	assert.Equal(t, 2.2, Max(2.1, 2.2, 1.0))
	assert.Equal(t, 0.0, Max([]float64{}...))
}

func TestMin(t *testing.T) {
	assert.Equal(t, 1, Min(2, 1))
	assert.Equal(t, 0, Min[int]())
	assert.Equal(t, 1.1, Min(1.2, 2.2, 1.1))
	assert.Equal(t, 0.0, Min([]float64{}...))
}

func TestStringToInt(t *testing.T) {
	assert.Equal(t, 14, StringToInt("14"))
	assert.Equal(t, -8, StringToInt("-8"))
	assert.Panics(t, func() {
		StringToInt("not an int")
	})
}
