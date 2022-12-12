package collection

import v "github.com/nealmcc/aoc2022/pkg/vector/twod"

// Slice2D is a generic square grid of items with O(1) access times.
//
// It prioritises very fast access times, but allocates memory as if
// all grid positions are used, and does not support negative X and Y values.
//
// Just like a slice, will panic if an invalid index is requested.
type Slice2D[T any] struct {
	size int
	data []T
}

// NewSlice2D initialises a new collection with the given width.
func NewSlice2D[T any](size int) *Slice2D[T] {
	return &Slice2D[T]{
		size: size,
		data: make([]T, size*size),
	}
}

// Set the given point to the given value.
func (s *Slice2D[T]) Set(p v.Point, x T) {
	s.data[s.indexOf(p)] = x
}

// Get the value at the given point.
func (s *Slice2D[T]) Get(p v.Point) T {
	return s.data[s.indexOf(p)]
}

// Neighbours4 returns a slice of all positions within the collection that
// are adjacent to the given position.
// Squares are only adjacent vertically and horizontally - not diagonally.
func (s *Slice2D[T]) Neighbours4(pos v.Point) []v.Point {
	points := make([]v.Point, 0, 4)
	for _, p := range pos.Neighbours4() {
		if p.X < 0 || p.X >= s.size || p.Y < 0 || p.Y >= s.size {
			continue
		}
		points = append(points, p)
	}
	return points
}

// indexOf returns the linear index that corresponds to the given 2D point.
//
// values are stored in top to bottom, left to right order, where (0, 0) is in
// the top left, and (x=size-1, y=size-1) is in the bottom right.
//
// For example, with a Slice2D with a size of 3:
//    0 1 2
//    3 4 5
//    6 7 8
//
func (s *Slice2D[T]) indexOf(p v.Point) int {
	return p.Y*s.size + p.X
}
