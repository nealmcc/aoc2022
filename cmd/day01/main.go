package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	start := time.Now()
	elves, err := readElves(file)
	if err != nil {
		log.Fatalf("read elves: %s", err)
	}

	sort.Slice(elves, func(i, j int) bool {
		return elves[i].calories > elves[j].calories
	})

	e1 := elves[0]
	e2 := elves[1]
	e3 := elves[2]

	p1 := e1.calories
	p2 := e1.calories + e2.calories + e3.calories

	end := time.Now()

	fmt.Println("part 1", p1)
	fmt.Println("part 2:", p2)
	fmt.Println("time taken:", end.Sub(start))
}

type elf struct {
	food     []int
	calories int
}

func readElves(r io.Reader) ([]elf, error) {
	elves := make([]elf, 0, 250)
	s := bufio.NewScanner(r)

	for {
		elf, ok, err := readElf(s)
		if err != nil {
			return nil, err
		}
		if !ok {
			return elves, nil
		}
		elves = append(elves, elf)
	}
}

func readElf(s *bufio.Scanner) (elf, bool, error) {
	e := elf{
		food:     make([]int, 0, 10),
		calories: 0,
	}

	var ok bool
	for s.Scan() {
		ok = true

		text := s.Text()
		if len(text) == 0 {
			break
		}

		n, err := strconv.Atoi(text)
		if err != nil {
			return elf{}, false, err
		}

		e.food = append(e.food, n)
		e.calories += n
	}

	return e, ok, s.Err()
}
