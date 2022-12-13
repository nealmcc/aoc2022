package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	flag.Parse()

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	start := time.Now()

	lines, err := read(file)
	if err != nil {
		log.Fatal(err)
	}
	p1 := part1(lines)

	middle := time.Now()
	// file.Seek(0, io.SeekStart)
	// p2 := part2(lines)
	// end := time.Now()

	fmt.Printf("part 1: %d] in %s\n", p1, middle.Sub(start))
	// fmt.Printf("part 2: %d in %s\n", p2, end.Sub(middle))
}

// data is either a float64 or a slice of float64.
type data []any

// read the lines from the given input.
func read(r io.Reader) ([][]any, error) {
	s := bufio.NewScanner(r)
	lines := make([][]interface{}, 0, 300)

	for s.Scan() {
		b := s.Bytes()
		if len(b) == 0 {
			continue
		}

		var line []any
		if err := json.Unmarshal(b, &line); err != nil {
			return nil, err
		}

		lines = append(lines, line)
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// part1 solves part 1 of the puzzle:
//
// From the given pairs add up the indices of pairs which are already
// in the correct order. Use 1-based indices instead of 0-based ones.
func part1(lines [][]any) int {
	pairs := makePairs(lines)
	sum := 0
	for i, pair := range pairs {
		fmt.Printf("== Pair %3d ==\n", i+1)

		diff := data(pair).Compare(0, 1, "- ")
		ok := diff <= 0
		fmt.Printf("pair %d is okay: %t\n", i+1, ok)
		if ok {
			sum += i + 1
		}
		fmt.Println()
	}
	return sum
}

func makePairs(lines [][]any) []data {
	pairs := make([]data, len(lines)/2)
	for i := 0; i < len(lines); i += 2 {
		pairs[i/2] = data{lines[i], lines[i+1]}
	}
	return pairs
}

// part2 solves part 2 of the puzzle:
func part2(lines [][]any) int {
	return 0
}

// Compare returns a negative number if d[i] < d[j],
// a positive number if d[j] > d[i], or zero if they are equal.
func (d data) Compare(i, j int, pad string) (diff int) {
	var (
		left, right           any
		leftIsNum, rightIsNum bool
		leftNum, rightNum     float64
		leftSlice, rightSlice []any
	)
	left, right = d[i], d[j]
	fmt.Printf("%sCompare %v vs %v\n", pad, left, right)

	defer func() {
		fmt.Printf("%s diff = %d\n", pad, diff)
	}()

	var ok bool
	if leftNum, leftIsNum = left.(float64); !leftIsNum {
		if leftSlice, ok = left.([]any); !ok {
			log.Fatalf("left is neither numeric nor a slice: %#v", left)
		}
	}

	if rightNum, rightIsNum = right.(float64); !rightIsNum {
		if rightSlice, ok = right.([]any); !ok {
			log.Fatalf("right is neither numeric nor a slice: %#v", left)
		}
	}

	if leftIsNum && rightIsNum {
		return int(leftNum) - int(rightNum)
	}

	if leftIsNum {
		leftSlice = data{leftNum}
		fmt.Printf("%sMixed types; convert left to %v and retry comparison\n", pad, leftSlice)
	}

	if rightIsNum {
		rightSlice = data{rightNum}
		fmt.Printf("%sMixed types; convert right to %v and retry comparison\n", pad, rightSlice)
	}

	if !leftIsNum || !rightIsNum {
		fmt.Printf("%sCompare %v vs %v\n", pad, leftSlice, rightSlice)
	}

	for i := 0; i < len(leftSlice); i++ {
		if i >= len(rightSlice) {
			return 1
		}

		pair := data{leftSlice[i], rightSlice[i]}
		if diff := pair.Compare(0, 1, "  "+pad); diff != 0 {
			return diff
		}
	}

	return len(leftSlice) - len(rightSlice)
}
