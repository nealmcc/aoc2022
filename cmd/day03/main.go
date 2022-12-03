package main

import (
	"bufio"
	"bytes"
	"errors"
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
	p1, err := part1(file)
	if err != nil {
		log.Fatalf("part1: %d", err)
	}

	middle := time.Now()
	file.Seek(0, 0)

	p2, err := part2(file)
	if err != nil {
		log.Fatalf("part2: %s", err)
	}
	end := time.Now()

	fmt.Printf("part 1: %d in %s\n", p1, middle.Sub(start))
	fmt.Printf("part 2: %d in %s\n", p2, end.Sub(middle))
}

// part1 determines:
// for each bag, which item is in both of that bag's compartments?
func part1(r io.Reader) (int, error) {
	var sum int
	s := bufio.NewScanner(r)

	for s.Scan() {
		b := bag(s.Bytes())
		left, right := b.left(), b.right()
		item, err := findCommonItem(left, right)
		if err != nil {
			return 0, err
		}
		sum += item.priority()
	}

	return sum, s.Err()
}

// part2 determines:
// for each group of three bags, which item is in all three bags?
func part2(r io.Reader) (int, error) {
	var count, sum int
	s := bufio.NewScanner(r)

	// we assume the file contains a multiple of 3 lines
	ok := s.Scan()
	for ok {
		// here we have to use s.Text() instead of s.Bytes() so that
		// we allocate new memory for each bag. Otherwise, the
		// subsequent scan operations will overwrite the memory.
		b1 := bag(s.Text())
		s.Scan()
		b2 := bag(s.Text())
		s.Scan()
		b3 := bag(s.Text())

		badge, err := findCommonItem(b1, b2, b3)
		if err != nil {
			return 0, err
		}
		sum += badge.priority()
		count++
		ok = s.Scan()
	}

	return sum, s.Err()
}

// a bag is a collection of items.
type bag []byte

// contains determines if this bag contains the given item.
func (b bag) contains(x item) bool {
	return bytes.Contains(b, []byte{byte(x)})
}

// left returns the bag's left compartment.
func (b bag) left() bag {
	return b[:len(b)/2]
}

// right returns the bag's right compartment.
func (b bag) right() bag {
	return b[len(b)/2:]
}

// an item is an uppercase or lowercase letter
type item byte

// priority gets the priority of this item.
func (x item) priority() int {
	// note that although 'a' follows 'Z',
	// the priority of 'a' is 1 and the priority of 'A' is 27.
	if x >= 'a' {
		return int(x - 'a' + 1)
	}
	return int(x - 'A' + 27)
}

// findCommonItem determines which item is in all of the given bags.
func findCommonItem(bags ...bag) (item, error) {
	var x item
	for x = 'A'; x <= 'z'; x++ {
		if allContain(x, bags...) {
			return x, nil
		}
	}

	return 0, errors.New("no duplicate found")
}

// allContain determines if all the given bags contain the given item.
func allContain(x item, bags ...bag) bool {
	for _, b := range bags {
		if !b.contains(x) {
			return false
		}
	}
	return true
}
