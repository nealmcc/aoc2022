package main

import (
	"container/heap"
	"fmt"
	"log"
	"os"
	"time"

	pq "github.com/nealmcc/aoc2022/pkg/collection/prioqueue"
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

	p2 := part2(graph)
	end := time.Now()

	fmt.Printf("part 1: %d in %s\n", p1, middle.Sub(start))
	fmt.Printf("part 2: %d in %s\n", p2, end.Sub(middle))
}

// state is a comparable struct that allows us to reduce the search space.
// See https://go.dev/ref/spec for more about 'comparable' in Go.
type state struct {
	pos1       int    // current position
	pos2       int    // current position of elephant (part 2)
	openValves uint64 // which valves are currently open, expressed as a bitmask
}

// part1 solves part 1 of the puzzle
//
// We simulate a turn-by-turn breadth-first traversal of the graph, and keep
// track of the best possible accumulated flow for each combination of
// (position, and set of open valves) for that turn. By discarding state that
// would have the same current position and open valves, but less accumulated
// flow, we significantly reduce the search space for the next step.
//
// In order to do so, we need the states for a given turn to be comparable.
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

	// each round of evaluation uses a new queue.  This is the first.
	qFirst := new(pq.Queue[state])
	q := &qFirst
	start := &pq.Node[state]{
		Value:    state{pos1: getIndex(K("AA"))},
		Priority: 0,
	}

	best := make(map[state]int)
	heap.Push(*q, start)
	for t := 1; t < limit; t++ {
		qNext := new(pq.Queue[state])
		pNext := make(map[state]*pq.Node[state])

		for (*q).Len() > 0 {
			curr := heap.Pop(*q).(*pq.Node[state])

			flow := valves[curr.Value.pos1].Flow
			mask := uint64(1) << curr.Value.pos1
			canOpen := flow > 0 && curr.Value.openValves&mask == 0
			neighbours := valves[curr.Value.pos1].Neighbours

			for _, n := range neighbours {
				ix := getIndex(n)
				k := state{pos1: ix, openValves: curr.Value.openValves}
				upsert(qNext, &best, &pNext, k, curr.Priority)
			}

			if canOpen {
				// add a state where we open this valve
				k := state{pos1: curr.Value.pos1, openValves: curr.Value.openValves | mask}
				score := curr.Priority + (limit-t)*flow
				upsert(qNext, &best, &pNext, k, score)
			}

		}
		*q = qNext
	}

	max := 0
	for _, v := range best {
		if v > max {
			max = v
		}
	}
	return max
}

// part 2 works similarly to part 1, except that we track the elephant's
// position as well as our own.
func part2(input map[ValveID]*Valve) int {
	const limit = 26
	_, valves := index(input)

	getIndex := func(k ValveID) int {
		return input[k].ix
	}

	// each round of evaluation uses a new queue.  This is the first.
	qFirst := new(pq.Queue[state])
	q := &qFirst
	start := &pq.Node[state]{
		Value: state{
			pos1: getIndex(K("AA")),
			pos2: getIndex(K("AA")),
		},
		Priority: 0,
	}

	best := make(map[state]int, 13000000)
	heap.Push(*q, start)
	for t := 1; t < limit; t++ {
		qNext := new(pq.Queue[state])
		pNext := make(map[state]*pq.Node[state])

		for (*q).Len() > 0 {
			curr := heap.Pop(*q).(*pq.Node[state])

			flow1 := valves[curr.Value.pos1].Flow
			mask1 := uint64(1) << curr.Value.pos1
			canOpen1 := flow1 > 0 && curr.Value.openValves&mask1 == 0
			neighbours1 := valves[curr.Value.pos1].Neighbours

			flow2 := valves[curr.Value.pos2].Flow
			mask2 := uint64(1) << curr.Value.pos2
			canOpen2 := flow2 > 0 && curr.Value.openValves&mask2 == 0
			neighbours2 := valves[curr.Value.pos2].Neighbours

			if canOpen1 && canOpen2 {
				// add a state where we both open the valve in our current room
				k := state{curr.Value.pos1, curr.Value.pos2, curr.Value.openValves | mask1 | mask2}
				score := curr.Priority + (limit-t)*(flow1+flow2)
				upsert(qNext, &best, &pNext, k, score)
			}

			if canOpen1 {
				// add states where I open the valve, but the elephant moves
				for _, n2 := range neighbours2 {
					ix := getIndex(n2)
					if ix != curr.Value.pos1 {
						k := state{curr.Value.pos1, ix, curr.Value.openValves | mask1}
						score := curr.Priority + (limit-t)*flow1
						upsert(qNext, &best, &pNext, k, score)
					}
				}
			}

			if canOpen2 {
				// add states where the elephant opens the valve, but I move
				for _, n1 := range neighbours1 {
					ix := getIndex(n1)
					if ix != curr.Value.pos2 {
						k := state{ix, curr.Value.pos2, curr.Value.openValves | mask2}
						score := curr.Priority + (limit-t)*flow2
						upsert(qNext, &best, &pNext, k, score)
					}
				}
			}

			// we both move
			for _, n1 := range neighbours1 {
				ix1 := getIndex(n1)
				for _, n2 := range neighbours2 {
					ix2 := getIndex(n2)
					if ix1 == ix2 {
						continue
					}
					k := state{ix1, ix2, curr.Value.openValves}
					upsert(qNext, &best, &pNext, k, curr.Priority)
				}
			}
		}

		*q = qNext
	}

	max := 0
	for _, v := range best {
		if v > max {
			max = v
		}
	}
	return max
}

// upsert checks to see if the given key, score combination beats what currently exists
// within the map 'best'. If so, it will either add or update the value in the queue.
// Performs some optimisations to avoid duplicate state in the queue.
func upsert(q *pq.Queue[state], best *map[state]int, pointers *map[state]*pq.Node[state], key state, score int) {
	if key.pos2 != 0 && key.pos1 > key.pos2 {
		key.pos1, key.pos2 = key.pos2, key.pos1
	}

	if sc, ok := (*best)[key]; ok && score <= sc {
		return
	}
	(*best)[key] = score

	p, ok := (*pointers)[key]
	if !ok {
		p = &pq.Node[state]{
			Value:    key,
			Priority: score,
		}
		heap.Push(q, p)
		(*pointers)[key] = p
	} else {
		q.Update(p, key, score)
	}
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
