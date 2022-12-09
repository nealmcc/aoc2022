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

	"github.com/nealmcc/aoc2022/pkg/rope"
	v "github.com/nealmcc/aoc2022/pkg/vector/twod"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	start := time.Now()

	p1, err := solve(file, 2)
	if err != nil {
		log.Fatal(err)
	}

	middle := time.Now()
	file.Seek(0, io.SeekStart)
	p2, err := solve(file, 10)
	if err != nil {
		log.Fatal(err)
	}
	end := time.Now()

	fmt.Printf("part 1: %d in %s\n", p1, middle.Sub(start))
	fmt.Printf("part 2: %d in %s\n", p2, end.Sub(middle))
}

// solve solves both part 1 and part 2
func solve(r io.Reader, n int) (int, error) {
	log := &logger{
		tailPositions: make(map[v.Point]struct{}),
	}
	rope := rope.New(n, log)

	s := bufio.NewScanner(r)
	line := 0
	for s.Scan() {
		line++
		dir, dist, err := parse(s.Text())
		if err != nil {
			return 0, fmt.Errorf("invalid input on line %d: %w", line, err)
		}

		for i := 0; i < dist; i++ {
			if err := rope.Move(dir); err != nil {
				return 0, err
			}
		}
	}

	if err := s.Err(); err != nil {
		return 0, err
	}

	return len(log.tailPositions), nil
}

// logger is responsible for recording the movement of the rope.
type logger struct {
	tailPositions map[v.Point]struct{}
}

// compile-time interface check:
var _ rope.Logger = new(logger)

// Log implements rope.Logger
func (l *logger) Log(knots []v.Point) {
	tail := knots[len(knots)-1]
	l.tailPositions[tail] = struct{}{}
}

// parse the given input line into a direction and distance.
func parse(s string) (v.Point, int, error) {
	parts := strings.Split(s, " ")
	if len(parts) != 2 {
		return v.Point{}, 0, fmt.Errorf("wanted 2 parts ; got %d", len(parts))
	}

	var dir v.Point
	switch parts[0] {
	case "U":
		dir = v.Point{Y: 1}
	case "D":
		dir = v.Point{Y: -1}
	case "R":
		dir = v.Point{X: 1}
	case "L":
		dir = v.Point{X: -1}
	default:
		return v.Point{}, 0, fmt.Errorf("invalid direction: %s", parts[0])
	}

	dist, err := strconv.Atoi(parts[1])
	if err != nil {
		return v.Point{}, 0, fmt.Errorf("invalid distance: %s", parts[1])
	}

	return dir, dist, nil
}
