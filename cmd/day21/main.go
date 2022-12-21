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

	start := time.Now()
	tree, err := parsetree(file)
	if err != nil {
		log.Fatalf("parse: %d", err)
	}

	var (
		eq1 = make(map[string][]string, len(tree))
		eq2 = make(map[string][]string, len(tree))
	)
	for k, v := range tree {
		eq1[k] = v
		eq2[k] = v
	}

	p1, err := part1(eq1, "root")
	if err != nil {
		log.Fatal("p1: ", err)
	}

	middle := time.Now()

	expr := eq2["root"]
	eq2["root"] = []string{expr[0], "=", expr[2]}
	p2, _, err := part2(eq2, "root")
	if err != nil {
		log.Fatal("p1: ", err)
	}

	end := time.Now()

	fmt.Printf("part 1: %d in %s\n", p1, middle.Sub(start))
	fmt.Printf("part 2: \n%v\n in %s\n", p2, end.Sub(middle))
}

// part1 solves the given equation from the root node
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

// part2 returns the string equation for the human (me) to solve:
func part2(tree map[string][]string, key string) (string, int, error) {
	val, ok := tree[key]
	if !ok {
		return "", 0, fmt.Errorf("node %s: not found", key)
	}

	if key == "humn" {
		return "humn", 0, nil
	}

	if len(val) == 1 {
		n, err := strconv.Atoi(val[0])
		return "", n, err
	}

	lhs, left, err := part2(tree, val[0])
	if err != nil {
		return "", 0, err
	}

	rhs, right, err := part2(tree, val[2])
	if err != nil {
		return "", 0, err
	}

	if len(lhs)+len(rhs) == 0 {
		var result int
		switch val[1] {
		case "+":
			result = left + right
			tree[key] = []string{strconv.Itoa(result)}
		case "-":
			result = left - right
			tree[key] = []string{strconv.Itoa(result)}
		case "*":
			result = left * right
			tree[key] = []string{strconv.Itoa(result)}
		case "/":
			result = left / right
			tree[key] = []string{strconv.Itoa(result)}
		default:
			return "", 0, fmt.Errorf("%q: invalid operation %q", key, val[1])
		}
		return "", result, nil
	}

	if len(lhs) > 0 && len(rhs) > 0 {
		return fmt.Sprintf("(%s %s %s)", lhs, val[1], rhs), 0, nil
	}

	if len(lhs) > 0 {
		return fmt.Sprintf("(%s %s %d)", lhs, val[1], right), 0, nil
	}

	return fmt.Sprintf("(%d %s %s)", left, val[1], rhs), 0, nil
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
