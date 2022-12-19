package main

import (
	"container/heap"
	"fmt"
	"io"
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

	blueprints, err := readInput(file)
	if err != nil {
		log.Fatal(err.Error())
	}
	start := time.Now()

	p1 := part1(blueprints)
	middle := time.Now()

	// p2 := part2(blueprints)
	// end := time.Now()

	fmt.Printf("part 1: %d in %s\n", p1, middle.Sub(start))
	// fmt.Printf("part 2: %d in %s\n", p2, end.Sub(middle))
}

func part1(bp []blueprint) int {
	return 0
}

type state struct {
	pos1       int    // current position
	pos2       int    // current position of elephant (part 2)
	openValves uint64 // which valves are currently open, expressed as a bitmask
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

func readInput(r io.Reader) ([]blueprint, error) {
	return nil, nil
}
