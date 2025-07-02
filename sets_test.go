package gotools

import (
    "cmp"
    "github.com/stretchr/testify/assert"
    "slices"
    "sort"
    "testing"
)

func assertIntSetEquals(t *testing.T, expected []int, actualSet Set[int], msgAndArgs ...interface{}) {
    actualSlice := actualSet.ToSlice()
    sort.Ints(actualSlice)
    sort.Ints(expected)
    assert.Equal(t, expected, actualSlice, msgAndArgs...)
}

func assertStringSetEquals(t *testing.T, expected []string, actualSet Set[string], msgAndArgs ...interface{}) {
    actualSlice := actualSet.ToSlice()
    sort.Strings(actualSlice)
    sort.Strings(expected)
    assert.Equal(t, expected, actualSlice, msgAndArgs...)
}

func TestMakeSet(t *testing.T) {
    assertIntSetEquals(t, []int{1, 2, 3, 4}, MakeSet(4, 2, 3, 1, 2), "Simple int set")
    assertIntSetEquals(t, []int{}, MakeSet[int](), "Empty int set")
    assertStringSetEquals(t, []string{"a", "b"}, MakeSet("a", "b", "a"), "Simple string set")
    assertStringSetEquals(t, []string{}, MakeSet[string](), "Empty string set")
}

func TestSet_Contains(t *testing.T) {
    assert.Equal(t, true, MakeSet(1, 2, 3).Contains(1), "simple contains")
    assert.Equal(t, false, MakeSet(1, 2, 3).Contains(0), "simple not contains")
    assert.Equal(t, false, MakeSet[int]().Contains(1), "empty not contains")
}

func TestSet_Add(t *testing.T) {
    assertIntSetEquals(t, []int{1, 2}, MakeSet(2).Add(1), "simple add")
    assertIntSetEquals(t, []int{1}, Set[int]{}.Add(1), "add to empty set")
    assertIntSetEquals(t, []int{1, 2, 3}, MakeSet(2).Add(1, 2, 3), "add multiple values")
    assertIntSetEquals(t, []int{}, Set[int]{}.Add(), "add nothing to empty set")
}

func TestSet_AddSet(t *testing.T) {
    assertIntSetEquals(t, []int{1, 2, 3, 4}, MakeSet(1, 2).AddSet(MakeSet(3, 4)), "simple add set")
    assertIntSetEquals(t, []int{3, 4}, Set[int]{}.AddSet(MakeSet(3, 4)), "add set to empty set")
}

func TestSet_Copy(t *testing.T) {
    s1 := MakeSet(1)
    s2 := s1.Copy()
    s1.Add(2)
    s2.Add(3)
    assertIntSetEquals(t, []int{1, 2}, s1, "changes at s2 are not made to s1")
    assertIntSetEquals(t, []int{1, 3}, s2, "changes at s1 are not made to s2")
}

func TestSet_Delete(t *testing.T) {
    s := MakeSet(1, 2)
    s.Delete(1)
    assertIntSetEquals(t, []int{2}, s, "simple delete")
}

func TestSet_Intersect(t *testing.T) {
    assertIntSetEquals(t, []int{2}, MakeSet(1, 2).Intersect(MakeSet(2, 3)), "simple intersection")
    assertIntSetEquals(t, []int{}, MakeSet(1, 2).Intersect(Set[int]{}), "intersect with empty set")
    assertIntSetEquals(t, []int{}, MakeSet[int]().Intersect(MakeSet(1)), "intersect empty set")
}

func TestSet_GetArbitraryElement(t *testing.T) {
    elt := MakeSet(2, 3).GetArbitraryElement()
    assert.True(t, elt == 2 || elt == 3, "Element must be from the given set")
    assert.Panics(t, func() {
        Set[int]{}.GetArbitraryElement()
    }, "Reading an element from an empty set panics")
}

func TestSet_IsSubsetOf(t *testing.T) {
    assert.True(t, MakeSet(1, 2).IsSubsetOf(MakeSet(1, 2, 3)), "simple subset")
    assert.True(t, MakeSet[int]().IsSubsetOf(MakeSet(1, 2, 3)), "empty subset")
    assert.False(t, MakeSet(4).IsSubsetOf(MakeSet(1, 2, 3)), "not subset")
    assert.False(t, MakeSet(4).IsSubsetOf(MakeSet[int]()), "not subset of empty set")
    assert.True(t, MakeSet[int]().IsSubsetOf(MakeSet[int]()), "empty set is subset of empty set")
}

func TestSet_IsSupersetOf(t *testing.T) {
    assert.True(t, MakeSet(1, 2, 3).IsSupersetOf(MakeSet(1, 2)), "simple superset")
    assert.True(t, MakeSet(1, 2).IsSupersetOf(MakeSet[int]()), "superset of empty set")
    assert.False(t, MakeSet(1, 2).IsSupersetOf(MakeSet(4)), "not superset")
    assert.False(t, MakeSet[int]().IsSupersetOf(MakeSet(4)), "empty set not superset")
    assert.True(t, MakeSet[int]().IsSupersetOf(MakeSet[int]()), "empty set is superset of empty set")
}

