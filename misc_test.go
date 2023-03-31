package gotools

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetIntegersInString(t *testing.T) {
	assert.Equal(t, []int{1, -2, 3}, GetIntegersInString("Und 1 und -2, ,+3"))
}
