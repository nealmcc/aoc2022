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
		log.Fatal(err)
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

func part1(f Forest, path []Step) int {
	curr := f.Origin()
	dir := Right

	for _, s := range path {
		if s.Rotation == 0 {
			fmt.Printf("walking %d %s\n", s.Dist, dir)
			curr, dir = f.Next(curr, s.Dist, dir, f.wrap1)
			continue
		}

		var next Facing
		if s.Rotation == 'L' {
			next = ((dir - 1) + 4) % 4
		} else {
			next = (dir + 1) % 4
		}
		dir = next
		fmt.Printf("turned %c ; standing at %s facing %s\n",
			s.Rotation, curr, dir)
	}
	fmt.Printf("standing at %s facing %s\n", curr, dir)
	return 1000*(curr.Y+1) + 4*(curr.X+1) + int(dir)
}

func part2(f Forest, path []Step) int {
	return 0
}

func parseInput(r io.Reader) (Forest, []Step, error) {
	s := bufio.NewScanner(r)

	f := Forest{
		grid: make(map[v.Point]byte, 50*50*6),
	}

	var x, y int
	for s.Scan() {
		b := s.Bytes()
		if len(b) == 0 {
			break
		}

		// use separate x and i values so we can align test input with tabs,
		// and skip the tab character when parsing.
		x = 0
		for i := 0; i < len(b); i, x = i+1, x+1 {
			switch b[i] {
			case '.', '#':
				f.grid[v.Point{X: x, Y: y}] = b[i]
				if x >= f.width {
					f.width = x + 1
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
	f.height = y

	f.setBounds()

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
			path = append(path, Step{Rotation: s[0]})
		default:
			n, _ := strconv.Atoi(s)
			path = append(path, Step{Dist: n})
		}
	}

	return path, nil
}
