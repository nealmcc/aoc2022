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

	p1 := forest.Visibility()

	middle := time.Now()

	// p2, err := part2(data)
	// if err != nil {
	// 	log.Fatalf("part2: %s", err)
	// }
	// end := time.Now()

	fmt.Printf("part 1: %d in %s\n", len(p1), middle.Sub(start))
	// fmt.Printf("part 2: %d in %s\n", p2, end.Sub(middle))
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
		for i, ch := range b {
			vals[i] = ch - '0'
		}
		g = append(g, vals)
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return g, nil
}

func (f Forest) Size() (int, int) {
	height := len(f)
	if height == 0 {
		return 0, 0
	}
	width := len(f[0])
	return width, height
}

// Visibility returns a 2d bitmask showing the directions that each tree is
// visible from. A tree is visible from a given direction if all of the other trees
// between it and an edge of the grid are shorter than it.
//
// In the resulting Mask, each coordinate will contain the bitwise AND of
// each of the four Directions.
func (f Forest) Visibility() Mask {
	vis := make(Mask)

	for r := 0; r < len(f); r++ {
		for c := 0; c < len(f[r]); c++ {
			pos := Pos{Row: r, Col: c}
			if dirs := f.VisibilityAt(pos); dirs > 0 {
				vis[pos] = dirs
			}
		}
	}

	return vis
}

func (f Forest) VisibilityAt(pos Pos) Direction {
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
