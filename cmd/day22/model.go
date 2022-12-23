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
func (f *Forest) Next(curr v.Point, dist int, dir Facing, wrapFn wrapFunc) (pNext v.Point, dirNext Facing) {
	count := 0
	defer func(start v.Point, dir Facing) {
		fmt.Printf("moved %d %s from %v, got to %s facing %s\n",
			count, dir, start, pNext, dirNext)
	}(curr, dir)

	var next v.Point
	for i := 0; i < dist; i++ {
		next, dirNext = wrapFn(curr, dir)
		if sq := f.grid[next]; sq == '#' {
			fmt.Println("hit a tree")
			return curr, dir
		}
		curr, dir, count = next, dirNext, count+1
	}

	return curr, dir
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
	// return a mod b.
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

type Zone int

const (
	ZOne Zone = iota + 1
	ZTwo
	ZThree
	ZFour
	ZFive
	ZSix
)

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

func (f Forest) wrap2(pos v.Point, dir Facing) (pNext v.Point, dirNext Facing) {
	z, fn := f.getZone(pos)
	defer func() {
		if pos.Add(dir.AsVector()) != pNext {
			zNext, _ := f.getZone(pNext)
			fmt.Printf("left zone %d at %s moving %s ; now in zone %d at %v moving %s\n",
				z, pos, dir, zNext, pNext, dirNext)
		}
	}()

	pNext, dirNext = fn(pos, dir)

	return
}

func (f Forest) getZone(p v.Point) (z Zone, wrapFn wrapFunc) {
	switch {
	case p.X >= 100:
		return ZTwo, zTwoWrap

	case p.Y < 50:
		return ZOne, zOneWrap

	case p.Y >= 150:
		return ZSix, zSixWrap

	case p.X < 50:
		return ZFour, zFourWrap

	case p.Y >= 100:
		return ZFive, zFiveWrap

	default:
		return ZThree, zThreeWrap
	}
}

// zOneWrap is the wrapper function to use while in zone one.
func zOneWrap(p v.Point, dir Facing) (v.Point, Facing) {
	switch {
	case dir == Up && p.Y == 0:
		// ✓ up from z1 -> right into z6
		return v.Point{
			X: 0,
			Y: 100 + p.X,
		}, Right

	case dir == Left && p.X == 50:
		// ✓ left from z1 -> right into z4
		return v.Point{
			X: 0,
			Y: 149 - p.Y,
		}, Right

	default:
		return p.Add(dir.AsVector()), dir
	}
}

// zTwoWrap is the wrapper function to use while in zone two.
func zTwoWrap(p v.Point, dir Facing) (v.Point, Facing) {
	switch {
	case dir == Up && p.Y == 0:
		// ✓ up from z2 -> up into z6
		return v.Point{
			X: p.X - 100,
			Y: 199,
		}, Up

	case dir == Right && p.X == 149:
		// ✓ right from z2 -> left into z5
		return v.Point{
			X: 99,
			Y: 149 - p.Y,
		}, Left

	case dir == Down && p.Y == 49:
		// ✓ down from z2 -> left into z3
		return v.Point{
			X: 99,
			Y: p.X - 50,
		}, Left

	default:
		return p.Add(dir.AsVector()), dir
	}
}

// zThreeWrap is the wrapper function to use while in zone three.
func zThreeWrap(p v.Point, dir Facing) (v.Point, Facing) {
	switch {
	case dir == Right && p.X == 99:
		// ✓ right from z3 -> up into z2
		return v.Point{
			X: p.Y + 50,
			Y: 49,
		}, Up

	case dir == Left && p.X == 50:
		// ✓ left from z3 -> down into z4
		return v.Point{
			X: p.Y - 50,
			Y: 100,
		}, Down

	default:
		return p.Add(dir.AsVector()), dir
	}
}

// zFourWrap is the wrapper function to use while in zone four.
func zFourWrap(p v.Point, dir Facing) (v.Point, Facing) {
	switch {
	case dir == Up && p.Y == 100:
		// ✓ up from z4 -> right into z3
		return v.Point{
			X: 50,
			Y: p.X + 50,
		}, Right

	case dir == Left && p.X == 0:
		// ✓ left from z4 -> right into z1
		return v.Point{
			X: 50,
			Y: 149 - p.Y,
		}, Right

	default:
		return p.Add(dir.AsVector()), dir
	}
}

// zFiveWrap is the wrapper function to use while in zone five.
func zFiveWrap(p v.Point, dir Facing) (v.Point, Facing) {
	switch {
	case dir == Right && p.X == 99:
		// ✓ right from z5 -> left into z2
		return v.Point{
			X: 149,
			Y: 149 - p.Y,
		}, Left

	case dir == Down && p.Y == 149:
		// ✓ down from z5 -> left into z6
		return v.Point{
			X: 49,
			Y: 100 + p.X,
		}, Left

	default:
		return p.Add(dir.AsVector()), dir
	}
}

// zSixWrap is the wrapper function to use while in zone siz.
func zSixWrap(p v.Point, dir Facing) (v.Point, Facing) {
	switch {
	case dir == Right && p.X == 49:
		// ✓ right from z6 -> up into z5
		return v.Point{
			X: p.Y - 100,
			Y: 149,
		}, Up

	case dir == Down && p.Y == 199:
		// ✓ down from z6 -> down into z2
		return v.Point{
			X: 100 + p.X,
			Y: 0,
		}, Down

	case dir == Left && p.X == 0:
		// ✓ left from z6 -> down into z1
		return v.Point{
			X: p.Y - 100,
			Y: 0,
		}, Down

	default:
		return p.Add(dir.AsVector()), dir
	}
}

// ✓
