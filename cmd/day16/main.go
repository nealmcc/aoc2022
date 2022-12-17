package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	valves, err := read(file)
	if err != nil {
		log.Fatal(err.Error())
	}
	start := time.Now()

	p1 := part1(valves, 30)
	middle := time.Now()

	// p2 := part2(valves)
	// end := time.Now()

	fmt.Printf("part 1: %d in %s\n", p1, middle.Sub(start))
	// fmt.Printf("part 2: %d in %s\n", p2, end.Sub(middle))
}

// read the lines from the given input.
func read(r io.Reader) (map[ValveID]*Valve, error) {
	s := bufio.NewScanner(r)
	valves := make(map[ValveID]*Valve, 58)

	for s.Scan() {
		v, err := ParseValve(s.Bytes())
		if err != nil {
			return nil, err
		}
		valves[v.ID] = &v
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return valves, nil
}

// part1 solves part 1 of the puzzle
func part1(valves map[ValveID]*Valve, limit int) int {
	// curr := ID("AA")

	return 0
}
