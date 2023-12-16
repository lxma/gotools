package gotools

// MapAtHash provides a map functionality for hashmaps, i.e. it applies a
// given function to the /values/ of a hashmap. Example:
//
//	MapAtHash(map[string]string{"a": "1", "b": "2"}, StringToInt)
//
// returns
//
//	map[string]int{"a": 1, "b": 2},
func MapAtHash[KeyType comparable, ValueTypeA any, ValueTypeB any](
	m map[KeyType]ValueTypeA,
	f func(ValueTypeA) ValueTypeB) map[KeyType]ValueTypeB {

	result := make(map[KeyType]ValueTypeB)
	for key, value := range m {
		result[key] = f(value)
	}
	return result
}

// MapHashKeys provides a map functionality for hashmaps, i.e. it applies a
// given function to the /keys/ of a hashmap. Example:
//
//	MapHashKeys(map[string]string{"1": "a", "2": "b"}, StringToInt)
//
// returns
//
//	map[int]string{1: "a", 2: "b"}
func MapHashKeys[KeyTypeA comparable, KeyTypeB comparable, ValueType any](
	m map[KeyTypeA]ValueType,
	f func(KeyTypeA) KeyTypeB) map[KeyTypeB]ValueType {
	result := make(map[KeyTypeB]ValueType)
	for key, value := range m {
		result[f(key)] = value
	}
	return result
}

// MapHasKey checks if a hashmap contains a given key
func MapHasKey[A comparable, B any](m map[A]B, key A) bool {
	_, ok := m[key]
	return ok
}
