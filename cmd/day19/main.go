package main

import (
	"bufio"
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

	p2 := part2(blueprints)
	end := time.Now()

	fmt.Printf("part 1: %d in %s\n", p1, middle.Sub(start))
	fmt.Printf("part 2: %d in %s\n", p2, end.Sub(middle))
}

func part1(blueprints []Blueprint) int {
	const limit = 24

	sum := 0
	for i, bp := range blueprints {
		score := evaluate(bp, limit)
		fmt.Printf("Blueprint %d: %d\n", i+1, score)
		sum += (i + 1) * score
	}

	return sum
}

func part2(blueprints []Blueprint) int {
	const limit = 32

	prod := 1
	for i, bp := range blueprints[:3] {
		score := evaluate(bp, limit)
		fmt.Printf("Blueprint %d: %d\n", i+1, score)
		prod *= score
	}

	return prod
}

func evaluate(bp Blueprint, limit int) int {
	// each round of evaluation uses a new queue.  This is the first.
	qFirst := new(pq.Queue[Factory])
	q := &qFirst
	start := &pq.Node[Factory]{
		Value:    Factory{bp: bp, bots: [4]int{1, 0, 0, 0}},
		Priority: 0,
	}

	best := make(map[Factory]int)
	heap.Push(*q, start)
	for t := 1; t <= limit; t++ {
		fmt.Println(t, (*q).Len())
		qNext := new(pq.Queue[Factory])
		pNext := make(map[Factory]*pq.Node[Factory])

		for (*q).Len() > 0 {
			curr := heap.Pop(*q).(*pq.Node[Factory])

			// any time we can build a geodebot, do so:
			if curr.Value.CanAfford(Geodebot) {
				// make a copy of the current state
				f := curr.Value
				score := f.Tick(Geodebot)
				upsert(qNext, &best, &pNext, f, score)
				// building a Geodebot will always be better than any other
				// options - don't bother evaluating the remaining states
				continue
			}

			// next priority: if we can build an obsidian bot, do do:
			if curr.Value.CanAfford(Obsbot) {
				f := curr.Value
				score := (&f).Tick(Obsbot)
				upsert(qNext, &best, &pNext, f, score)
				// building an obsidian bot will always
				// be better than building an ore bot or a clay bot,
				// but *might* not be as good as building nothing. This would
				// only be true if it means we will end the simulation soon,
				// and would be able to afford a geodebot sooner if we wait
			} else {
				if curr.Value.CanAfford(Claybot) {
					f := curr.Value
					score := (&f).Tick(Claybot)
					upsert(qNext, &best, &pNext, f, score)
				}

				if curr.Value.CanAfford(Orebot) {
					f := curr.Value
					score := (&f).Tick(Orebot)
					upsert(qNext, &best, &pNext, f, score)
				}
			}

			// finally, consider the option of not building anything:
			f := curr.Value
			score := (&f).Tick()
			upsert(qNext, &best, &pNext, f, score)
		}
		*q = qNext
	}

	max := 0
	for k, v := range best {
		if v > max {
			fmt.Println("score", v, "state", k)
			max = v
		}
	}
	return max
}

// upsert checks to see if the given key, score combination beats what currently exists
// within the map 'best'. If so, it will either add or update the value in the queue.
// Performs some optimisations to avoid duplicate state in the queue.
func upsert(q *pq.Queue[Factory], best *map[Factory]int, pointers *map[Factory]*pq.Node[Factory], key Factory, score int) {
	if sc, ok := (*best)[key]; ok && score <= sc {
		return
	}
	(*best)[key] = score

	p, ok := (*pointers)[key]
	if !ok {
		p = &pq.Node[Factory]{
			Value:    key,
			Priority: score,
		}
		heap.Push(q, p)
		(*pointers)[key] = p
	} else {
		q.Update(p, key, score)
	}
}

func readInput(r io.Reader) ([]Blueprint, error) {
	s := bufio.NewScanner(r)

	blueprints := make([]Blueprint, 0, 30)
	for s.Scan() {
		bp, err := ParseBlueprint(s.Text())
		if err != nil {
			return nil, err
		}
		blueprints = append(blueprints, bp)
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return blueprints, nil
}
