package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/nealmcc/aoc2022/pkg/quickperm"
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

	p1 := part1(valves)
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
func part1(valves map[ValveID]*Valve) int {
	network := NewNetwork(valves)

	var maxflow int

	destset := make(ValveSet, len(valves)/3)
	for id, valve := range network.v {
		if valve.Flow != 0 {
			maxflow += valve.Flow
			destset[id] = struct{}{}
		}
	}

	destinations := make([]ValveID, 0, len(destset))
	for id := range destset {
		destinations = append(destinations, id)
	}

	best := state{
		missedFlow: 1<<63 - 1,
	}
	for route := range quickperm.Permutations(destinations) {
		s := routeStringer(route).String()
		state := network.TransitionMany(s)
		if state.mins > 30 {
			fmt.Printf("re-evaluate algorithm due to route %s \n", s)
		}
		if state.missedFlow < best.missedFlow {
			best = state
			fmt.Printf("new best route %s with total flow %d after %d mins\n", s, best.totalFlow, best.mins)
		}
	}

	return (30-best.mins)*maxflow + best.totalFlow
}

type routeStringer []ValveID

func (r routeStringer) String() string {
	b := make([]byte, 2*len(r))
	for i, id := range r {
		b[2*i] = byte((id & 0xFF00) >> 8)
		b[2*i+1] = byte(id & 0x00FF)
	}
	return string(b)
}
