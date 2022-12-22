package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
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
	forest, path, err := parseInput(file)
	if err != nil {
		log.Fatalf("parse: %d", err)
	}

	p1 := part1(forest, path)

	middle := time.Now()

	// file.Seek(0, io.SeekStart)
	// tree, err = parsetree(file)
	// if err != nil {
	// 	log.Fatal("parse: ", err)
	// }

	// p2, err := part2(tree)
	// if err != nil {
	// 	log.Fatalf("part2: %d", err)
	// }

	// end := time.Now()

	fmt.Printf("part 1: %d in %s\n", p1, middle.Sub(start))
	// fmt.Printf("part 2: %d in %s\n", p2, end.Sub(middle))
}

type Forest struct {
	grid   map[v.Point]byte
	origin v.Point // the point (x, y) where y = 0, and x is the left-most position
	max    v.Point // the largest individual values for x and y. Note (x, y) may not be present.
}

type Step struct {
	qty int
	dir byte
}

func part1(f Forest, path []Step) int {
	return 0
}

func part2(f Forest, path []Step) int {
	return 0
}

func parseInput(r io.Reader) (Forest, []Step, error) {
	s := bufio.NewScanner(r)

	f := Forest{
		grid: make(map[v.Point]byte, 50*50*6),
	}

	var foundOrigin bool

	var x, y int
	for s.Scan() {
		b := s.Bytes()
		if len(b) == 0 {
			break
		}

		// use separate x and i values so we can align test input with tabs,
		// and skip the tab character when parsing.
		x = 0 // coordinate on the grid
		for i := 0; i < len(b); i, x = i+1, x+1 {
			switch b[i] {
			case '.', '#':
				f.grid[v.Point{X: x, Y: y}] = b[i]
				if !foundOrigin {
					foundOrigin = true
					f.origin.X = x
				}
				if x > f.max.X {
					f.max.X = x
				}

			case ' ':
				continue
			case '\t':
				x--
				continue
			default:
				return Forest{}, nil, fmt.Errorf("line %d: character %d: got '%c'", y, i, b[i])
			}
		}
		y++
	}
	f.max.Y = y - 1

	s.Scan()
	if err := s.Err(); err != nil {
		return Forest{}, nil, err
	}

	p, err := parsePath(s.Text())
	if err != nil {
		return Forest{}, nil, err
	}

	return f, p, nil
}

var _re = regexp.MustCompile(`([0-9]+|[LR])`)

func parsePath(s string) ([]Step, error) {
	path := make([]Step, 0, 100)

	for _, s := range _re.FindAllString(s, -1) {
		switch s {
		case "L", "R":
			path = append(path, Step{dir: s[0]})
		default:
			n, _ := strconv.Atoi(s)
			path = append(path, Step{qty: n, dir: 'F'})
		}
	}

	return path, nil
}
