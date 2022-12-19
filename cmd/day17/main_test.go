package main

import (
	"context"
	"os"
	"strings"
	"testing"
)

const _sample = `>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>`

func TestPart1(t *testing.T) {
	t.Parallel()
	ctx := withInterrupt(context.Background())
	got, want := part1(ctx, _sample, os.Stderr), 3068
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
	t.Parallel()

	tt := []struct {
		name string
		in   string
		want Shape
	}{
		{
			name: "empty",
			in:   ``,
			want: Shape{},
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
			want: Shape{0, 0, 0, 1},
		},
		{
			name: "top left",
			in:   `x`,
			want: Shape{0x40, 0, 0, 0},
		},
		{
			name: "all filled",
			in: `#######
#######
#######
#######`,
			want: Shape{0x7F, 0x7F, 0x7F, 0x7F},
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

func TestBoard_WriteTo(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name string
		in   Board
		want string
	}{
		{
			name: "empty board",
			want: "+-------+",
		},
		{
			name: "board with height 3",
			in: Board{
				top:  2,
				rows: make([]Row, 3),
			},
			want: `|.......|
|.......|
|.......|
+-------+`,
		},
		{
			name: "a board with its first piece placed:",
			in: Board{
				top:  2,
				rows: make([]Row, 3),
			},
			want: `|.......|
|.......|
|.......|
+-------+`,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			var b strings.Builder
			_, err := tc.in.WriteTo(&b)
			if err != nil {
				t.Log(err)
				t.FailNow()
			}

			got := b.String()
			t.Log(got)
			if got != tc.want {
				t.Logf("%#v.WriteTo() got\n%q\nwant:\n%q", tc.in, got, tc.want)
				t.Fail()
			}
		})
	}
}
