package main

import (
	"bufio"
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
	defer func() { file.Close() }()

	start := time.Now()
	forest, err := parseInput(file)
	if err != nil {
		log.Fatal(err)
	}

	p1 := part1(forest)
	middle := time.Now()

	// p2 := part2(forest, path)
	// end := time.Now()

	fmt.Printf("part 1: %d in %s\n", p1, middle.Sub(start))
	// fmt.Printf("part 2: %d in %s\n", p2, end.Sub(middle))
}

func part1(f Forest) int {
	fmt.Printf("%12v\n", f)
	return -1
}

func parseInput(r io.Reader) (Forest, error) {
	s := bufio.NewScanner(r)

	f := Forest{
		Grid: make(map[v.Point]struct{}, 70*70),
	}

	var x, y int
	for s.Scan() {
		b := s.Bytes()

		x = 0
		for i := 0; i < len(b); i, x = i+1, x+1 {
			switch b[i] {
			case '#':
				f.Grid[v.Point{X: x, Y: y}] = struct{}{}

			case '.', ' ':
				continue

			case '\t':
				x--
				continue

			default:
				return Forest{}, fmt.Errorf("line %d: character %d: got '%c'", y, i, b[i])
			}
		}
		y++
	}

	f.setBounds()

	if err := s.Err(); err != nil {
		return Forest{}, err
	}

	return f, nil
}
