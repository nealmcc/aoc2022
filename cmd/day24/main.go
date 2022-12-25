package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	pq "github.com/nealmcc/aoc2022/pkg/collection/prioqueue"
	v "github.com/nealmcc/aoc2022/pkg/vector/twod"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { file.Close() }()

	start := time.Now()
	storm, err := parse(file)
	if err != nil {
		log.Fatal(err)
	}

	p1, err := part1(storm)
	if err != nil {
		log.Fatal(err)
	}
	middle := time.Now()

	p2, err := part2(storm, p1, false)
	if err != nil {
		log.Fatal(err)
	}
	end := time.Now()

	fmt.Printf("part 1: %d in %s\n", p1, middle.Sub(start))
	fmt.Printf("part 2: %d in %s\n", p2, end.Sub(middle))
}

// State is a current state of the grid. It combines the current position with
// the positions of all the ice in the storm.
//
// We represent all ice combinations by using some value t that where that value
// of t will produce the ice combination, when passed to storm.IceAt()
// This acts as a unique hash of the ice locations.
type State struct {
	v.Point     // the position on the grid
	iceHash int // the set of current positions, identified by t mod lcm
}

// part1 finds the fewest number of turns to travel through the storm.
// Uses the A* search algorithm.
func part1(st Storm) (int, error) {
	return solve(st, 0)
}

func part2(st Storm, p1 int, verbose bool) (int, error) {
	st.start, st.end = st.end, st.start
	cost, err := solve(st, p1, verbose)
	if err != nil {
		return 0, err
	}

	sum := p1 + cost
	st.start, st.end = st.end, st.start
	cost, err = solve(st, sum, verbose)
	if err != nil {
		return 0, err
	}
	sum += cost

	return sum, nil
}

func solve(storm Storm, startTime int, verbose ...bool) (int, error) {
	debug := len(verbose) > 0 && verbose[0]

	var (
		cost = map[State]int{{
			Point:   storm.start,
			iceHash: startTime,
		}: 0} // known exact costs for each state
		visited  = make(map[State]bool, 100)       // states we are finished with
		q        = new(pq.Queue[State])            // next states to examine
		pointers = make(map[State]*pq.Node[State]) // pointers to those states
		path     = make(map[State]State)           // the previous node
	)

	heap.Push(q, &pq.Node[State]{
		Value: State{Point: storm.start},
	})
	for q.Len() > 0 {
		node := heap.Pop(q).(*pq.Node[State])
		keyCurr := node.Value
		delete(pointers, keyCurr)

		visited[keyCurr] = true
		costCurr := cost[keyCurr]

		if debug {
			fmt.Printf("\n== priority %d ==\n\tarrived at %v from %v at time t=%d (+%d)\n",
				node.Priority, keyCurr, path[keyCurr], costCurr, startTime)

			buf := storm.At(startTime + costCurr).Render()
			compose(buf, map[v.Point]byte{
				keyCurr.Add(v.Point{X: 1, Y: 1}): 'E',
			})
			fmt.Println(string(bytes.Join(buf, []byte("\n"))))
		}

		if keyCurr.Point == storm.end {
			return cost[keyCurr], nil
		}

		for _, move := range [...]Ice{South, East, None, North, West} {
			posNext, costNext := keyCurr.Add(move.AsVector()), costCurr+1

			if keyCurr.Point != storm.start && posNext == storm.start {
				// don't bother re-entering the start point after we've left it.
				continue
			}

			if debug {
				fmt.Printf("considering moving %s to %s at t = %d", move, posNext, startTime+costNext)
			}

			ice, ok := storm.IceAt(posNext, startTime+costNext)
			if !ok {
				if debug {
					fmt.Println("; out of bounds.")
				}
				continue
			}

			if ice > None {
				if debug {
					fmt.Println("; there will be ice.")
				}
				continue
			}

			keyNext := State{
				Point:   posNext,
				iceHash: startTime + costNext,
			}

			if debug {
				fmt.Printf("; icehash %d ", keyNext.iceHash)
			}

			if visited[keyNext] {
				if debug {
					fmt.Printf("; we already visited %s with icehash %d\n",
						keyNext, keyNext.iceHash)
				}
				continue
			}

			bestSoFar, ok := cost[keyNext]
			if ok {
				if debug {
					fmt.Printf("; costNext: %d vs best found so far: %d", costNext, bestSoFar)
				}
				if costNext >= bestSoFar {
					if debug {
						fmt.Println("; this not better - skipping.")
					}
					continue
				} else {
					if debug {
						fmt.Print("; this is better; ")
					}
				}
			} else {
				if debug {
					fmt.Print("; this is a new state for us")
				}
			}

			cost[keyNext] = costNext
			path[keyNext] = keyCurr

			prio := -1 * (costNext + v.ManhattanLength(storm.end.Sub(posNext)))
			if p, ok := pointers[keyNext]; ok {
				if debug {
					fmt.Println("; updating priority to ", prio)
				}
				q.Update(p, keyNext, prio)
			} else {
				p = &pq.Node[State]{
					Value:    keyNext,
					Priority: prio,
				}
				if debug {
					fmt.Println("; adding the state to the queue with priority ", prio)
				}
				heap.Push(q, p)
			}
		}
	}

	return 0, errors.New("no path found")
}

func parse(r io.Reader) (Storm, error) {
	s := bufio.NewScanner(r)

	storm := Storm{
		grid: make(map[v.Point]Ice, 25*121),
	}

	y := -1

	for s.Scan() {
		row := s.Bytes()
		for i, x := 0, -1; i < len(row); i, x = i+1, x+1 {
			if row[i] == '#' {
				if y == -1 && x > storm.extents.Max.X {
					storm.extents.Max.X = x
				}
				continue
			}

			val, err := ParseIce(row[i])
			if err != nil {
				return Storm{}, fmt.Errorf("line %d: character %d: %w", y+1, i, err)
			}
			p := v.Point{X: x, Y: y}
			if y == -1 {
				storm.start = p
			}
			storm.grid[p] = val
			storm.end = p
		}
		y++
	}

	// the walls and entry / exit points are not included in the storm extents:
	storm.extents.Max.X -= 1
	storm.extents.Max.Y = y - 2

	if err := s.Err(); err != nil {
		return Storm{}, err
	}

	return storm, nil
}
