package gotools

import "fmt"

type Set[C comparable] map[C]struct{}

func (s Set[C]) IsSubsetOf(super Set[C]) bool {
	for elt, _ := range s {
		if !super.Contains(elt) {
			return false
		}
	}
	return true
}

func (s Set[C]) IsSupersetOf(super Set[C]) bool {
	for elt, _ := range super {
		if !s.Contains(elt) {
			return false
		}
	}
	return true
}

func MakeSet[C comparable](input []C) Set[C] {
	result := Set[C]{}
	for _, elt := range input {
		result[elt] = struct{}{}
	}
	return result
}

func (s Set[C]) Copy() Set[C] {
	result := Set[C]{}
	for elt, _ := range s {
		result.Add(elt)
	}
	return result
}

func (s Set[C]) Contains(elt C) bool {
	_, ok := s[elt]
	return ok
}

func (s Set[C]) Add(elt C) Set[C] {
	s[elt] = struct{}{}
	return s
}

func (s Set[C]) Delete(elt C) Set[C] {
	delete(s, elt)
	return s
}

func (s Set[C]) AddSet(secondSet Set[C]) Set[C] {
	for key, _ := range secondSet {
		s[key] = struct{}{}
	}
	return s
}

func (s Set[C]) SubtractSet(secondSet Set[C]) Set[C] {
	for key, _ := range secondSet {
		delete(s, key)
	}
	return s
}

func (s Set[C]) Intersect(secondSet Set[C]) Set[C] {
	for key, _ := range s {
		if !secondSet.Contains(key) {
			delete(s, key)
		}
	}
	return s
}

func (s Set[C]) ToSlice() []C {
	slice := make([]C, len(s))
	i := 0
	for elt, _ := range s {
		slice[i] = elt
		i++
	}
	return slice
}

func (s Set[C]) GetArbitraryElement() C {
	for elt, _ := range s {
		return elt
	}
	panic(fmt.Errorf("requested element from an empty set"))
}

func Union[C comparable](set1 Set[C], set2 Set[C]) Set[C] {
	result := Set[C]{}
	for key, _ := range set1 {
		result[key] = struct{}{}
	}
	for key, _ := range set2 {
		result[key] = struct{}{}
	}
	return result
}

func SetAddElement[C comparable](set1 Set[C], elt C) Set[C] {
	result := set1.Copy()
	result.Add(elt)
	return result
}

func SetDeleteElement[C comparable](set1 Set[C], elt C) Set[C] {
	result := set1.Copy()
	result.Delete(elt)
	return result
}

func Intersection[C comparable](set1 Set[C], set2 Set[C]) Set[C] {
	if len(set1) > len(set2) {
		return Intersection(set2, set1)
	}
	result := Set[C]{}
	for key, _ := range set1 {
		if set2.Contains(key) {
			result[key] = struct{}{}
		}
	}
	return result
}

func SetDifference[C comparable](set1 Set[C], set2 Set[C]) Set[C] {
	result := Set[C]{}
	for key, _ := range set1 {
		if !set2.Contains(key) {
			result[key] = struct{}{}
		}
	}
	return result
}

func MapSet[A comparable, B comparable](s Set[A], f func(A) B) Set[B] {
	mappedSet := Set[B]{}
	for elt := range s {
		mappedSet.Add(f(elt))
	}
	return mappedSet
}
