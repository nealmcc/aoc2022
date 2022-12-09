// Package twod models two-dimension vectors.
package twod

// Point is a 2-dimensional integer coordinate.
type Point struct {
	X int
	Y int
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

// gcd calculates the greatest common divisor of a and b.
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
