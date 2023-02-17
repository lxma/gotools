package gotools

func MapAtHash[KeyType comparable, ValueTypeA any, ValueTypeB any](
	m map[KeyType]ValueTypeA,
	f func(ValueTypeA) ValueTypeB) map[KeyType]ValueTypeB {

	result := make(map[KeyType]ValueTypeB)
	for key, value := range m {
		result[key] = f(value)
	}
	return result
}

func MapHashKeys[KeyTypeA comparable, KeyTypeB comparable, ValueType any](
	m map[KeyTypeA]ValueType,
	f func(KeyTypeA) KeyTypeB) map[KeyTypeB]ValueType {
	result := make(map[KeyTypeB]ValueType)
	for key, value := range m {
		result[f(key)] = value
	}
	return result
}

func MapHasKey[A comparable, B any](m map[A]B, key A) bool {
	_, ok := m[key]
	return ok
}
