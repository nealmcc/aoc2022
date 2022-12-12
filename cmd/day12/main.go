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

	"github.com/nealmcc/aoc2022/pkg/vector/twod"
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
	// p2 := part2(hill)
	// end := time.Now()

	fmt.Printf("part 1: %d in %s\n", p1, middle.Sub(start))
	// fmt.Printf("part 2: %d in %s\n", p2, end.Sub(middle))
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
		if err := hill.Set(twod.Point{X: x, Y: 0}, row[x]); err != nil {
			return grid{}, err
		}
	}

	for y := 1; s.Scan(); y++ {
		row := s.Bytes()
		for x := 0; x < len(row); x++ {
			if err := hill.Set(twod.Point{X: x, Y: y}, row[x]); err != nil {
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
	start   twod.Point
	end     twod.Point
	terrain map[twod.Point]byte
}

func newGrid(width int) grid {
	return grid{
		size:    width,
		terrain: make(map[twod.Point]byte, width*width),
	}
}

// Set the elevation of the given coordinate.
func (g *grid) Set(p twod.Point, b byte) error {
	switch true {
	case b == 'S':
		g.start = p
		g.terrain[p] = 'a'

	case b == 'E':
		g.end = p
		g.terrain[p] = 'z'

	case 'a' <= b && b <= 'z':
		g.terrain[p] = b

	default:
		return fmt.Errorf("invalid elevation: %v", b)
	}
	return nil
}

// part1 computes the shortest distance from the grid's start point to end point,
// assuming you an only climb a maximum of 1 height per step.
// Uses Dijkstra's Algorithm.
func part1(g grid) int {
	dist := make(map[twod.Point]int, len(g.terrain))
	dist[g.start] = 0

	// push each node on the graph into the queue, with an initial distance
	q := new(pq.Queue)
	// save a pointer to each node, so we can update it in the queue
	pointers := make([]*pq.Node, g.size*g.size)
	const infinity = 1<<63 - 1
	for pos := range g.terrain {
		if pos != g.start {
			dist[pos] = infinity
		}
		node := &pq.Node{Value: pos, Priority: -1 * dist[pos]}
		pointers[pos.X*g.size+pos.Y] = node
		heap.Push(q, node)
	}

	// the first node that we definitely know the distance to is the start.
	// It has the highest priority (0 - all others are negative) so we pop it
	// off the queue, and find the distance to all adjacent nodes.
	// We keep working out from the next nearest node until all nodes have a defined cost.
	// Since we cannot climb up a height difference of more than one, we treat
	// the incremental cost of such a climb as infinity.
	for q.Len() > 0 {
		n := heap.Pop(q).(*pq.Node)
		currPos, currTotal := n.Value, n.Priority*-1

		for _, next := range g.neighbours(currPos) {
			currHeight := g.terrain[currPos]
			nextHeight := g.terrain[next]

			if nextHeight > currHeight+1 {
				// too steep - try another way.
				continue
			}

			// the distance from start to next, if we arrive via curr:
			alt := currTotal + 1
			if alt < dist[next] {
				dist[next] = alt
				p := pointers[next.X*g.size+next.Y]
				q.Update(p, p.Value, -1*alt)
			}
		}
	}

	return dist[g.end]
}

// neighbours returns a slice of all positions within the grid that
// are adjacent to the given position.
// Squares are only adjacent vertically and horizontally - not diagonally.
func (g grid) neighbours(pos twod.Point) []twod.Point {
	points := make([]twod.Point, 0, 4)
	for _, p := range pos.Neighbours4() {
		if _, ok := g.terrain[p]; ok {
			points = append(points, p)
		}
	}
	return points
}
