package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"time"
)

func main() {
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

	file.Seek(0, io.SeekStart)
	p2, err := part2(lines)
	if err != nil {
		log.Fatal(err)
	}

	end := time.Now()

	fmt.Printf("part 1: %d in %s\n", p1, middle.Sub(start))
	fmt.Printf("part 2: %d in %s\n", p2, end.Sub(middle))
}

// data is either a float64 or a slice of float64.
type data []any

// read the lines from the given input.
func read(r io.Reader) ([]data, error) {
	s := bufio.NewScanner(r)
	lines := make([]data, 0, 300)

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
func part1(lines []data) int {
	pairs := makePairs(lines)
	sum := 0
	for i, pair := range pairs {
		diff := data(pair).Compare(0, 1)
		if diff <= 0 {
			sum += i + 1
		}
	}
	return sum
}

func makePairs(lines []data) []data {
	pairs := make([]data, len(lines)/2)
	for i := 0; i < len(lines); i += 2 {
		pairs[i/2] = data{any(lines[i]), any(lines[i+1])}
	}
	return pairs
}

// part2 solves part 2 of the puzzle:
//
// If the lines were sorted, determine the indices (using 1-based indices)
// of the two divider packets, and multiply them together.
func part2(lines []data) (int, error) {
	d := make(data, len(lines), len(lines)+2)
	for i := 0; i < len(lines); i++ {
		d[i] = lines[i]
	}
	var div1, div2 data
	if err := json.Unmarshal([]byte("[[2]]"), &div1); err != nil {
		return 0, err
	}
	if err := json.Unmarshal([]byte("[[6]]"), &div2); err != nil {
		return 0, err
	}

	d = append(d, div1, div2)
	sort.Slice(d, func(i, j int) bool {
		diff := d.Compare(i, j)
		return diff < 0
	})

	// sadly, []any is not comparable in Go, so we have to compare strings
	var i2, i6 int
loop:
	for i, data := range d {
		switch text := fmt.Sprintf("%v", data); text {
		case "[[2]]":
			i2 = i + 1

		case "[[6]]":
			i6 = i + 1
			break loop
		}
	}

	return i2 * i6, nil
}

// Compare returns a number indicating the relative sizes of d[i] vs d[j].
// It returns a negative number if d[i] < d[j],
// a positive number if d[j] > d[i],
// or zero if they are equal.
func (d data) Compare(i, j int) int {
	var (
		left, right           any
		leftIsNum, rightIsNum bool
		leftNum, rightNum     float64 // floats because we parsed this as JSON.
		leftSlice, rightSlice data
	)
	left, right = d[i], d[j]

	var ok bool
	if leftNum, leftIsNum = left.(float64); !leftIsNum {
		if leftSlice, ok = toSlice(left); !ok {
			log.Fatalf("left is neither numeric nor a slice: %#v", left)
		}
	}

	if rightNum, rightIsNum = right.(float64); !rightIsNum {
		if rightSlice, ok = toSlice(right); !ok {
			log.Fatalf("right is neither numeric nor a slice: %#v", right)
		}
	}

	if leftIsNum && rightIsNum {
		return int(leftNum) - int(rightNum)
	}

	if leftIsNum {
		leftSlice = data{leftNum}
	}

	if rightIsNum {
		rightSlice = data{rightNum}
	}

	for i := 0; i < len(leftSlice); i++ {
		if i >= len(rightSlice) {
			return 1
		}

		pair := data{leftSlice[i], rightSlice[i]}
		if diff := pair.Compare(0, 1); diff != 0 {
			return diff
		}
	}

	return len(leftSlice) - len(rightSlice)
}

// toSlice asserts that the given value is 'data'.
func toSlice(x any) (data, bool) {
	d, ok := x.(data)
	if ok {
		return d, true
	}

	slice, ok := x.([]any)
	if ok {
		return data(slice), true
	}

	return nil, false
}
