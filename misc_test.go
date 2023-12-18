package gotools

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetIntegersInString(t *testing.T) {
	assert.Equal(t, []int{1, -2, 3}, GetIntegersInString("Und 1 und -2, ,+3"))
}

func TestPartial(t *testing.T) {
	f1 := PartialR(func(a int, b int) int { return a / b }, 10)
	assert.Equal(t, 3, f1(3))
	f2 := PartialL(func(a int, b int) int { return a / b }, 2)
	assert.Equal(t, 2, f2(5))
	f3 := Partial2R(func(a int, b int, c int) int { return a - b - c }, 10, 3)
	assert.Equal(t, 5, f3(2))
	f4 := Partial2L(func(a int, b int, c int) int { return a - b - c }, 2, 3)
	assert.Equal(t, 5, f4(10))

}
