package main

import (
	"os"
	"strings"
	"testing"

	v "github.com/nealmcc/aoc2022/pkg/vector/twod"
)

func TestPart1(t *testing.T) {
	file, err := os.Open("sample.txt")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	forest, path, err := parseInput(file)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	got, want := part1(forest, path), 6032

	if got != want {
		t.Logf("part1(sample) = %d ; want %d", got, want)
		t.Fail()
	}
}

func TestParseForest_grid(t *testing.T) {
	tt := []struct {
		name   string
		in     string
		origin v.Point
		max    v.Point
		grid   map[v.Point]byte
	}{
		{
			name: "basic 3x3 square forest",
			in: `..#
.##
###`,
			origin: v.Point{X: 0, Y: 0},
			max:    v.Point{X: 2, Y: 2},
			grid: map[v.Point]byte{
				{X: 0, Y: 0}: '.', {X: 1, Y: 0}: '.', {X: 2, Y: 0}: '#',
				{X: 0, Y: 1}: '.', {X: 1, Y: 1}: '#', {X: 2, Y: 1}: '#',
				{X: 0, Y: 2}: '#', {X: 1, Y: 2}: '#', {X: 2, Y: 2}: '#',
			},
		},
		{
			name: "sample forest",
			in: `	        ...#
					        .#..
					        #...
					        ....
					...#.......#
					........#...
					..#....#....
					..........#.
					        ...#....
					        .....#..
					        .#......
					        ......#.`,
			origin: v.Point{X: 8, Y: 0},
			max:    v.Point{X: 15, Y: 11},
			grid: map[v.Point]byte{
				// top section
				{X: 8, Y: 0}: '.', {X: 9, Y: 0}: '.', {X: 10, Y: 0}: '.', {X: 11, Y: 0}: '#',
				{X: 8, Y: 1}: '.', {X: 9, Y: 1}: '#', {X: 10, Y: 1}: '.', {X: 11, Y: 1}: '.',
				{X: 8, Y: 2}: '#', {X: 9, Y: 2}: '.', {X: 10, Y: 2}: '.', {X: 11, Y: 2}: '.',
				{X: 8, Y: 3}: '.', {X: 9, Y: 3}: '.', {X: 10, Y: 3}: '.', {X: 11, Y: 3}: '.',
				// middle section
				{X: 0, Y: 4}: '.', {X: 1, Y: 4}: '.', {X: 2, Y: 4}: '.', {X: 3, Y: 4}: '#', {X: 4, Y: 4}: '.', {X: 5, Y: 4}: '.', {X: 6, Y: 4}: '.', {X: 7, Y: 4}: '.', {X: 8, Y: 4}: '.', {X: 9, Y: 4}: '.', {X: 10, Y: 4}: '.', {X: 11, Y: 4}: '#',
				{X: 0, Y: 5}: '.', {X: 1, Y: 5}: '.', {X: 2, Y: 5}: '.', {X: 3, Y: 5}: '.', {X: 4, Y: 5}: '.', {X: 5, Y: 5}: '.', {X: 6, Y: 5}: '.', {X: 7, Y: 5}: '.', {X: 8, Y: 5}: '#', {X: 9, Y: 5}: '.', {X: 10, Y: 5}: '.', {X: 11, Y: 5}: '.',
				{X: 0, Y: 6}: '.', {X: 1, Y: 6}: '.', {X: 2, Y: 6}: '#', {X: 3, Y: 6}: '.', {X: 4, Y: 6}: '.', {X: 5, Y: 6}: '.', {X: 6, Y: 6}: '.', {X: 7, Y: 6}: '#', {X: 8, Y: 6}: '.', {X: 9, Y: 6}: '.', {X: 10, Y: 6}: '.', {X: 11, Y: 6}: '.',
				{X: 0, Y: 7}: '.', {X: 1, Y: 7}: '.', {X: 2, Y: 7}: '.', {X: 3, Y: 7}: '.', {X: 4, Y: 7}: '.', {X: 5, Y: 7}: '.', {X: 6, Y: 7}: '.', {X: 7, Y: 7}: '.', {X: 8, Y: 7}: '.', {X: 9, Y: 7}: '.', {X: 10, Y: 7}: '#', {X: 11, Y: 7}: '.',
				// bottom section
				{X: 8, Y: 8}: '.', {X: 9, Y: 8}: '.', {X: 10, Y: 8}: '.', {X: 11, Y: 8}: '#', {X: 12, Y: 8}: '.', {X: 13, Y: 8}: '.', {X: 14, Y: 8}: '.', {X: 15, Y: 8}: '.',
				{X: 8, Y: 9}: '.', {X: 9, Y: 9}: '.', {X: 10, Y: 9}: '.', {X: 11, Y: 9}: '.', {X: 12, Y: 9}: '.', {X: 13, Y: 9}: '#', {X: 14, Y: 9}: '.', {X: 15, Y: 9}: '.',
				{X: 8, Y: 10}: '.', {X: 9, Y: 10}: '#', {X: 10, Y: 10}: '.', {X: 11, Y: 10}: '.', {X: 12, Y: 10}: '.', {X: 13, Y: 10}: '.', {X: 14, Y: 10}: '.', {X: 15, Y: 10}: '.',
				{X: 8, Y: 11}: '.', {X: 9, Y: 11}: '.', {X: 10, Y: 11}: '.', {X: 11, Y: 11}: '.', {X: 12, Y: 11}: '.', {X: 13, Y: 11}: '.', {X: 14, Y: 11}: '#', {X: 15, Y: 11}: '.',
			},
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			f, _, err := parseInput(strings.NewReader(tc.in))
			if err != nil {
				t.Log(err)
				t.FailNow()
			}

			if f.origin != tc.origin {
				t.Logf("got origin = %v; want %v", f.origin, tc.origin)
				t.Fail()
			}

			if f.max != tc.max {
				t.Logf("got max = %v; want %v", f.max, tc.max)
				t.Fail()
			}

			for y := 0; y < 24; y++ {
				for x := 0; x < 24; x++ {
					p := v.Point{X: x, Y: y}
					gotByte, gotOK := f.grid[p]
					wantByte, wantOK := tc.grid[p]
					if gotOK != wantOK || gotByte != wantByte {
						t.Logf("grid[%d, %d] = %q, %t; want %q, %t",
							x, y, gotByte, gotOK, wantByte, wantOK)
						t.Fail()
					}
				}
			}
		})
	}
}
