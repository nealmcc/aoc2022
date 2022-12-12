package main

import (
	"os"
	"strings"
	"testing"

	"github.com/nealmcc/aoc2022/pkg/vector/twod"
	"github.com/stretchr/testify/assert"
)

const _sample = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
`

func TestRead(t *testing.T) {
	t.Parallel()

	hill, err := read(strings.NewReader(_sample))
	if err != nil {
		t.Log("error reading sample", err)
		t.FailNow()
	}

	a := assert.New(t)
	a.Equal(8, hill.size)
	a.Equal(twod.Point{}, hill.start)
	a.Equal(twod.Point{X: 5, Y: 2}, hill.end)
}

func TestPart1(t *testing.T) {
	t.Parallel()

	hill, err := read(strings.NewReader(_sample))
	if err != nil {
		t.Log("error reading sample", err)
		t.FailNow()
	}

	got, want := part1(hill), 31

	if got != want {
		t.Logf("part1() =  %d; want %d", got, want)
		t.Fail()
	}
}

func TestPart2(t *testing.T) {
	t.Parallel()

	hill, err := read(strings.NewReader(_sample))
	if err != nil {
		t.Log("error reading sample", err)
		t.FailNow()
	}

	got, want := part2(hill), 29

	if got != want {
		t.Logf("part2() =  %d; want %d", got, want)
		t.Fail()
	}
}

var _result int // prevent the compiler from optimising away the call.

func BenchmarkPart2(b *testing.B) {
	file, err := os.Open("input.txt")
	if err != nil {
		b.Fatal(err)
	}
	hill, err := read(file)
	if err != nil {
		b.Fatal(err)
	}

	var result int
	for n := 0; n < b.N; n++ {
		result = part2(hill)
	}
	_result = result
}
