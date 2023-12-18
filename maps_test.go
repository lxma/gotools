package gotools

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestMapAtHash(t *testing.T) {
	assert.Equal(t, map[string]int{"a": 1, "b": 2}, MapAtHash(map[string]string{"a": "1", "b": "2"}, StringToInt),
		"simple mapping")
	assert.Equal(t, map[string]int{}, MapAtHash(map[string]string{}, StringToInt),
		"empty mapping")
}

func TestMapHashKeys(t *testing.T) {
	assert.Equal(t, map[int]string{1: "a", 2: "b"}, MapHashKeys(map[string]string{"1": "a", "2": "b"}, StringToInt),
		"simple mapping")
	assert.Equal(t, map[int]string{}, MapHashKeys(map[string]string{}, StringToInt),
		"empty mapping")
}

func TestMapHasKey(t *testing.T) {
	assert.True(t, MapHasKey(map[int]string{1: "a", 2: "b"}, 1), "map has key 1")
	assert.False(t, MapHasKey(map[int]string{1: "a", 2: "b"}, 3), "map does not have key 3")
}

func TestGetKeysAndValues(t *testing.T) {
	values := GetValues(map[int]string{1: "one", 2: "two"})
	sort.Strings(values)
	assert.Equal(t, []string{"one", "two"}, values)
	assert.Equal(t, []string{}, GetValues(map[int]string{}))
	keys := GetKeys(map[int]string{1: "one", 2: "two"})
	sort.Ints(keys)
	assert.Equal(t, []int{1, 2}, keys)
	assert.Equal(t, []int{}, GetKeys(map[int]string{}))
}
