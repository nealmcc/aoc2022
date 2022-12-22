// Package twod models two-dimension vectors.
package twod

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"

	"github.com/nealmcc/aoc2022/pkg/vector"
)

// Point is a 2-dimensional integer coordinate.
type Point struct {
	X int
	Y int
}

// parse the given text in the form x,y as a point.
// The text must have two numbers separated by a comma.
func (p *Point) Parse(b []byte) error {
	parts := bytes.Split(b, []byte{','})
	if len(parts) != 2 {
		return errors.New("parse requires two parts")
	}

	var err error
	if (*p).X, err = strconv.Atoi(string(parts[0])); err != nil {
		return fmt.Errorf("invalid value for x: %w", err)
	}

	if (*p).Y, err = strconv.Atoi(string(parts[1])); err != nil {
		return fmt.Errorf("invalid value for y: %w", err)
	}

	return nil
}

// Add returns the vector sum of a + b.
func (a Point) Add(b Point) Point {
	return Point{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

// Sub returns the vector difference of a - b.
func (a Point) Sub(b Point) Point {
	return Point{
		X: a.X - b.X,
		Y: a.Y - b.Y,
	}
}

// Times returns a copy of this Point scaled by n.
func (a Point) Times(n int) Point {
	return Point{
		X: a.X * n,
		Y: a.Y * n,
	}
}

// Neighbours4 returns the four points adjacent to this one.
func (p Point) Neighbours4() []Point {
	return []Point{
		{X: p.X, Y: p.Y - 1},
		{X: p.X, Y: p.Y + 1},
		{X: p.X - 1, Y: p.Y},
		{X: p.X + 1, Y: p.Y},
	}
}

// Reduce returns the shortest vector with the same slope as this one
// that can still be represented with integer values for X and Y.
// Also returns the largest positive integer that evenly divides this one.
func (a Point) Reduce() (Point, int) {
	if (a == Point{}) {
		return a, 1
	}

	if a.X == 0 {
		if a.Y > 0 {
			return Point{X: 0, Y: 1}, a.Y
		}
		return Point{X: 0, Y: -1}, -1 * a.Y
	}

	if a.Y == 0 {
		if a.X > 0 {
			return Point{X: 1, Y: 0}, a.X
		}
		return Point{X: -1, Y: 0}, -1 * a.X
	}

	scale := gcd(a.X, a.Y)
	if scale < 0 {
		scale *= -1
	}
	return Point{
		X: a.X / scale,
		Y: a.Y / scale,
	}, scale
}

func (a Point) String() string {
	return fmt.Sprintf("(%d, %d)", a.X, a.Y)
}

func ManhattanLength(p Point) int {
	if p.X < 0 {
		p.X *= -1
	}
	if p.Y < 0 {
		p.Y *= -1
	}
	return p.X + p.Y
}

// Rot90 returns this vector rotated 90 degrees.
// With Y up, this is Left. With Y down, this is Light.
func (p Point) Rot90() Point {
	m := vector.Matrix{
		{0, -1},
		{1, 0},
	}

	// express this point as a 2x1 column matrix and then apply the
	// matrix transformation
	v, _ := m.Cross(vector.Matrix{
		{p.X},
		{p.Y},
	})

	return Point{X: v[0][0], Y: v[1][0]}
}

// Rot270 returns this vector rotated by 270 degrees.
// With Y up, this is Right. With Y down, this is Left.
func (p Point) Rot270() Point {
	m := vector.Matrix{
		{0, 1},
		{-1, 0},
	}

	v, _ := m.Cross(vector.Matrix{
		{p.X},
		{p.Y},
	})

	return Point{X: v[0][0], Y: v[1][0]}
}

// gcd calculates the greatest common divisor of a and b.
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
