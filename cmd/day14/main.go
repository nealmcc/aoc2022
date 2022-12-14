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

	"github.com/nealmcc/aoc2022/pkg/vector/twod"
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

	p1 := part1(&cave)
	middle := time.Now()

	p2 := part2(&cave)
	end := time.Now()

	fmt.Printf("part 1: %d in %s\n", p1, middle.Sub(start))
	fmt.Printf("part 2: %d in %s\n", p2, end.Sub(middle))
}

// read the lines from the given input.
func read(r io.Reader) (Cavern, error) {
	s := bufio.NewScanner(r)
	cave := Cavern{}

	// re-use a buffer of points to reduce memory allocation
	rocks := make([]twod.Point, 0, 200)
	for s.Scan() {
		rocks = rocks[:0]
		if err := parseRow(s.Bytes(), &rocks); err != nil {
			return Cavern{}, err
		}
		for _, sq := range rocks {
			cave.Set(sq, Rock)
		}
	}

	if err := s.Err(); err != nil {
		return Cavern{}, err
	}

	return cave, nil
}

func parseRow(b []byte, buf *[]twod.Point) error {
	corners := bytes.Split(b, []byte(" -> "))
	if len(corners) < 2 {
		return errors.New("row must have at least two corners")
	}

	var curr twod.Point
	if err := curr.Parse(corners[0]); err != nil {
		return err
	}
	*buf = append(*buf, curr)

	var next twod.Point
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
func part1(cave *Cavern) int {
	src := twod.Point{X: 500, Y: 0}

	for {
		_, ok := cave.DropSand(src)
		if !ok {
			break
		}
	}
	return cave.Count(Sand)
}

// part2 solves part 2 of the puzzle:
//
// Assume the cave has a floor two squares lower than what the scan results show.
// Drop sand until it reaches the peak (at 500,0).
// When that happens, how many units of sand are in the cave?
func part2(cave *Cavern) int {
	src := twod.Point{X: 500, Y: 0}
	y := cave.lowest + 2
	for x := src.X - y - 1; x < src.X+y+2; x++ {
		cave.Set(twod.Point{X: x, Y: y}, Rock)
	}

	for {
		_, ok := cave.DropSand(src)
		if !ok {
			break
		}
	}
	return cave.Count(Sand)
}
