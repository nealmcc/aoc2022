package main

import (
	"fmt"

	v "github.com/nealmcc/aoc2022/pkg/vector/twod"
)

type Forest struct {
	grid        map[v.Point]byte
	width       int
	height      int
	boundsHoriz []Bound // for each y coordinate, the smallest and largest X value
	boundsVert  []Bound // for each x coordinate, the smallest and largest Y value
}

// Bound defines a lower and upper bound (inclusive at both ends)
type Bound struct {
	Min int
	Max int
}

func (f Forest) Origin() v.Point {
	return v.Point{
		X: f.boundsHoriz[0].Min,
		Y: 0,
	}
}

// NextPart1 determines the next coordinate in the forest the traveller
// can get to, based on the given distance and facing.
func (f *Forest) NextPart1(curr v.Point, dist int, dir Facing) v.Point {
	count := 0
	defer func(prev v.Point) {
		fmt.Printf("moved %d %s from %v, got to %s\n", count, dir, prev, curr)
	}(curr)

	vect := dir.AsVector()
	for i := 0; i < dist; i++ {
		next := f.wrap(curr, vect)
		if square := f.grid[next]; square == '#' {
			fmt.Println("hit a tree")
			return curr
		}
		curr, count = next, count+1
	}
	return curr
}

func (f Forest) wrap(pos, delta v.Point) v.Point {
	mod := func(a, b int) int {
		return (a%b + b) % b
	}

	next := pos.Add(delta)
	if delta.Y != 0 {
		bound := f.boundsVert[pos.X]
		size := bound.Max - bound.Min + 1
		next.Y = bound.Min + mod(next.Y-bound.Min, size)
		if next != pos.Add(delta) {
			fmt.Printf("wrap adjusted for vertical bounds at %v moving %v\n", pos, delta)
		}
		return next
	}

	bound := f.boundsHoriz[pos.Y]
	size := bound.Max - bound.Min + 1
	next.X = bound.Min + mod(next.X-bound.Min, size)
	if next != pos.Add(delta) {
		fmt.Printf("wrap adjusted for horizontal bounds at %v moving %v\n", pos, delta)
	}
	return next
}

func (f *Forest) setBounds() {
	const large = 1<<63 - 1

	f.boundsVert = make([]Bound, f.width)
	for x := 0; x < f.width; x++ {
		f.boundsVert[x] = Bound{
			Min: large,
			Max: -1,
		}
	}

	f.boundsHoriz = make([]Bound, f.height)
	for y := 0; y < f.height; y++ {
		f.boundsHoriz[y] = Bound{
			Min: large,
			Max: -1,
		}
	}

	for p := range f.grid {
		horiz := &f.boundsHoriz[p.Y]
		if p.X < horiz.Min {
			horiz.Min = p.X
		}
		if p.X > horiz.Max {
			horiz.Max = p.X
		}

		vert := &f.boundsVert[p.X]
		if p.Y < vert.Min {
			vert.Min = p.Y
		}
		if p.Y > vert.Max {
			vert.Max = p.Y
		}
	}
}

type Step struct {
	Dist     int  // the number of spaces to attempt to move forward
	Rotation byte // one of L=left, R=right
}

type Facing int

const (
	Right Facing = iota
	Down
	Left
	Up
)

func (f Facing) String() string {
	return [...]string{
		"Right",
		"Down",
		"Left",
		"Up",
	}[f]
}

func (f Facing) AsVector() v.Point {
	return [...]v.Point{
		{X: 1},
		{Y: 1},
		{X: -1},
		{Y: -1},
	}[f]
}
