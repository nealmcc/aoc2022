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

func (f Forest) Origin() v.Point {
	return v.Point{
		X: f.boundsHoriz[0].Min,
		Y: 0,
	}
}

// wrapFunc is a function that examines the current position and facing,
// and returns the next position and facing if the traveller were
// to take that step.  Does not move the traveller.
type wrapFunc func(pos v.Point, f Facing) (v.Point, Facing)

// Next determines the next coordinate and facing the traveller moves to
func (f *Forest) Next(curr v.Point, dist int, dir Facing, wrapFn wrapFunc) (v.Point, Facing) {
	count := 0
	defer func(prev v.Point) {
		fmt.Printf("moved %d %s from %v, got to %s facing %s\n",
			count, dir, prev, curr, dir)
	}(curr)

	var next v.Point
	for i := 0; i < dist; i++ {
		next, dir = wrapFn(curr, dir)
		if sq := f.grid[next]; sq == '#' {
			fmt.Println("hit a tree")
			return curr, dir
		}
		curr, count = next, count+1
	}

	return curr, dir
}

func (f Forest) wrap1(pos v.Point, dir Facing) (v.Point, Facing) {
	delta := dir.AsVector()
	next := pos.Add(delta)

	isVertical := dir == Up || dir == Down

	if isVertical {
		bound := f.boundsVert[pos.X]
		next.Y = bound.Mod(next.Y)
		if next != pos.Add(delta) {
			fmt.Printf("wrap1 adjusted for vertical bounds at %v moving %v\n", pos, dir)
		}
		return next, dir
	}

	bound := f.boundsHoriz[pos.Y]
	next.X = bound.Mod(next.X)
	if next != pos.Add(delta) {
		fmt.Printf("wrap1 adjusted for horizontal bounds at %v moving %v\n", pos, delta)
	}
	return next, dir
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

// Bound defines a lower and upper bound (inclusive at both ends)
type Bound struct {
	Min int
	Max int
}

// Len returns the size of this boundary
func (b Bound) Len() int {
	return b.Max - b.Min + 1
}

// Mod returns the given value adjusted to fit within this boundary
func (b Bound) Mod(n int) int {
	// mod return a mod b.
	// (% is the remainder operator)
	mod := func(a, b int) int {
		return (a%b + b) % b
	}
	return b.Min + mod(n-b.Min, b.Len())
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

type zone int

const (
	ZOne zone = iota + 1
	ZTwo
	ZThree
	ZFour
	ZFive
	ZSix
)

func (f Forest) wrap2(pos v.Point, dir Facing) (v.Point, Facing) {
	return pos, dir
}

// func (f Forest) zoneWrappers(p v.Point) (z Zone, h, v wrapFunc) {
// 	if p.X >= 100 {
// 		// zone two, we need a wrapper to move up or left
// 		h = f.wrapTwoHoriz
// 		v = f.wrapTwoVert
// 		return ZTwo, h, v
// 	}
// 	if p.Y < 50 {
// 		return ZOne
// 	}
// 	if p.Y >= 150 {
// 		return ZSix
// 	}
// 	if p.X < 50 {
// 		return ZFour
// 	}
// 	if p.Y >= 100 {
// 		return ZFive
// 	}
// 	return ZThree
// }

// func (f Forest) wrapTwoHoriz(pos, delta v.Point) (v.Point, Facing) {
// }
