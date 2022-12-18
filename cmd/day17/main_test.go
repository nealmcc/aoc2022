package main

import (
	"testing"
)

const _sample = `>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>`

func TestPart1(t *testing.T) {
	t.Parallel()

	got, want := part1(_sample), 3068
	if got != want {
		t.Logf("part1() = %d; want %d", got, want)
		t.Fail()
	}
}

func TestGenerator(t *testing.T) {
	t.Parallel()

	next := generator(0, []int{1, 2, 3})
	for i := 0; i < 10; i++ {
		got, want := next(), i%3+1
		if got != want {
			t.Logf("%d * next() = %d ; want %d", i, got, want)
			t.Fail()
		}
	}
}

func TestMakeShape(t *testing.T) {
	tt := []struct {
		name string
		in   string
		want shape
	}{
		{
			name: "empty",
			in:   ``,
			want: shape{},
		},
		{
			name: "all dots",
			in: `.......
.......
.......
.......
`,
		},
		{
			name: "bottom right",
			in: `.......
.......
.......
......#
`,
			want: shape{0, 0, 0, 1},
		},
		{
			name: "top left",
			in:   `x`,
			want: shape{0x40, 0, 0, 0},
		},
		{
			name: "all filled",
			in: `#######
#######
#######
#######`,
			want: shape{0x7F, 0x7F, 0x7F, 0x7F},
		},
		{
			name: "dash",
			in: `.......
.......
.......
..####.`,
			want: _dash,
		},
		{
			name: "plus",
			in: `
...#...
..###..
...#...`,
			want: _plus,
		},
		{
			name: "corner",
			in: `
....#..
....#..
..###..`,
			want: _corner,
		},
		{
			name: "bar",
			in: `..#....
..#....
..#....
..#....`,
			want: _bar,
		},
		{
			name: "square",
			in: `
.......
..##...
..##...`,
			want: _square,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := makeShape(tc.in)
			if got != tc.want {
				t.Logf("makeShape(%s) = %v ; want %v", tc.in, got, tc.want)
				t.Fail()
			}
		})
	}
}
