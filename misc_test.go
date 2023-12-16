package gotools

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetIntegersInString(t *testing.T) {
	assert.Equal(t, []int{1, -2, 3}, GetIntegersInString("Und 1 und -2, ,+3"))
}

func TestPartial(t *testing.T) {
	f := Partial(func(a int, b int) int { return a / b }, 2)
	assert.Equal(t, 2, f(5))
	g := PartialR(func(a int, b int) int { return a / b }, 10)
	assert.Equal(t, 3, g(3))
	h := Partial2(func(a int, b int, c int) int { return a + b + c }, 2, 3)
	assert.Equal(t, 9, h(4))
}
