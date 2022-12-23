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
	Grid    map[v.Point]struct{}
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
			if _, ok := f.Grid[v.Point{X: x, Y: y}]; ok {
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
