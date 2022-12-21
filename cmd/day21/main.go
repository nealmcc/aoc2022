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

	p1, err := part1(tree, "root")
	if err != nil {
		log.Fatal("p1: ", err)
	}

	middle := time.Now()

	file.Seek(0, io.SeekStart)
	tree, err = parsetree(file)
	if err != nil {
		log.Fatal("parse: ", err)
	}

	p2, err := part2(tree)
	if err != nil {
		log.Fatalf("part2: %d", err)
	}

	end := time.Now()

	fmt.Printf("part 1: %d in %s\n", p1, middle.Sub(start))
	fmt.Printf("part 2: %d in %s\n", p2, end.Sub(middle))
}

type Tree map[string][]string

// part1 solves the given equation from the root node
func part1(tree Tree, lhs string) (int, error) {
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

func part2(tree Tree) (int, error) {
	expr := tree["root"]
	tree["root"] = []string{expr[0], "=", expr[2]}
	_, _, err := part2a(tree, "root")
	if err != nil {
		return 0, fmt.Errorf("part2A: %w", err)
	}

	// fmt.Println(inOrder(tree, "root"))
	p2, err := part2b(tree)
	if err != nil {
		return 0, fmt.Errorf("part2B: %w", err)
	}

	return p2, nil
}

// part2a modifies the tree so that a future traversal down from the
// root will have an easy job 'reversing' the equation.
func part2a(tree Tree, key string) (string, int, error) {
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

	lhs, left, err := part2a(tree, val[0])
	if err != nil {
		return "", 0, err
	}

	rhs, right, err := part2a(tree, val[2])
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

	// both sides are expressions:
	if len(lhs) > 0 && len(rhs) > 0 {
		return fmt.Sprintf("(%s %s %s)", lhs, val[1], rhs), 0, nil
	}

	// rhs is numeric
	if len(lhs) > 0 {
		return fmt.Sprintf("(%s %s %d)", lhs, val[1], right), 0, nil
	}

	switch val[1] {
	case "+", "*", "=":
		val[0], val[2] = val[2], val[0]
		tree[key] = val
		return fmt.Sprintf("(%s %s %d)", rhs, val[1], left), 0, nil

	case "-":
		// multiply the right value by neg1:
		newleft := val[2] + "_negative"
		tree["neg1"] = []string{"-1"}
		tree[newleft] = []string{val[2], "*", "neg1"}

		// add them together at this node, instead of subtracting:
		newright := val[0]
		val = []string{newleft, "+", newright}
		tree[key] = val

		// re-evaluate this node
		return part2a(tree, key)

	default:
		// all of the division operations already have a numeric RHS
		return fmt.Sprintf("(%d %s %s)", left, val[1], rhs), 0, nil
	}
}

// part2b walks the tree and reverses all of the operations as it goes,
// returning the value for "humn" at the end.
func part2b(tree Tree) (int, error) {
	readNext := func(k string) (keyLeft, op string, numRight int) {
		expr := tree[k]
		keyLeft, op, right := expr[0], expr[1], expr[2]
		rhs, ok := tree[right]
		if !ok {
			rhs = []string{right}
		}
		numRight, err := strconv.Atoi(rhs[0])
		if err != nil {
			panic(fmt.Sprintf("key %q: %s", k, err))
		}
		return
	}

	var (
		left, op      string
		right, answer int
	)
	left, _, answer = readNext("root")
	for left != "humn" {
		var next string
		next, op, right = readNext(left)
		switch op {
		case "/":
			answer *= right
		case "*":
			answer /= right
		case "+":
			answer -= right
		case "-":
			answer += right

		default:
			return 0, fmt.Errorf("unexpected op: %q", op)
		}
		left = next
	}

	return answer, nil
}

// inOrder prints the tree using an in-order traversal.
func inOrder(t Tree, key string) (string, error) {
	if key == "humn" {
		return key, nil
	}
	val, ok := t[key]
	if !ok {
		return key, nil
	}
	if len(val) == 1 {
		return val[0], nil
	}

	if len(val) != 3 {
		return "", fmt.Errorf("malformed value: %q", val)
	}

	lhs, err := inOrder(t, val[0])
	if err != nil {
		return "", fmt.Errorf("inOrder(%q): %w", key, err)
	}
	rhs, err := inOrder(t, val[2])
	if err != nil {
		return "", fmt.Errorf("inOrder(%q): %w", key, err)
	}

	return fmt.Sprintf("(%s %s %s)", lhs, val[1], rhs), nil
}

func (t Tree) String() string {
	s, err := inOrder(t, "root")
	if err != nil {
		panic(err)
	}
	return s
}

// gvnh: 3
// jfhb: cmbm + nlgl
var _re = regexp.MustCompile(`([0-9]+)|(([a-z]{4}) ([-+*\/]) ([a-z]{4}))`)

func parsetree(r io.Reader) (Tree, error) {
	s := bufio.NewScanner(r)

	tree := make(Tree)
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
