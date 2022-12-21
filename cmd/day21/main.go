package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { file.Close() }()

	equation, err := parsetree(file)
	if err != nil {
		log.Fatalf("parse: %d", err)
	}

	start := time.Now()
	p1, err := part1(equation, "root")
	if err != nil {
		log.Fatal("parse: ", err)
	}
	middle := time.Now()

	// p2 := part2(equation)
	// end := time.Now()

	fmt.Printf("part 1: %d in %s\n", p1, middle.Sub(start))
	// fmt.Printf("part 2: %d in %s\n", p2, end.Sub(middle))
}

// part1 solves the given equation from the given node
func part1(tree map[string][]string, lhs string) (int, error) {
	rhs, ok := tree[lhs]
	if !ok {
		return 0, fmt.Errorf("node %s: not found", lhs)
	}

	if len(rhs) == 1 {
		return strconv.Atoi(rhs[0])
	}

	left, err := part1(tree, rhs[0])
	if err != nil {
		return 0, err
	}

	right, err := part1(tree, rhs[2])
	if err != nil {
		return 0, err
	}

	var result int
	switch rhs[1] {
	case "+":
		result = left + right
		tree[lhs] = []string{strconv.Itoa(result)}
	case "-":
		result = left - right
		tree[lhs] = []string{strconv.Itoa(result)}
	case "*":
		result = left * right
		tree[lhs] = []string{strconv.Itoa(result)}
	case "/":
		result = left / right
		tree[lhs] = []string{strconv.Itoa(result)}
	default:
		return 0, fmt.Errorf("%q: invalid operation %q", lhs, rhs[1])
	}

	return result, nil
}

// gvnh: 3
// jfhb: cmbm + nlgl
var _re = regexp.MustCompile(`([0-9]+)|(([a-z]{4}) ([-+*\/]) ([a-z]{4}))`)

func parsetree(r io.Reader) (map[string][]string, error) {
	s := bufio.NewScanner(r)

	tree := make(map[string][]string)
	for s.Scan() {
		line := s.Text()
		lhs := line[:4]
		parts := _re.FindAllStringSubmatch(line[5:], -1)
		if len(parts[0][1]) > 0 {
			tree[lhs] = []string{parts[0][1]}
		} else {
			tree[lhs] = parts[0][3:]
		}
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return tree, nil
}
