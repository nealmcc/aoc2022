package main

import "fmt"

// Pos is a position on a 2d grid
type Pos struct {
	Row int
	Col int
}

// Direction is a bitmask for one or more of N, E, S, W
type Direction byte

const (
	top    Direction = 1 << iota // 1
	right                        // 2
	bottom                       // 4
	left                         // 8
)

// Mask is a 2d grid of bitwise masks showing which directions each tree is visible from.
type Mask map[Pos]Direction

func (m Mask) Filter(dir Direction) Mask {
	out := make(Mask)
	for coord, val := range m {
		if val&dir > 0 {
			out[coord] = dir
		}
	}
	return out
}

func (m Mask) size() (int, int) {
	width, height := 0, 0
	for k := range m {
		if k.Col > width {
			width = k.Col
		}
		if k.Row > height {
			height = k.Row
		}
	}
	return width, height
}

// ToSlice converts the sparsely populated bitmask to a fully populated matrix.
func (m Mask) ToSlice() [][]Direction {
	w, h := m.size()
	data := make([][]Direction, h)
	for row := 0; row < w; row++ {
		data[row] = make([]Direction, w)
	}

	for pos, val := range m {
		data[pos.Row][pos.Col] = val
	}

	return data
}

// Format implements fmt.Formatter.
// The width is used to increase padding on the left and right if desired.
func (m Mask) Format(s fmt.State, verb rune) {
	width, ok := s.Width()
	w, height := m.size()
	if !ok || width < w {
		width = w
	}

	buf := make([]byte, width+1)
	pad := (width - w) / 2
	for row := 0; row <= height; row++ {
		for col := -1 * pad; col <= width+pad; col++ {
			buf[col+pad] = fmt.Sprintf("%x", m[Pos{row, col}])[0]
		}
		s.Write(append(buf, '\n'))
	}
}

// String implements fmt.Stringer.
func (m Mask) String() string {
	return fmt.Sprintf("%v", m)
}
