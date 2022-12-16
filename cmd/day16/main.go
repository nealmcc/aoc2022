package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
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
func read(r io.Reader) ([]Valve, error) {
	s := bufio.NewScanner(r)
	valves := make([]Valve, 0, 58)

	for s.Scan() {
		v, err := parseRow(s.Bytes())
		if err != nil {
			return nil, err
		}
		valves = append(valves, v)
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return valves, nil
}

var _re = regexp.MustCompile(`Valve ([A-Z][A-Z]) has flow rate=([\d]+); tunnels? leads? to valves? ((?:[A-Z][A-Z])(?:(?:, )?(?:[A-Z][A-Z]))*)`)

func parseRow(b []byte) (Valve, error) {
	m := _re.FindAllSubmatch(b, -1)
	if len(m) != 1 {
		return Valve{}, fmt.Errorf("parse(%q): want 1 match got %d",
			b, len(m))
	}

	flow, err := strconv.Atoi(string(m[0][2]))
	if err != nil {
		return Valve{}, fmt.Errorf("parse flow %q: %w", m[0][2], err)
	}

	v := Valve{
		ID:         ID(string(m[0][1])),
		Flow:       flow,
		Neighbours: make([]ValveID, 0, 5),
	}

	parts := bytes.Split(m[0][3], []byte(", "))
	for _, x := range parts {
		v.Neighbours = append(v.Neighbours, ID(string(x)))
	}

	return v, nil
}

// part1 solves part 1 of the puzzle
func part1(valves []Valve, limit int) int {
	sum := 0

	return sum
}
