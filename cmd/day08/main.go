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
		log.Fatal(err)
	}
	defer file.Close()

	start := time.Now()

	forest, err := NewForest(file)
	if err != nil {
		log.Fatal(err)
	}

	p1 := part1(forest)
	middle := time.Now()
	p2 := part2(forest)
	end := time.Now()

	fmt.Printf("part 1: %d in %s\n", p1, middle.Sub(start))
	fmt.Printf("part 2: %d in %s\n", p2, end.Sub(middle))
}

// Visibility counts the number of trees in the forest that are visible from
// outside the forest.  A tree is visible from a given direction if all of the
// other trees between it and an edge of the grid are shorter than it.
func part1(f Forest) int {
	return len(f.Visibility())
}

// part2 returns the largest scene score from the forest
func part2(f Forest) int {
	_, sc := f.SceneScore()
	return sc
}

// Forest is a grid of trees. Each tree has a height in the range ['0','9'].
type Forest [][]byte

// NewForest initialises a forest from the given input.
func NewForest(r io.Reader) (Forest, error) {
	g := make(Forest, 0, 100)

	s := bufio.NewScanner(r)
	for s.Scan() {
		b := s.Bytes()
		vals := make([]byte, len(b))
		copy(vals, b)
		g = append(g, vals)
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return g, nil
}

func (f Forest) Visibility() Mask {
	visibleTrees := make(Mask)

	for r := 0; r < len(f); r++ {
		for c := 0; c < len(f[r]); c++ {
			pos := Pos{Row: r, Col: c}
			if dirs := f.visibilityAt(pos); dirs > 0 {
				visibleTrees[pos] = dirs
			}
		}
	}

	return visibleTrees
}

// SceneScore returns the best scene score in the forest.
func (f Forest) SceneScore() (Pos, int) {
	var best Pos
	max := 0

	for r := 0; r < len(f); r++ {
		for c := 0; c < len(f[r]); c++ {
			pos := Pos{Row: r, Col: c}
			if sc := f.sceneScoreAt(pos); sc > max {
				best = pos
				max = sc
			}
		}
	}

	return best, max
}

func (f Forest) visibilityAt(pos Pos) Direction {
	vis := top + bottom + right + left
	height := f[pos.Row][pos.Col]

	// visibility from the top:
	for r, c := pos.Row-1, pos.Col; r >= 0; r-- {
		if f[r][c] >= height {
			vis -= top
			break
		}
	}

	// visibility from the right:
	for r, c := pos.Row, pos.Col+1; c < len(f[r]); c++ {
		if f[r][c] >= height {
			vis -= right
			break
		}
	}

	// visibility from the bottom:
	for r, c := pos.Row+1, pos.Col; r < len(f); r++ {
		if f[r][c] >= height {
			vis -= bottom
			break
		}
	}

	// visibility from the left:
	for r, c := pos.Row, pos.Col-1; c >= 0; c-- {
		if f[r][c] >= height {
			vis -= left
			break
		}
	}

	return vis
}

// sceneScoreAt returns the scene score for the given position.
func (f Forest) sceneScoreAt(pos Pos) int {
	sc := 1
	height := f[pos.Row][pos.Col]

	// number of trees visible to the top:
	count := 0
	for r, c := pos.Row-1, pos.Col; r >= 0; r-- {
		count++
		if f[r][c] >= height {
			break
		}
	}
	sc *= count

	// number of trees visible to the right:
	count = 0
	for r, c := pos.Row, pos.Col+1; c < len(f[r]); c++ {
		count++
		if f[r][c] >= height {
			break
		}
	}
	sc *= count

	// number of trees visible to the bottom:
	count = 0
	for r, c := pos.Row+1, pos.Col; r < len(f); r++ {
		count++
		if f[r][c] >= height {
			break
		}
	}
	sc *= count

	// number of trees visible to the left:
	count = 0
	for r, c := pos.Row, pos.Col-1; c >= 0; c-- {
		count++
		if f[r][c] >= height {
			break
		}
	}
	sc *= count

	return sc
}
