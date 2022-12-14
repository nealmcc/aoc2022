package main

import (
	v "github.com/nealmcc/aoc2022/pkg/vector/twod"
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
func (c *Cavern) Get(p v.Point) (Material, bool) {
	m, ok := c.grid[p.X][p.Y]
	return m, ok
}

// Place the given material into the cavern.
func (c *Cavern) Set(p v.Point, m Material) {
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
func (c *Cavern) DropSand(src v.Point) (v.Point, bool) {
	p := src
	var (
		down      = v.Point{Y: 1}
		downleft  = v.Point{X: -1, Y: 1}
		downright = v.Point{X: 1, Y: 1}
	)

	for {
		if p.Y >= c.lowest {
			return v.Point{}, false
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
