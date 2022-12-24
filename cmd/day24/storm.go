package main

import (
	"bytes"

	"github.com/nealmcc/aoc2022/pkg/bound"
	v "github.com/nealmcc/aoc2022/pkg/vector/twod"
)

type Storm struct {
	grid        map[v.Point]Ice // a map of the ice at time t = 0
	extents     bound.Rect      // the rectangle that contains the storm itself
	entry, exit v.Point         // the entry and exit points
}

// String implements fmt.Stringer.
// It draws the walls, even though they are not stored within the grid data.
func (st Storm) String() string {
	return string(bytes.Join(st.Render(), []byte("\n")))
}

// Render prepares an output buffer to display this storm.
// Additional symbols may be added to the buffer using the compose() function.
func (st Storm) Render() [][]byte {
	size := st.extents.Size()
	width := size.X + 2
	height := size.Y + 2

	lines := make([][]byte, height)

	for row, y := 0, -1; row < height; row, y = row+1, y+1 {
		buf := make([]byte, width)
		for col, x := 0, -1; col < width; col, x = col+1, x+1 {
			if ice, ok := st.grid[v.Point{X: x, Y: y}]; ok {
				buf[col] = ice.Render()
			} else {
				buf[col] = '#'
			}
		}
		lines[row] = buf
	}

	compose(lines, map[v.Point]byte{
		{X: st.entry.X + 1, Y: st.entry.Y + 1}: '.',
		{X: st.exit.X + 1, Y: st.exit.Y + 1}:   '.',
	})

	return lines
}

// compose adds the given points to the given lines buffer, so that
// we can add symbols to the displayed storm before converting it to a string.
//
// The lines buffer should be in row-major order, with Y increasing downwards.
func compose(lines [][]byte, points map[v.Point]byte) {
	for k, char := range points {
		lines[k.Y][k.X] = char
	}
}

// At returns a copy of this storm at time t.
func (st Storm) At(t int) Storm {
	storm2 := Storm{
		extents: st.extents,
		entry:   st.entry,
		exit:    st.exit,
		grid:    make(map[v.Point]Ice, len(st.grid)),
	}

	for pos := range st.grid {
		next, _ := st.IceAt(pos, t)
		storm2.grid[pos] = next
	}
	return storm2
}

// IceAt looks at the given position, at some time t, to determine
// what ice will be there.
// returns false if the point is not within the storm grid.
func (st Storm) IceAt(pos v.Point, t int) (Ice, bool) {
	_, ok := st.grid[pos]
	if !ok {
		return None, false
	}

	var result Ice

	for _, ice := range [...]Ice{North, East, South, West} {
		delta := ice.AsVector().Times(-1 * t)
		p := st.extents.Mod(pos.Add(delta))
		incoming := st.grid[p]
		result |= incoming & ice
	}

	return result, true
}
