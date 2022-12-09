// Package rope models a rope from Advent of Code 2022, day 9.
//
// https://adventofcode.com/2022/day/9
package rope

import (
	"fmt"

	v "github.com/nealmcc/aoc2022/pkg/vector/twod"
)

// Logger is anything that records changes to a rope over time.
type Logger interface {
	Log(knots []v.Point)
}

// New creates a new Rope that uses the given Logger.
func New(size int, l Logger) *Rope {
	return &Rope{
		knots: make([]v.Point, size),
		l:     l,
	}
}

// Rope models a Rope from Advent of Code 2022, day 9.
type Rope struct {
	knots []v.Point
	l     Logger
}

var (
	_up    = v.Point{Y: 1}
	_right = v.Point{X: 1}
	_down  = v.Point{Y: -1}
	_left  = v.Point{X: -1}
)

// Move the head of the rope by the given increment, and adjust the tail to follow.
// Assumes that the step will always be one of up, down, left or right.
func (r *Rope) Move(step v.Point) error {
	if err := move(0, step, r.knots); err != nil {
		return err
	}

	r.l.Log(r.knots)
	return nil
}

func move(i int, step v.Point, knots []v.Point) error {
	// base case:
	if i == len(knots)-1 {
		knots[i] = knots[i].Add(step)
		return nil
	}

	head := &knots[i]
	next := knots[i+1]
	h2 := head.Add(step)
	*head = h2

	diff := h2.Sub(next)
	switch diff {
	// base case - no further movement needed:
	case v.Point{}, _up, _down, _left, _right,
		v.Point{X: 1, Y: 1}, v.Point{X: 1, Y: -1},
		v.Point{X: -1, Y: 1}, v.Point{X: -1, Y: -1}:
		return nil

	case v.Point{Y: 2}:
		step = _up

	case v.Point{X: 2}:
		step = _right

	case v.Point{Y: -2}:
		step = _down

	case v.Point{X: -2}:
		step = _left

	case v.Point{X: 2, Y: 1}, v.Point{X: 2, Y: 2}, v.Point{X: 1, Y: 2}:
		step = _up.Add(_right)

	case v.Point{X: 2, Y: -1}, v.Point{X: 2, Y: -2}, v.Point{X: 1, Y: -2}:
		step = _down.Add(_right)

	case v.Point{X: -2, Y: -1}, v.Point{X: -2, Y: -2}, v.Point{X: -1, Y: -2}:
		step = _down.Add(_left)

	case v.Point{X: -2, Y: 1}, v.Point{X: -2, Y: 2}, v.Point{X: -1, Y: 2}:
		step = _up.Add(_left)

	default:
		return fmt.Errorf("invalid step: %v gave difference of %v", step, diff)
	}

	return move(i+1, step, knots)
}
