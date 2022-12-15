package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoinSegments(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name string
		in   []Segment
		want []Segment
	}{
		{
			name: "an empty list produces an empty list",
			in:   nil,
			want: nil,
		},
		{
			name: "a list of one segment should return itself",
			in: []Segment{
				{From: -2, To: 2},
			},
			want: []Segment{
				{From: -2, To: 2},
			},
		},
		{
			name: "two overlapping segments can be merged into one",
			in: []Segment{
				{From: -2, To: 2},
				{From: 0, To: 2},
			},
			want: []Segment{
				{From: -2, To: 2},
			},
		},
		{
			name: "two adjacing segments can be merged into one",
			in: []Segment{
				{From: -2, To: 2},
				{From: 3, To: 5},
			},
			want: []Segment{
				{From: -2, To: 5},
			},
		},
		{
			name: "two distinct segments remain distinct",
			in: []Segment{
				{From: -2, To: 1},
				{From: 3, To: 5},
			},
			want: []Segment{
				{From: -2, To: 1},
				{From: 3, To: 5},
			},
		},
		{
			name: "some segments are joined and others are not",
			in: []Segment{
				{From: -2, To: 2},
				{From: -1, To: 3},
				{From: 5, To: 8},
			},
			want: []Segment{
				{From: -2, To: 3},
				{From: 5, To: 8},
			},
		},
		{
			name: "handles segments in arbitrary order",
			in: []Segment{
				{From: -1, To: 3},
				{From: 5, To: 8},
				{From: -2, To: 2},
				{From: -12, To: -4},
			},
			want: []Segment{
				{From: -12, To: -4},
				{From: -2, To: 3},
				{From: 5, To: 8},
			},
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := JoinSegments(tc.in...)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestConstrain(t *testing.T) {
	t.Parallel()
	tt := []struct {
		name     string
		min, max int
		in       []Segment
		want     []Segment
	}{
		{
			name: "domain of length 1 => segment of length 1",
			min:  0,
			max:  0,
			in:   []Segment{{From: -1, To: 1}},
			want: []Segment{{From: 0, To: 0}},
		},
		{
			name: "domain of length 4, inputs overlap each end",
			min:  0,
			max:  3,
			in: []Segment{
				{From: -1, To: 1},
				{From: 3, To: 4},
			},
			want: []Segment{
				{From: 0, To: 1},
				{From: 3, To: 3},
			},
		},
		{
			name: "input outside the domain is discarded",
			min:  0,
			max:  3,
			in:   []Segment{{From: -10, To: -8}},
			want: []Segment{},
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := Constrain(tc.min, tc.max, tc.in)
			assert.Equal(t, tc.want, got)
		})
	}
}
