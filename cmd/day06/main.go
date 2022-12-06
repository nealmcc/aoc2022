package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	p1, err := part1(data)
	if err != nil {
		log.Fatalf("part1: %d", err)
	}

	middle := time.Now()

	p2, err := part2(data)
	if err != nil {
		log.Fatalf("part2: %s", err)
	}
	end := time.Now()

	fmt.Printf("part 1: %d in %s\n", p1, middle.Sub(start))
	fmt.Printf("part 2: %d in %s\n", p2, end.Sub(middle))
}

// part1 returns the index of the start-of-packet marker in the given data.
// A start-of-packet marker occurs *after* a sequence of 4 unique bytes.
// Returns io.EOF if the start-of-packet marker is not found.
func part1(data []byte) (int, error) {
	if len(data) < 4 {
		return 0, errors.New("input too short")
	}

	for j := 4; j < len(data); {
		i, _ := findPair(data[j-4 : j])
		if i == -1 {
			return j, nil
		}
		j += i + 1
	}

	return 0, io.EOF
}

// part2 returns the index of the start-of-message marker in the given data.
// A start-of-message marker occurs *after* a sequence of 14 unique bytes.
// Returns io.EOF if the start-of-message marker is not found.
func part2(data []byte) (int, error) {
	if len(data) < 14 {
		return 0, errors.New("input too short")
	}

	for j := 14; j < len(data); {
		i, _ := findPair(data[j-14 : j])
		if i == -1 {
			return j, nil
		}
		j += i + 1
	}

	return 0, io.EOF
}

// findPair looks for a pair of duplicate bytes in the data.  If it finds such
// a pair, then it returns their indices.
// This function will return the largest pair of indices that it can.
// If there is no duplicate, then both return values will be -1.
func findPair(data []byte) (int, int) {
	for j := len(data) - 1; j >= 1; j-- {
		for i := j - 1; i >= 0; i-- {
			if data[i] == data[j] {
				return i, j
			}
		}
	}
	return -1, -1
}
