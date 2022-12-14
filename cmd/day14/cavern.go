package main

import (
	"bytes"

	"github.com/nealmcc/aoc2022/pkg/vector/twod"
)

// Material represents the contents of a given square in a cavern.
type Material byte

const (
	Air Material = iota
	Sand
	Rock
)

// Cavern represents the two-dimensional slice of the cave behind the waterfall.
// Each square in the cavern has either zero or one material in it.
// In this cavern, the coordinates start at 0,0 at the top left,
// and increase to the right and down.
type Cavern struct {
	grid   map[int]map[int]Material
	lowest int
}

// Get the material at the given point, if any.
func (c *Cavern) Get(p twod.Point) (Material, bool) {
	m, ok := c.grid[p.X][p.Y]
	return m, ok
}

// Place the given material into the cavern.
func (c *Cavern) Set(p twod.Point, m Material) {
	if c.grid == nil {
		c.grid = make(map[int]map[int]Material, 8)
	}

	col, ok := c.grid[p.X]
	if !ok {
		col = make(map[int]Material, 8)
		c.grid[p.X] = col
	}

	if p.Y > c.lowest {
		c.lowest = p.Y
	}
	col[p.Y] = m
}

// DropSand drops sand into the cavern from the given point.
// The sand will fall until it (possibly) comes to rest.
// Returns the position where the sand comes to rest, or false if
// it will fall forever, or block the entry point.
func (c *Cavern) DropSand(src twod.Point) (twod.Point, bool) {
	p := src
	var (
		down      = twod.Point{Y: 1}
		downleft  = twod.Point{X: -1, Y: 1}
		downright = twod.Point{X: 1, Y: 1}
	)

	for {
		if p.Y >= c.lowest {
			return twod.Point{}, false
		}

		if _, full := c.Get(p.Add(down)); !full {
			p = p.Add(down)
			continue
		}
		if _, full := c.Get(p.Add(downleft)); !full {
			p = p.Add(downleft)
			continue
		}
		if _, full := c.Get(p.Add(downright)); !full {
			p = p.Add(downright)
			continue
		}
		break
	}

	c.Set(p, Sand)
	return p, p != src
}

// Count returns the total number of the given material in the cave.
// Note that the cave may have a bottomless pit below it, so it only makes sense
// to count sand and rock.
func (c *Cavern) Count(m Material) int {
	sum := 0
	for _, column := range c.grid {
		for _, sq := range column {
			if sq == m {
				sum++
			}
		}
	}
	return sum
}

// Render the given portion of this cavern as a1 string.
// The portion that will be rendered will include (x1, y1) at the top left,
// and will *exclude* (x2, y2) in the bottom right
func (c *Cavern) Render(x1, y1, x2, y2 int) string {
	height := y2 - y1
	width := x2 - x1
	rows := make([][]byte, height)
	for r := 0; r < height; r++ {
		row := make([]byte, width)
		for k := 0; k < width; k++ {
			switch mat := c.grid[x1+k][y1+r]; mat {
			case Air:
				row[k] = '.'
			case Sand:
				row[k] = 'o'
			case Rock:
				row[k] = '#'
			}
		}
		rows[r] = row
	}
	return string(bytes.Join(rows, []byte{'\n'}))
}
