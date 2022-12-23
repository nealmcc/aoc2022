package main

import (
	"bytes"
	"fmt"

	"github.com/nealmcc/aoc2022/pkg/bound"
	v "github.com/nealmcc/aoc2022/pkg/vector/twod"
)

// Forest is a two dimensional grid containing elves
type Forest struct {
	extents bound.Rect
	Grid    map[Elf]struct{}
}

// Elf is a v.Point with additional methods to look adjacent to itself.
type Elf struct {
	v.Point
}

// Tick processes a full round of tree planting
// The given directions determine the sequence of directions that the elves will
// use during this round when looking for a destination.
func (f *Forest) Tick(dirs [4]Facing) bool {
	next := make(map[Elf]struct{}, len(f.Grid))
	// proposals maps from a destination to a list of elves that want to move there.
	proposals := make(map[v.Point][]Elf)
	for elf := range f.Grid {
		if elf.IsAlone(*f) {
			next[elf] = struct{}{}
			continue
		}
		if dest, ok := elf.Survey(*f, dirs); ok {
			proposals[dest] = append(proposals[dest], elf)
		} else {
			next[elf] = struct{}{}
		}
	}
	for p, list := range proposals {
		if len(list) == 1 {
			next[Elf{p}] = struct{}{}
		} else {
			for _, elf := range list {
				next[elf] = struct{}{}
			}
		}
	}
	f.Grid = next
	return len(proposals) == 0
}

func (f *Forest) setBounds() {
	const large = 1<<63 - 1

	b := bound.Rect{
		Min: v.Point{X: large, Y: large},
		Max: v.Point{X: -1 * large, Y: -1 * large},
	}

	for p := range f.Grid {
		if p.X < b.Min.X {
			b.Min.X = p.X
		}
		if p.X > b.Max.X {
			b.Max.X = p.X
		}

		if p.Y < b.Min.Y {
			b.Min.Y = p.Y
		}
		if p.Y > b.Max.Y {
			b.Max.Y = p.Y
		}
	}
	f.extents = b
}

func (f Forest) CountEmpty() int {
	count := 0
	f.setBounds()
	min := f.extents.Min
	max := f.extents.Max
	for x := min.X; x <= max.X; x++ {
		for y := min.Y; y <= max.Y; y++ {
			if _, ok := f.Grid[Elf{v.Point{X: x, Y: y}}]; !ok {
				count++
			}
		}
	}
	return count
}

// Format implements fmt.Formatter.
// The width is used to increase padding around the forest if desired.
func (f Forest) Format(s fmt.State, verb rune) {
	width, ok := s.Width()

	size := f.extents.Size()

	if !ok || width < size.X {
		width = size.X
	}

	pad := (width - size.X) / 2
	height := size.Y + 2*pad

	lines := make([][]byte, height)
	for row, y := 0, -1*pad+f.extents.Min.Y; row < height; row, y = row+1, y+1 {
		buf := make([]byte, width)
		for col, x := 0, -1*pad+f.extents.Min.X; col < width; col, x = col+1, x+1 {
			if _, ok := f.Grid[Elf{v.Point{X: x, Y: y}}]; ok {
				buf[col] = '#'
			} else {
				buf[col] = '.'
			}
		}
		lines[row] = buf
	}
	buf := bytes.Join(lines, []byte("\n"))

	s.Write(buf)
}

// String implements fmt.Stringer.
func (f Forest) String() string {
	return fmt.Sprintf("%v", f)
}

type Facing int

const (
	North Facing = iota
	South
	West
	East
)

func (f Facing) String() string {
	return [...]string{
		"N",
		"S",
		"W",
		"E",
	}[f]
}

func (f Facing) AsVector() v.Point {
	return [...]v.Point{
		{Y: -1},
		{Y: 1},
		{X: -1},
		{X: 1},
	}[f]
}

func dirSequence(n int) [4]Facing {
	res := [4]Facing{}
	for i := 0; i < 4; i++ {
		res[i] = Facing((n + i) % 4)
	}
	return res
}

// IsAlone asks the elf if there are any other elves beside it.
func (e Elf) IsAlone(f Forest) bool {
	for _, p := range e.Neighbours8() {
		if _, ok := f.Grid[Elf{p}]; ok {
			return false
		}
	}
	return true
}

// Survey asks the elf to look around it in each direction in turn and look
// for a good place to move to. If the elf sees a good place, it returns it.
// If there is no suitable place then the elf returns false.
func (e Elf) Survey(f Forest, dirs [4]Facing) (v.Point, bool) {
dirloop:
	for _, d := range dirs {
		for _, p := range e.Adjacent3(d) {
			if _, isFull := f.Grid[Elf{p}]; isFull {
				continue dirloop
			}
		}
		next := e.Add(d.AsVector())
		return next, true
	}
	return e.Point, false
}

// Neighbours8 returns the eight points adjacent to this elf.
func (e Elf) Neighbours8() []v.Point {
	points := make([]v.Point, 0, 8)
	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			if x == 0 && y == 0 {
				continue
			}
			points = append(points, e.Add(v.Point{X: x, Y: y}))
		}
	}
	return points
}

// Adjacent3 returns the three points adjacent to this elf in the given direction
func (e Elf) Adjacent3(dir Facing) [3]v.Point {
	return [...][3]v.Point{
		// North
		{e.add(-1, -1), e.add(0, -1), e.add(1, -1)},
		// South
		{e.add(-1, 1), e.add(0, 1), e.add(1, 1)},
		// West
		{e.add(-1, -1), e.add(-1, 0), e.add(-1, 1)},
		// East
		{e.add(1, -1), e.add(1, 0), e.add(1, 1)},
	}[dir]
}

// add is a utility function that makes it easier to read
func (e Elf) add(x, y int) v.Point {
	return e.Add(v.Point{X: x, Y: y})
}
