package main

import (
	"sort"
)

// Segment is a horizontal or vertical line segment
// that includes all points in the range [From, To].
type Segment struct {
	From int
	To   int
}

// Length returns the number of values in this segment
func (s Segment) Length() int {
	return s.To - s.From + 1
}

// JoinSegments takes the given list of segments, and merges them together,
// removing any overlapping x values, and minimizing the number of resulting
// segments.  Assumes all of the segments have the same y value.
func JoinSegments(segments ...Segment) []Segment {
	if len(segments) < 2 {
		return segments
	}

	sort.Slice(segments, func(i, j int) bool {
		diff := segments[i].From - segments[j].From
		if diff != 0 {
			return diff < 0
		}
		diff = segments[i].To - segments[j].To
		return diff < 0
	})

	result := make([]Segment, 0, 4)

	curr := segments[0]
	for i := 1; i < len(segments); i++ {
		next := segments[i]
		if curr.To+1 >= next.From {
			if curr.To >= next.To {
				continue
			}
			curr.To = next.To
			continue
		}
		result = append(result, curr)
		curr = next
	}
	result = append(result, curr)
	return result
}

// Constrain modifies a collection of segments so that for all Segments s in the
// results, it be true that min <= s.From and (s.From+s.Len) <= max.
// If a given segment was fully outside this range, then it will be discarded.
func Constrain(min, max int, segments []Segment) []Segment {
	result := make([]Segment, 0, 4)
	for _, s := range segments {
		if s.To < min || s.From > max {
			continue
		}

		if s.From < min {
			s.From = min
		}

		if s.To > max {
			s.To = max
		}

		result = append(result, s)
	}
	return result
}
