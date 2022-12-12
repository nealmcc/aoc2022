package main

import (
	"bufio"
	"container/heap"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/nealmcc/aoc2022/pkg/collection"
	v "github.com/nealmcc/aoc2022/pkg/vector/twod"
	pq "github.com/nealmcc/aoc2022/pkg/vector/twod/priorityqueue"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	hill, err := read(file)
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()

	p1 := part1(hill)
	middle := time.Now()
	p2 := part2(hill)
	end := time.Now()

	fmt.Printf("part 1: %d in %s\n", p1, middle.Sub(start))
	fmt.Printf("part 2: %d in %s\n", p2, end.Sub(middle))
}

// read the terrain from the given input.
func read(r io.Reader) (grid, error) {
	s := bufio.NewScanner(r)

	if !s.Scan() {
		return grid{}, errors.New("empty input")
	}
	if err := s.Err(); err != nil {
		return grid{}, err
	}

	row := s.Bytes()
	hill := newGrid(len(row))

	for x := 0; x < len(row); x++ {
		if err := hill.Set(v.Point{X: x, Y: 0}, row[x]); err != nil {
			return grid{}, err
		}
	}

	for y := 1; s.Scan(); y++ {
		row := s.Bytes()
		for x := 0; x < len(row); x++ {
			if err := hill.Set(v.Point{X: x, Y: y}, row[x]); err != nil {
				return grid{}, err
			}
		}
	}

	if err := s.Err(); err != nil {
		return grid{}, err
	}

	return hill, nil
}

// grid is a square topographical map with coordinates
// ranging from (0, 0) at the top left to (size-1, size-1) at the bottom right.
// Each position on the map has an elevation from 'a' (lowest) to 'z' (highest).
type grid struct {
	size    int
	start   v.Point
	end     v.Point
	terrain *collection.Slice2D[byte]
}

func newGrid(width int) grid {
	return grid{
		size:    width,
		terrain: collection.NewSlice2D[byte](width),
	}
}

// Set the elevation of the given coordinate.
func (g *grid) Set(p v.Point, b byte) error {
	switch true {
	case b == 'S':
		g.start = p
		g.terrain.Set(p, 'a')

	case b == 'E':
		g.end = p
		g.terrain.Set(p, 'z')

	case 'a' <= b && b <= 'z':
		g.terrain.Set(p, b)

	default:
		return fmt.Errorf("invalid elevation: %v", b)
	}
	return nil
}

const _infinity = 1<<63 - 1

// part1 computes the shortest distance from the grid's start point to end point,
// assuming you an only climb a maximum of 1 height per step.
// Uses Dijkstra's Algorithm.
func part1(g grid) int {
	// save the shortest distance to each point:
	dist := collection.NewSlice2D[int](g.size)

	// save a pointer to each in the queue, so we can update their priorities:
	pointers := collection.NewSlice2D[*pq.Node](g.size)

	// push each node on the graph into the queue, with an initial distance
	q := new(pq.Queue)

	for y := 0; y < g.size; y++ {
		for x := 0; x < g.size; x++ {
			coord := v.Point{X: x, Y: y}
			if coord != g.start {
				dist.Set(coord, _infinity)
			}
			node := &pq.Node{
				Value:    coord,
				Priority: -1 * dist.Get(coord),
			}
			pointers.Set(coord, node)
			heap.Push(q, node)
		}
	}

	// the first node that we definitely know the distance to is the start.
	// It has the highest priority (0 - all others are negative) so we pop it
	// off the queue, and find the distance to all adjacent nodes.
	// We keep working out from the next nearest node until all nodes have a defined distance.
	for q.Len() > 0 {
		n := heap.Pop(q).(*pq.Node)
		curr, currTotal := n.Value, n.Priority*-1

		for _, next := range g.terrain.Neighbours4(curr) {
			currHeight := g.terrain.Get(curr)
			nextHeight := g.terrain.Get(next)

			if nextHeight > currHeight+1 {
				// too steep - try another way.
				continue
			}

			// the distance from start to next, if we arrive via curr:
			alt := currTotal + 1
			if alt < dist.Get(next) {
				dist.Set(next, alt)
				p := pointers.Get(next)
				q.Update(p, p.Value, -1*alt)
			}
		}
	}

	return dist.Get(g.end)
}

// part2 computes the shortest distance from any point with elevation 'a',
// to the grid's end point, assuming you an only climb 1 height per step.
// Uses Dijkstra's Algorithm.
func part2(g grid) int {
	dist := collection.NewSlice2D[int](g.size)
	pointers := collection.NewSlice2D[*pq.Node](g.size)

	// push each node on the graph into the queue, with an initial distance
	q := new(pq.Queue)

	for y := 0; y < g.size; y++ {
		for x := 0; x < g.size; x++ {
			coord := v.Point{X: x, Y: y}
			if g.terrain.Get(coord) != 'a' {
				dist.Set(coord, _infinity)
			}
			node := &pq.Node{
				Value:    coord,
				Priority: -1 * dist.Get(coord),
			}
			pointers.Set(coord, node)
			heap.Push(q, node)
		}
	}

	// Work out from the next nearest node until all nodes have a defined cost.
	for q.Len() > 0 {
		n := heap.Pop(q).(*pq.Node)
		curr, currTotal := n.Value, n.Priority*-1

		for _, next := range g.terrain.Neighbours4(curr) {
			currHeight := g.terrain.Get(curr)
			nextHeight := g.terrain.Get(next)

			if nextHeight > currHeight+1 {
				// too steep - try another way.
				continue
			}

			// the distance from start to next, if we arrive via curr:
			alt := currTotal + 1
			if alt < dist.Get(next) {
				dist.Set(next, alt)
				p := pointers.Get(next)
				q.Update(p, p.Value, -1*alt)
			}
		}
	}

	return dist.Get(g.end)
}
