package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { file.Close() }()

	nums, err := parseInts(file)
	if err != nil {
		log.Fatalf("parse: %d", err)
	}

	start := time.Now()
	in := make([]int, len(nums))
	copy(in, nums)

	p1 := part1(in)
	middle := time.Now()

	p2 := part2(in)
	end := time.Now()

	fmt.Printf("part 1: %d in %s\n", p1, middle.Sub(start))
	fmt.Printf("part 2: %d in %s\n", p2, end.Sub(middle))
}

// part1 solves part 1
func part1(nums []int) int {
	list := NewList(nums, 0)
	list.Mix()
	return list.Coordinates()
}

// part2 solves part 2
func part2(nums []int) int {
	for i := range nums {
		nums[i] *= 811589153
	}
	list := NewList(nums, 0)
	for i := 0; i < 10; i++ {
		list.Mix()
	}
	return list.Coordinates()
}

func parseInts(r io.Reader) ([]int, error) {
	s := bufio.NewScanner(r)

	nums := make([]int, 0, 5000)
	for s.Scan() {
		row := s.Text()
		n, err := strconv.Atoi(row)
		if err != nil {
			return nil, fmt.Errorf("parse %q: %w", row, err)
		}
		nums = append(nums, n)
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return nums, nil
}