func TestSet_SubtractSet(t *testing.T) {
    assertIntSetEquals(t, []int{1, 2}, MakeSet(1, 2, 3, 4).SubtractSet(MakeSet(3, 4, 5)), "simple subtraction")
    assertIntSetEquals(t, []int{1, 2}, MakeSet(1, 2).SubtractSet(MakeSet[int]()), "subtracting an empty set")
    assertIntSetEquals(t, []int{}, MakeSet[int]().SubtractSet(MakeSet(1, 2)), "subtraction of empty set")
}

func TestIntersection(t *testing.T) {
    s1 := MakeSet(1, 2)
    s2 := MakeSet(2, 3)
    assertIntSetEquals(t, []int{2}, Intersection(s1, s2), "simple intersection")
    assertIntSetEquals(t, []int{1, 2}, s1, "s1 was not changed")
    assertIntSetEquals(t, []int{2, 3}, s2, "s2 was not changed")
    assertIntSetEquals(t, []int{}, Intersection(MakeSet(1, 2), Set[int]{}), "intersect with empty set")
    assertIntSetEquals(t, []int{}, Intersection(Set[int]{}, Set[int]{}), "intersect two empty sets")
}

func TestUnion(t *testing.T) {
    s1 := MakeSet(1, 2)
    s2 := MakeSet(2, 3)
    assertIntSetEquals(t, []int{1, 2, 3}, Union(s1, s2), "simple intersection")
    assertIntSetEquals(t, []int{1, 2}, s1, "s1 was not changed")
    assertIntSetEquals(t, []int{2, 3}, s2, "s2 was not changed")
}

func TestSetDifference(t *testing.T) {
    s1 := MakeSet(1, 2)
    s2 := MakeSet(2, 3)
    assertIntSetEquals(t, []int{1}, SetDifference(s1, s2), "simple intersection")
    assertIntSetEquals(t, []int{1, 2}, s1, "s1 was not changed")
    assertIntSetEquals(t, []int{2, 3}, s2, "s2 was not changed")
}

func TestSet_ToSlice(t *testing.T) {
    slc := MakeSet(1, 2).ToSlice()
    sort.Ints(slc)
    assert.Equal(t, []int{1, 2}, slc, "Getting data as slice")
    assert.Equal(t, []int{}, Set[int]{}.ToSlice(), "Getting empty set data as slice")
}

func TestMapSet(t *testing.T) {
    assertIntSetEquals(t, []int{4, 9}, MapSet(MakeSet(2, 3, -2), func(a int) int { return a * a }), "Simple map set")
    assertIntSetEquals(t, []int{}, MapSet(Set[int]{}, func(a int) int { return a * a }), "mapping an empty set")
}

func TestSet_Some_Every(t *testing.T) {
    assert.Equal(t, false, MakeSet(2, 3, 4).Some(func(n int) bool { return n < 1 }))
    assert.Equal(t, true, MakeSet(2, 0, 4).Some(func(n int) bool { return n < 1 }))
    assert.Equal(t, false, MakeSet[int]().Some(func(n int) bool { return n < 1 }))
    assert.Equal(t, true, MakeSet(2, 3, 4).Every(func(n int) bool { return n > 1 }))
    assert.Equal(t, false, MakeSet(2, 0, 4).Every(func(n int) bool { return n > 1 }))
    assert.Equal(t, true, MakeSet[int]().Every(func(n int) bool { return n > 1 }))
}

func collectSet[C cmp.Ordered](s Set[C], n int) []C {
    result := make([]C, min(len(s), n))
    i := 0
    for elt := range s.Iter() {
        if i == n {
            break
        }
        result[i] = elt
        i++
    }
    slices.Sort(result)
    return result
}

func TestSet_Iter(t *testing.T) {
    assert.Equal(t, []int{1, 2, 4}, collectSet(MakeSet(1, 2, 4, 2), 100))
    assert.Equal(t, 2, len(collectSet(MakeSet(1, 2, 4, 3), 2)), "ensure break works")
    assert.Equal(t, []int{1}, collectSet(MakeSet(1), 100))
    assert.Equal(t, []int{}, collectSet(MakeSet[int](), 100))
    assert.Equal(t, []string{"a", "ab"}, collectSet(MakeSet("ab", "a"), 100))
}

func TestSet_Len(t *testing.T) {
    assert.Equal(t, 2, MakeSet(1, 2).Len())
    assert.Equal(t, 0, MakeSet[int]().Len())
    var nilSet Set[int]
    assert.Equal(t, 0, nilSet.Len())
    assert.Equal(t, 3, MakeSet("a", "b", "c").Len())
}
