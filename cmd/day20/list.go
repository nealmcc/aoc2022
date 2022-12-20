package main

import (
	"container/ring"
)

// List is a circular, doubly-linked list.
type List struct {
	Root     *ring.Ring   // a pointer to the marker
	offset   int          // the index of the root, within the initial input
	pointers []*ring.Ring // allows for indexed lookup, based on the input sequence
}

// NewList initializes a new List using the given sequence of inputs and marker.
// The marker is assumed to be unique within the input.
func NewList(nums []int, marker int) List {
	size := len(nums)
	l := List{
		pointers: make([]*ring.Ring, size),
	}

	r := ring.New(size)
	for i, n := range nums {
		r.Value = n
		l.pointers[i] = r
		if n == marker {
			l.Root = r
			l.offset = i
		}
		r = r.Next()
	}
	return l
}

// Coordinates examines the list to find the 3 values as specified in part 1,
// and adds them up.
func (l List) Coordinates() int {
	sum := 0
	r := l.Root
	for i := 0; i < 3; i++ {
		r = r.Move(1000 % len(l.pointers))
		sum += r.Value.(int)
	}
	return sum
}

// Mix shuffles the elements in the list according to the algorithm in part 1.
func (l List) Mix() {
	for _, p := range l.pointers {
		n := p.Value.(int) % (l.Len() - 1)
		r := p.Move(-1)
		r.Unlink(1)
		r = r.Move(n)
		r.Link(p)
	}
}

// Fixed returns the values from the list in fixed form - that is,
// as if the marker is immobile at the same place it started.
func (l List) Fixed() []int {
	fixed := make([]int, len(l.pointers))
	i := l.offset
	l.Root.Do(func(val any) {
		fixed[i%len(l.pointers)] = val.(int)
		i++
	})
	return fixed
}

// Normal returns the values from the list in normal form - that is,
// the first element will be the marker followed by the rest.
func (l List) Normal() []int {
	normal := make([]int, len(l.pointers))
	i := 0
	l.Root.Do(func(val any) {
		normal[i] = val.(int)
		i++
	})
	return normal
}

// Len returns the number of items in the list
func (l List) Len() int {
	return len(l.pointers)
}
