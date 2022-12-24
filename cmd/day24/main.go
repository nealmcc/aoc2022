package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/nealmcc/aoc2022/pkg/vector/twod"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { file.Close() }()

	start := time.Now()
	storm, err := parse(file)
	if err != nil {
		log.Fatal(err)
	}

	p1 := part1(storm)
	middle := time.Now()

	// p2 := part2(storm)
	// end := time.Now()

	fmt.Printf("part 1: %d in %s\n", p1, middle.Sub(start))
	// fmt.Printf("part 2: %d in %s\n", p2, end.Sub(middle))
}

func part1(st Storm) int {
	return 0
}

func part2(st Storm) int {
	return 0
}

func parse(r io.Reader) (Storm, error) {
	s := bufio.NewScanner(r)

	storm := Storm{
		grid: make(map[twod.Point]Ice, 25*121),
	}

	y := -1

	for s.Scan() {
		row := s.Bytes()
		for i, x := 0, -1; i < len(row); i, x = i+1, x+1 {
			if row[i] == '#' {
				if y == -1 && x > storm.extents.Max.X {
					storm.extents.Max.X = x
				}
				continue
			}

			val, err := ParseIce(row[i])
			if err != nil {
				return Storm{}, fmt.Errorf("line %d: character %d: %w", y+1, i, err)
			}
			p := twod.Point{X: x, Y: y}
			if y == -1 {
				storm.entry = p
			}
			storm.grid[p] = val
			storm.exit = p
		}
		y++
	}

	// the walls and entry / exit points are not included in the storm extents:
	storm.extents.Max.X -= 1
	storm.extents.Max.Y = y - 2

	if err := s.Err(); err != nil {
		return Storm{}, err
	}

	return storm, nil
}
