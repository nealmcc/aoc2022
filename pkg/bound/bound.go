// Package bound provides utility functions for evaluating boundary conditions
package bound

import v "github.com/nealmcc/aoc2022/pkg/vector/twod"

// Linear defines a lower and upper bound (inclusive at both ends)
type Linear struct {
	Min int
	Max int
}

// Size returns the size of this boundary
func (b Linear) Size() int {
	return b.Max - b.Min + 1
}

// Contains checks to see if this boundary contains the given value.
func (b Linear) Contains(n int) bool {
	return b.Min <= n && n <= b.Max
}

// Mod returns the given value adjusted to fit within this boundary
func (b Linear) Mod(n int) int {
	return b.Min + mod(n-b.Min, b.Size())
}

// Rect defines a two-dimensional bounding box (inclusive at both ends)
type Rect struct {
	Min v.Point
	Max v.Point
}

// Size returns the size of this boundary
func (b Rect) Size() v.Point {
	return b.Max.Sub(b.Min).Add(v.Point{X: 1, Y: 1})
}

// Contains checks to see if this boundary contains the given value.
func (b Rect) Contains(p v.Point) bool {
	return b.Min.X <= p.X && p.X <= b.Max.X &&
		p.Y <= b.Min.Y && p.Y <= b.Max.Y
}

// Mod returns the given value adjusted to fit within this boundary
func (b Rect) Mod(p v.Point) v.Point {
	size := b.Size()
	x := b.Min.X + mod(p.X-b.Min.X, size.X)
	y := b.Min.Y + mod(p.Y-b.Min.Y, size.Y)
	return v.Point{X: x, Y: y}
}

// mod return a mod b.
// (% is the remainder operator)
func mod(a, b int) int {
	return (a%b + b) % b
}
