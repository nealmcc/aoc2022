package main

import (
	"bufio"
	"bytes"
	"container/heap"
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
func read(r io.Reader) (map[ValveID]*Valve, error) {
	s := bufio.NewScanner(r)
	valves := make(map[ValveID]*Valve, 58)

	for s.Scan() {
		v, err := parseRow(s.Bytes())
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
func part1(valves map[ValveID]*Valve, limit int) int {
	sum := 0

	return sum
}

// distances finds the shortest path in minutes to move from each valve to each
// other valve without opening any along the way.
// Uses Dijkstra's algorithm.
func distances(valves map[ValveID]*Valve) map[ValveID]map[ValveID]int {
	result := make(map[ValveID]map[ValveID]int)

	// we store a pointer to each node so that we can update its priority in the queue.
	// we allocate this memory once, outside the loop to reduce garbage collection.
	pointers := make(map[ValveID]*Node, len(valves))

	const infinity = 1<<63 - 1
	shortestPath := func(start ValveID) {
		cost := make(map[ValveID]int, len(valves))
		cost[start] = 0

		// assign initial priorities for each node:
		q := new(Queue)
		for id := range valves {
			if id != start {
				cost[id] = infinity
			}
			node := &Node{Value: id, Priority: -1 * cost[id]}
			pointers[id] = node
			heap.Push(q, node)
		}

		// progress outward from the next nearest node:
		for q.Len() > 0 {
			node := heap.Pop(q).(*Node)
			curr, cumulativeDist := node.Value, node.Priority*-1
			v := valves[curr]
			for _, next := range v.Neighbours {
				// the distance from start to next, if we arrive via curr:
				alt := cumulativeDist + 1
				if alt < cost[next] {
					cost[next] = alt
					p := pointers[next]
					q.Update(p, p.Value, -1*alt)
				}
			}
		}

		result[start] = cost
	}

	for id := range valves {
		shortestPath(id)
	}
	return result
}
