package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	v "github.com/nealmcc/aoc2022/pkg/vector/twod"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	start := time.Now()

	cave, err := read(file)
	if err != nil {
		log.Fatal(err)
	}

	animate := os.Getenv("ANIMATE")
	var r RenderFunc
	if animate != "" {
		min := v.Point{X: 332, Y: -1}
		max := v.Point{X: 669, Y: 168}
		renderer := NewRenderer(cave, "part1", min, max, 10)
		r = renderer.SaveNext
	}

	p1, err := part1(cave, r)
	if err != nil {
		log.Fatal(err)
	}
	middle := time.Now()

	p2, err := part2(cave, r)
	if err != nil {
		log.Fatal(err)
	}
	end := time.Now()

	fmt.Printf("part 1: %d in %s\n", p1, middle.Sub(start))
	fmt.Printf("part 2: %d in %s\n", p2, end.Sub(middle))
}

// read the lines from the given input.
func read(r io.Reader) (*Cavern, error) {
	s := bufio.NewScanner(r)
	cave := &Cavern{}

	// re-use a buffer of points to reduce memory allocation
	rocks := make([]v.Point, 0, 200)
	for s.Scan() {
		rocks = rocks[:0]
		if err := parseRow(s.Bytes(), &rocks); err != nil {
			return nil, err
		}
		for _, sq := range rocks {
			cave.Set(sq, Rock)
		}
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return cave, nil
}

func parseRow(b []byte, buf *[]v.Point) error {
	corners := bytes.Split(b, []byte(" -> "))
	if len(corners) < 2 {
		return errors.New("row must have at least two corners")
	}

	var curr v.Point
	if err := curr.Parse(corners[0]); err != nil {
		return err
	}
	*buf = append(*buf, curr)

	var next v.Point
	for i := 1; i < len(corners); i++ {
		if err := next.Parse(corners[i]); err != nil {
			return err
		}
		step, n := next.Sub(curr).Reduce()
		for j := 0; j < n; j++ {
			curr = curr.Add(step)
			*buf = append(*buf, curr)
		}
	}
	return nil
}

// part1 solves part 1 of the puzzle:
//
// Assume the cave has no floor. Drop sand until it falls into the abyss.
// When that happens, how many units of sand are in the cave?
func part1(cave *Cavern, render RenderFunc) (int, error) {
	src := v.Point{X: 500, Y: 0}

	if render != nil {
		render()
	}

	for {
		p, ok := cave.DropSand(src)
		if render != nil {
			if err := render(p); err != nil {
				return 0, fmt.Errorf("render: %w", err)
			}
		}
		if !ok {
			break
		}
	}
	return cave.Count(Sand), nil
}

// part2 solves part 2 of the puzzle:
//
// Assume the cave has a floor two squares lower than what the scan results show.
// Drop sand until it reaches the peak (at 500,0).
// When that happens, how many units of sand are in the cave?
func part2(cave *Cavern, render RenderFunc) (int, error) {
	src := v.Point{X: 500, Y: 0}
	y := cave.lowest + 2
	for x := src.X - y - 1; x < src.X+y+2; x++ {
		cave.Set(v.Point{X: x, Y: y}, Rock)
	}

	if render != nil {
		render()
	}

	for {
		p, ok := cave.DropSand(src)
		if render != nil {
			if err := render(p); err != nil {
				return 0, fmt.Errorf("render: %w", err)
			}
		}
		if !ok {
			break
		}
	}
	return cave.Count(Sand), nil
}
