package main

import (
	"container/heap"

	coll "github.com/nealmcc/aoc2022/pkg/collection"
)

// Graph is the full set of distances between any two points in the graph,
// and also can be used to find the path between them.
type Graph map[ValveID]map[ValveID]Marker

// Marker is the distance needed to reach a valve, along with the ID
// of the previous valve in the path.
type Marker struct {
	Dist int
	Prev ValveID
}

// Route is a potential path through the graph
type Route = string

// NewGraph finds the shortest path in to move from each valve to each
// other valve without opening any along the way.
// Uses Dijkstra's algorithm.
func NewGraph(valves map[ValveID]*Valve) Graph {
	result := make(Graph)

	// we store a pointer to each node so that we can update its priority in the queue.
	// we allocate this memory once, outside the loop to reduce garbage collection.
	pointers := make(map[ValveID]*coll.Node[ValveID], len(valves))

	const infinity = 1<<63 - 1
	shortestPath := func(start ValveID) {
		cost := make(map[ValveID]Marker, len(valves))

		// assign initial priorities for each node:
		q := new(coll.PrioQueue[ValveID])
		for id := range valves {
			if id != start {
				cost[id] = Marker{Dist: infinity}
			}
			node := &coll.Node[ValveID]{Value: id, Priority: -1 * cost[id].Dist}
			pointers[id] = node
			heap.Push(q, node)
		}

		// progress outward from the next nearest node:
		for q.Len() > 0 {
			node := heap.Pop(q).(*coll.Node[ValveID])
			curr, cumulativeDist := node.Value, node.Priority*-1
			v := valves[curr]
			for _, next := range v.Neighbours {
				// the distance from start to next, if we arrive via curr:
				alt := cumulativeDist + 1
				if alt < cost[next].Dist {
					cost[next] = Marker{Dist: alt, Prev: curr}
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
