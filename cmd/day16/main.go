package main

import (
	"fmt"
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

	graph, err := ReadValves(file)
	if err != nil {
		log.Fatal(err.Error())
	}
	start := time.Now()

	p1 := part1(graph)
	middle := time.Now()

	// p2 := part2(valves)
	// end := time.Now()

	fmt.Printf("part 1: %d in %s\n", p1, middle.Sub(start))
	// fmt.Printf("part 2: %d in %s\n", p2, end.Sub(middle))
}

// state1 is a comparable struct that allows us to reduce the search space
// for each turn during part1.
// See https://go.dev/ref/spec for more about 'comparable'.
type state1 struct {
	key1
	score int
}

type key1 struct {
	ix   int    // current position
	mask uint64 // which valves are currently open, expressed as a bitmask
}

// part1 solves part 1 of the puzzle
//
// We simulate a turn-by-turn breadth-first traversal of the graph, and keep
// track of the best possible accumulated flow for each combination of
// (position, and set of open valves) for that turn. By discarding state that
// would have the same current position and open valves, but less accumulated
// flow, we significantly reduce the search space for the next step.
//
// In order to do so, we need the stateKeys for a given turn to be comparable.
// That is, they must support the == and != operations.
//
// Due to the above requirement, we can't use a map or a slice to store the set
// of currently open valves. Thankfully, we have less than 64 valves, so we can
// use a uint64 bitmask, where each bit corresponds to a different valve, based
// on its index within a slice. That's why we've had to add the 'ix' property to
// the Valve, and store the Key in a slice, accessed by its index.
func part1(input map[ValveID]*Valve) int {
	const limit = 30
	_, valves := index(input)

	getIndex := func(k ValveID) int {
		return input[k].ix
	}

	t := 0 // turn number (aka minutes elapsed)

	buffers := [][]state1{
		make([]state1, 0, 2048),
		make([]state1, 0, 2048),
	}

	// For each turn, all states and best score for each
	states := append(buffers[0], state1{key1: key1{ix: getIndex(K("AA"))}})
	best := make(map[key1]int, 100000)
	for t = 1; t < limit; t++ {
		newstates := buffers[t%2][:0]
		for _, next := range states {
			k := key1{ix: next.ix, mask: next.mask}
			hiscore, ok := best[k]
			if ok && hiscore >= next.score {
				continue
			}
			best[k] = next.score

			flow := valves[next.ix].Flow
			mask := uint64(1) << next.ix
			if flow > 0 && next.mask&mask == 0 {
				newstates = append(newstates, state1{
					key1:  key1{ix: next.ix, mask: next.mask | mask},
					score: next.score + (limit-t)*flow,
				})
			}

			for _, neighbour := range valves[next.ix].Neighbours {
				newstates = append(newstates, state1{
					key1:  key1{ix: input[neighbour].ix, mask: next.mask},
					score: next.score,
				})
			}

			states = newstates
		}
	}

	max := 0
	for _, s := range states {
		if s.score > max {
			max = s.score
		}
	}
	return max
}

// index organises the input data so it's more efficient to work with.
func index(input map[ValveID]*Valve) (names []ValveID, valves []*Valve) {
	names = make([]ValveID, len(input))
	valves = make([]*Valve, len(input))
	i := 0
	for k, v := range input {
		names[i] = k
		valves[i] = v
		v.ix = i
		i++
	}
	return names, valves
}
