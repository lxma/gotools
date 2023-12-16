package gotools

import "fmt"

// Set is a set of objects in the mathematical sense. I.e. it is
// unsorted and can contain each element only once. It provides
// expected methods such as intersections, unions, and information functions.
//
// Create a set from a slice using MakeSet()
// A set can be transformed into a slice using ToSlice()
//
// There are no .Empty or .Len functions, as this can be done via the regular len(...)
// function. .Equal can be done via reflect.DeepEqual()k
type Set[C comparable] map[C]struct{}

// IsSubsetOf checks if the set is a subset of (or equal to) another set.
// Example:
//
//	MakeSet([]int{1, 2, 3}).IsSubsetOf([]int{1, 2})
//
// returns true
func (s Set[C]) IsSubsetOf(super Set[C]) bool {
	for elt, _ := range s {
		if !super.Contains(elt) {
			return false
		}
	}
	return true
}

// IsSupersetOf checks if the set is a superset of (or equal to) another set.
// Example:
//
//	MakeSet([]int{1, 2}).IsSupersetOf([]int{1, 2, 3})
//
// returns true
func (s Set[C]) IsSupersetOf(sub Set[C]) bool {
	for elt, _ := range sub {
		if !s.Contains(elt) {
			return false
		}
	}
	return true
}

// MakeSet creates a set from a slice, e.g.
//
//	MakeSet([]int{2,4,6,2})
//
// creates the set {2,4,6}.
func MakeSet[C comparable](input []C) Set[C] {
	result := Set[C]{}
	for _, elt := range input {
		result[elt] = struct{}{}
	}
	return result
}

// Copy copies a set. It copies only the set itself, not the values.
func (s Set[C]) Copy() Set[C] {
	result := Set[C]{}
	for elt, _ := range s {
		result.Add(elt)
	}
	return result
}

// Contains checks if a given set contains an element
// Example:
//
//	MakeSet([]int{1, 2, 3}).Contains(1)
//
// returns true
func (s Set[C]) Contains(elt C) bool {
	_, ok := s[elt]
	return ok
}

// Add adds one or more elements to a set. The set is altered and it is returned.
// It's ok to add an existing element (the set is not changed).
// Examples:
//
//	MakeSet([]int{1, 2, 3}).Add(4)
//	MakeSet([]int{1, 2, 3}).Add(2,3,4,5)
func (s Set[C]) Add(elts ...C) Set[C] {
	for _, elt := range elts {
		s[elt] = struct{}{}
	}
	return s
}

// Delete deletes one element from a set. The set is altered and it is returned.
// It's ok to delete an element that does not exist (the set is not changed).
// Examples:
//
//	MakeSet([]int{1, 2, 3}).Delete(4)
func (s Set[C]) Delete(elt C) Set[C] {
	delete(s, elt)
	return s
}

// AddSet (similar to Union) adds elements from another set to a set.
// It alters and returns the given set.
//
//	MakeSet([]int{1, 2}).AddSet(MakeSet([]int{3, 4})
//
// returns the set {1, 2, 3, 4}
func (s Set[C]) AddSet(secondSet Set[C]) Set[C] {
	for key, _ := range secondSet {
		s[key] = struct{}{}
	}
	return s
}

// SubtractSet (similar to SetDifference) removes elements from another set to a set.
// It alters and returns the given set.
//
//	MakeSet([]int{1, 2}).SubtractSet(MakeSet([]int{2, 3})
//
// returns the set {1}
func (s Set[C]) SubtractSet(secondSet Set[C]) Set[C] {
	for key, _ := range secondSet {
		delete(s, key)
	}
	return s
}

// Intersect (similar to Intersection) removes elements /not/ contained in another set.
// It alters and returns the given set.
//
//	MakeSet([]int{1, 2}).Intersect(MakeSet([]int{2, 3})
//
// returns the set {2}
func (s Set[C]) Intersect(secondSet Set[C]) Set[C] {
	for key, _ := range s {
		if !secondSet.Contains(key) {
			delete(s, key)
		}
	}
	return s
}

// ToSlice turns a set into a slice of its elements. The order
// of the elements is undefined.
//
//	MakeSet([]int{1, 2}).ToSlice()
//
// returns []int{1,2} or []int{2,1}
func (s Set[C]) ToSlice() []C {
	slice := make([]C, len(s))
	i := 0
	for elt, _ := range s {
		slice[i] = elt
		i++
	}
	return slice
}

// GetArbitraryElement returns an arbitrary element from a set.
// It panics when called on an empty set.
// Example:
//
//	MakeSet([]int{1, 2}).GetArbitraryElement()
//
// returns 1 or 2
func (s Set[C]) GetArbitraryElement() C {
	for elt, _ := range s {
		return elt
	}
	panic(fmt.Errorf("requested element from an empty set"))
}

// Union (similar to .AddSet) returns the union of two sets. The original sets
// are not changed.
// Example:
//
//	Union(MakeSet([]int{1, 2}), MakeSet([]int{2, 3}))
//
// returns {1,2,3}
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

// Intersection (similar to .Intersect) delivers the intersection of two sets. The original
// sets are not changed.
// Example:
//
//	Intersection(MakeSet([]int{1, 2}), MakeSet([]int{2, 3}))
//
// returns {2}
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

// SetDifference (similar to .SubtractSet) delivers the set difference of two sets.
// The original sets are not changed.
// Example:
//
//	SetDifference(MakeSet([]int{1, 2}), MakeSet([]int{2, 3}))
//
// returns {1}
func SetDifference[C comparable](set1 Set[C], set2 Set[C]) Set[C] {
	result := Set[C]{}
	for key, _ := range set1 {
		if !set2.Contains(key) {
			result[key] = struct{}{}
		}
	}
	return result
}

// MapSet maps a function to all elements of a set. It returns a new set, containing all
// function results.
// Example:
//
//	MapSet(MakeSet([]int{1, 2, 3}), func(a int) int { return 2 * a })
//
// returns {2, 4, 6}
func MapSet[A comparable, B comparable](s Set[A], f func(A) B) Set[B] {
	mappedSet := Set[B]{}
	for elt := range s {
		mappedSet.Add(f(elt))
	}
	return mappedSet
}
