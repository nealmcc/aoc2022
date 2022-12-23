package main

import (
	"os"
	"strings"
	"testing"

	"github.com/nealmcc/aoc2022/pkg/bound"
)

func TestPart1(t *testing.T) {
	file, err := os.Open("sample.txt")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	forest, err := parseInput(file)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	got, want := part1(&forest), 110

	if got != want {
		t.Logf("part1(sample) = %d ; want %d", got, want)
		t.Fail()
	}
}

func TestPart2(t *testing.T) {
	file, err := os.Open("sample.txt")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	forest, err := parseInput(file)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	part1(&forest)
	got, want := part2(&forest), 20

	if got != want {
		t.Logf("part2(sample) = %d ; want %d", got, want)
		t.Fail()
	}
}

func TestParseForest(t *testing.T) {
	tt := []struct {
		name   string
		in     string
		want   string
		bounds bound.Rect
	}{
		{
			name: "small sample",
			in: `	.....
					..##.
					..#..
					.....
					..##.`,
			want: `##
#.
..
##`,
		},
		{
			name: "larger sample",
			in: `	....#..
					..###.#
					#...#.#
					.#...##
					#.###..
					##.#.##
					.#..#..`,
			want: `....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..`,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			forest, err := parseInput(strings.NewReader(tc.in))
			if err != nil {
				t.Log(err)
				t.FailNow()
			}

			got := forest.String()
			if got != tc.want {
				t.Logf("parsed forest: \n%v\ngot:\n%v\nwant:\n%v\n",
					tc.in, got, tc.want)
				t.Fail()
			}
		})
	}
}
