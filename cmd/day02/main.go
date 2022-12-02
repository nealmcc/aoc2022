package main

import (
	"bufio"
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
	p1, err := addScores(file, part1)
	if err != nil {
		log.Fatalf("part1: %s", err)
	}

	middle := time.Now()
	file.Seek(0, 0)

	p2, err := addScores(file, part2)
	if err != nil {
		log.Fatalf("part2: %s", err)
	}
	end := time.Now()

	fmt.Printf("part 1: %d in %s\n", p1, middle.Sub(start))
	fmt.Printf("part 2: %d in %s\n", p2, end.Sub(middle))
}

func addScores(r io.Reader, score func(row string) (int, error)) (int, error) {
	var sum int
	s := bufio.NewScanner(r)

	for s.Scan() {
		row := s.Text()
		n, err := score(row)
		if err != nil {
			return 0, err
		}
		sum += n
	}

	return sum, s.Err()
}

// part1 is the scorer function for part 1
func part1(s string) (int, error) {
	switch s {
	// I choose Rock:
	case "A X":
		return 1 + 3, nil
	case "B X":
		return 1 + 0, nil
	case "C X":
		return 1 + 6, nil

	// I choose Paper:
	case "A Y":
		return 2 + 6, nil
	case "B Y":
		return 2 + 3, nil
	case "C Y":
		return 2 + 0, nil

	// I choose Scissors:
	case "A Z":
		return 3 + 0, nil
	case "B Z":
		return 3 + 6, nil
	case "C Z":
		return 3 + 3, nil

	default:
		return 0, errors.New("invalid input")
	}
}

// part2 is the scorer function for part 2
func part2(s string) (int, error) {
	switch s {
	// I lose:
	case "A X":
		return 3 + 0, nil
	case "B X":
		return 1 + 0, nil
	case "C X":
		return 2 + 0, nil

	// I draw:
	case "A Y":
		return 1 + 3, nil
	case "B Y":
		return 2 + 3, nil
	case "C Y":
		return 3 + 3, nil

	// I win:
	case "A Z":
		return 2 + 6, nil
	case "B Z":
		return 3 + 6, nil
	case "C Z":
		return 1 + 6, nil

	default:
		return 0, errors.New("invalid input")
	}
}
