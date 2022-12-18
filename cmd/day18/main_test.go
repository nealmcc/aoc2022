package main

import (
	"strings"
	"testing"
)

var _sample = `2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5`

func TestPart1(t *testing.T) {
	r := strings.NewReader(_sample)
	blocks, err := parseBlocks(r)
	if err != nil {
		t.Fatal(err)
	}

	got, want := part1(blocks), 64
	if got != want {
		t.Logf("part1(sample) = %d ; want %d", got, want)
		t.Fail()
	}
}

func TestPart2(t *testing.T) {
	r := strings.NewReader(_sample)
	sh, err := parseBlocks(r)
	if err != nil {
		t.Fatal(err)
	}

	got, want := part2(sh), 58
	if got != want {
		t.Logf("part2(sample) = %d ; want %d", got, want)
		t.Fail()
	}
}

func TestBoundsCheck(t *testing.T) {
	t.Parallel()
	bounds := map[point]int{
		{1, 0, 0}:  3, // most positive x value is 3
		{0, 1, 0}:  4,
		{0, 0, 1}:  5,
		{-1, 0, 0}: 3, // most negative x value is -3
		{0, -1, 0}: 4,
		{0, 0, -1}: 5,
	}

	tt := []struct {
		name string
		in   []point
		want bool
	}{
		{
			name: "the origin is inside all bounds",
			in:   []point{{}},
			want: true,
		},
		{
			name: "points within positive and negative quadrants",
			in: []point{
				{1, 2, 1},
				{-1, -2, -3},
			},
			want: true,
		},
		{
			name: "outside each positive axis",
			in: []point{
				{4, 1, 1},
				{1, 5, 1},
				{1, 1, 6},
			},
			want: false,
		},
		{
			name: "outside each negative axis",
			in: []point{
				{-4, 1, 1},
				{1, -5, 1},
				{1, 1, -6},
			},
			want: false,
		},
		{
			name: "exactly on a corner",
			in: []point{
				{-3, 4, -5},
			},
			want: true,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			for _, p := range tc.in {
				got := inbounds(bounds, p)
				if got != tc.want {
					t.Logf("inbounds(%v, %v) = %t ; want %t",
						bounds, tc.in, got, tc.want)
					t.Fail()
				}
			}
		})
	}
}
