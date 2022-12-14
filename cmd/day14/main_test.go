package main

import (
	"strings"
	"testing"
)

const _sample = `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9
`

func TestRead(t *testing.T) {
	t.Parallel()

	cave, err := read(strings.NewReader(_sample))
	if err != nil {
		t.Log("error reading sample:", err)
		t.FailNow()
	}

	got, want := cave.Count(Rock), 20
	if got != want {
		t.Logf("cave.Count(Rock) =  %d; want %d", got, want)
		t.Fail()
	}
}

func TestRender(t *testing.T) {
	t.Parallel()

	cave, err := read(strings.NewReader(_sample))
	if err != nil {
		t.Log("error reading sample:", err)
		t.FailNow()
	}

	got := cave.Render(494, 0, 504, 10)
	want := `..........
..........
..........
..........
....#...##
....#...#.
..###...#.
........#.
........#.
#########.`
	if got != want {
		t.Logf("cave.Render():\n%s\nwant:\n%s\n", got, want)
		t.Fail()
	}
}

func TestPart1(t *testing.T) {
	t.Parallel()

	cave, err := read(strings.NewReader(_sample))
	if err != nil {
		t.Log("error reading sample:", err)
		t.FailNow()
	}

	got, want := part1(&cave), 24
	t.Log(cave.Render(488, -1, 513, 12))
	if got != want {
		t.Logf("part1() = %d; want %d", got, want)
		t.Fail()
	}
}

func TestPart2(t *testing.T) {
	t.Parallel()

	cave, err := read(strings.NewReader(_sample))
	if err != nil {
		t.Log("error reading sample:", err)
		t.FailNow()
	}

	got, want := part2(&cave), 93
	t.Log(cave.Render(488, -1, 513, 12))
	if got != want {
		t.Logf("part2() = %d; want %d", got, want)
		t.Fail()
	}
}
