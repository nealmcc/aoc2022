package main

import (
	"bytes"
	"fmt"

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

	lines[st.entry.Y+1][st.entry.X+1] = '.'
	lines[st.exit.Y+1][st.exit.X+1] = '.'

	// for debugging
	// meta := fmt.Sprintf("%v, %v\n", st.extents, st.extents.Size())
	// meta += fmt.Sprintf("%v\n", st.entry)
	// meta += fmt.Sprintf("%v\n", st.exit)
	// meta += fmt.Sprintf("%v\n", st.Grid)

	return string(bytes.Join(lines, []byte("\n")))
}

// IceAt looks at the given position, at some time t, to determine
// what ice will be there, if any.
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
		fmt.Println("looking at", p, "for", ice)
		incoming := st.grid[p]
		result |= incoming & ice
	}

	return result, true
}
