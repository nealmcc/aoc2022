package main

import (
	"strings"
	"testing"

	v "github.com/nealmcc/aoc2022/pkg/vector/twod"
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

func TestText(t *testing.T) {
	t.Parallel()

	cave, err := read(strings.NewReader(_sample))
	if err != nil {
		t.Log("error reading sample:", err)
		t.FailNow()
	}

	got := cave.Text(494, 0, 504, 10)
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

	got, err := part1(cave, nil)
	if err != nil {
		t.Log("error reading sample:", err)
		t.FailNow()
	}
	want := 24
	if got != want {
		t.Logf("part1() = %d; want %d", got, want)
		t.Fail()
	}

	t.Log(cave.Text(488, -1, 513, 12))
}

func TestPart2(t *testing.T) {
	t.Parallel()

	cave, err := read(strings.NewReader(_sample))
	if err != nil {
		t.Log("error reading sample:", err)
		t.FailNow()
	}

	got, err := part2(cave, nil)
	if err != nil {
		t.Log("error reading sample:", err)
		t.FailNow()
	}

	want := 93
	if got != want {
		t.Logf("part2() = %d; want %d", got, want)
		t.Fail()
	}
	t.Log(cave.Text(488, -1, 513, 12))
}

func TestRender(t *testing.T) {
	t.Parallel()

	cave, err := read(strings.NewReader(_sample))
	if err != nil {
		t.Log("error reading sample:", err)
		t.FailNow()
	}

	min := v.Point{X: 488, Y: -1}
	max := v.Point{X: 513, Y: 12}
	r := NewRenderer(cave, "sample1", min, max, 10)

	_, err = part1(cave, r.SaveNext)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	r.prefix = "sample2"
	r.frame = 0

	_, err = part2(cave, r.SaveNext)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}
