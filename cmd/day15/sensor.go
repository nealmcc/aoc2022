package main

import (
	v "github.com/nealmcc/aoc2022/pkg/vector/twod"
)

// Sensor represents a 'Manhattan circle' which has a center point and radius.
// It knows where its nearest beacon is.
type Sensor struct {
	Center v.Point
	Beacon v.Point
}

// radius calculates the radius of the sensor in manhattan distance
func (s Sensor) Radius() int {
	return v.ManhattanLength(s.Beacon.Sub(s.Center))
}

// SegmentAt returns the line segment of points where this circle
// intersects with the line y = n.
func (s Sensor) SegmentAt(y int) (Segment, bool) {
	rad := s.Radius()

	dy := y - s.Center.Y
	if dy < 0 {
		dy *= -1
	}

	if dy > rad {
		return Segment{}, false
	}

	dx := rad - dy
	seg := Segment{
		From: s.Center.X - dx,
		To:   s.Center.X + dx,
	}

	return seg, true
}
