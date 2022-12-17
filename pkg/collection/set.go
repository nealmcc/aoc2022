package collection

import (
	"fmt"
	"sort"
)

// Set implements a generic collection of unique elements.
// Not safe for concurrent use.
type Set[T comparable] map[T]struct{}

func NewSet[T comparable](vals ...T) Set[T] {
	s := make(Set[T])
	s.Init(vals)
	return s
}

// Init initialises this set with the given values.
// Any previous values will be lost.
func (s *Set[T]) Init(values []T) {
	*s = make(Set[T], len(values))
	for _, k := range values {
		(*s)[k] = struct{}{}
	}
}

// Add the given element to the set.
func (s *Set[T]) Add(k T) {
	(*s)[k] = struct{}{}
}

// Remove the given element from the set.
func (s *Set[T]) Remove(k T) {
	delete((*s), k)
}

// Contains returns true iff this set contains the given key
func (s Set[T]) Contains(k T) bool {
	_, ok := s[k]
	return ok
}

// Union returns the union of A and B
func Union[T comparable](a, b Set[T]) Set[T] {
	res := make(Set[T], len(a))
	for k := range a {
		res[k] = struct{}{}
	}
	for k := range b {
		res[k] = struct{}{}
	}
	return res
}

// Difference returns A minus B
func Difference[T comparable](a, b Set[T]) Set[T] {
	res := make(Set[T], len(a))
	for k := range a {
		if _, ok := b[k]; !ok {
			res[k] = struct{}{}
		}
	}
	return res
}

// Intersect returns the intersection of A and B
func Intersect[T comparable](a, b Set[T]) Set[T] {
	if len(b) < len(a) {
		a, b = b, a
	}
	res := make(Set[T], len(a))
	for k := range a {
		if _, ok := b[k]; ok {
			res[k] = struct{}{}
		}
	}
	return res
}

// Format implements fmt.Formatter.
// It writes the elements of s in lexicographic order.
func (s Set[T]) Format(f fmt.State, verb rune) {
	if len(s) == 0 {
		f.Write([]byte("[]"))
		return
	}

	parts := make([]string, 0, len(s))
	for k := range s {
		parts = append(parts, fmt.Sprintf("%v", k))
	}
	sort.Strings(parts)

	f.Write([]byte{'['})

	for i := 0; i < len(parts)-1; i++ {
		f.Write([]byte(parts[i]))
		f.Write([]byte(", "))
	}
	f.Write([]byte(parts[len(parts)-1]))

	f.Write([]byte{']'})
}

// String implements fmt.Stringer.
func (s *Set[T]) String() string {
	return fmt.Sprint(s)
}
