package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/nealmcc/aoc2022/pkg/collection"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { file.Close() }()

	start := time.Now()
	sh, err := parseBlocks(file)
	if err != nil {
		log.Fatalf("parse: %d", err)
	}

	p1 := part1(sh)
	middle := time.Now()

	p2 := part2(sh)
	end := time.Now()

	fmt.Printf("part 1: %d in %s\n", p1, middle.Sub(start))
	fmt.Printf("part 2: %d in %s\n", p2, end.Sub(middle))
}

type point struct {
	x, y, z int
}

func parseBlocks(r io.Reader) (map[point]struct{}, error) {
	s := bufio.NewScanner(r)

	// all of the 1x1x1 cubes in the shape
	blocks := map[point]struct{}{}

	for s.Scan() {
		p, _ := parsePoint(s.Text())
		blocks[p] = struct{}{}
	}

	return blocks, nil
}

func parsePoint(s string) (point, error) {
	parts := strings.Split(s, ",")
	if len(parts) != 3 {
		return point{}, fmt.Errorf("parse(%q): got %d parts; want 3", s, len(parts))
	}
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return point{}, err
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return point{}, err
	}
	z, err := strconv.Atoi(parts[2])
	if err != nil {
		return point{}, err
	}
	return point{x, y, z}, nil
}

// part1 finds the total surface area of the given shape.
func part1(blocks map[point]struct{}) int {
	sum := 0
	for p := range blocks {
		for _, dir := range dir6() {
			neighbour := p.add(dir)
			if _, ok := blocks[neighbour]; !ok {
				sum++
			}
		}
	}
	return sum
}

// part2 finds the exterior surface area of the given shape.
func part2(blocks map[point]struct{}) int {
	bounds := setFloodBoundary(blocks)
	stk := collection.Stack[point]{}

	// 0,0,0 is outside the shape but within our flood fill boundary
	stk.Push(point{})

	// each point that we find within the flood boundary that is not
	// part of the shape will be tracked here.
	exterior := make(map[point]struct{}, 64)

	sum := 0
	for stk.Len() > 0 {
		curr, _ := stk.Pop()
		if _, ok := exterior[curr]; ok {
			continue
		}
		if !inbounds(bounds, curr) {
			continue
		}
		if _, ok := blocks[curr]; ok {
			sum++
			continue
		}
		exterior[curr] = struct{}{}
		for _, dir := range dir6() {
			stk.Push(curr.add(dir))
		}
	}

	return sum
}

func identity() []point {
	return []point{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
}

func dir6() []point {
	return []point{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
		{-1, 0, 0},
		{0, -1, 0},
		{0, 0, -1},
	}
}

func (a point) add(b point) point {
	return point{
		x: a.x + b.x,
		y: a.y + b.y,
		z: a.z + b.z,
	}
}

func (a point) times(n int) point {
	return point{a.x * n, a.y * n, a.z * n}
}

// dot returns the dot product of vectors a and b
func (a point) dot(b point) int {
	return a.x*b.x + a.y*b.y + a.z*b.z
}

// inbounds checks to see if the given point is within
// the given boundary map.
func inbounds(bounds map[point]int, p point) bool {
	for dir, max := range bounds {
		if p.dot(dir) > max {
			return false
		}
	}
	return true
}

// setFloodBoundary defines the upper and lower bounds to use when performing
// a flood fill around the given shape.
func setFloodBoundary(blocks map[point]struct{}) map[point]int {
	bounds := make(map[point]int, 6)

	// set initial upper boundary to -1
	for _, dir := range identity() {
		bounds[dir] = -1
	}

	// adjust upper boundary
	for p := range blocks {
		for _, dir := range identity() {
			n := p.dot(dir) + 1
			if n > bounds[dir] {
				bounds[dir] = n
			}
		}
	}

	// set the lower boundary to 1 for each negative direction:
	for _, dir := range identity() {
		neg1 := dir.times(-1)
		bounds[neg1] = 1
	}

	return bounds
}
