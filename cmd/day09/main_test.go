package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	const sample = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`
	got, err := solve(strings.NewReader(sample), 2)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	want := 13
	if got != want {
		t.Logf("part1(sample) = %d ; want %d", got, want)
		t.Fail()
	}
}

func TestPart2(t *testing.T) {
	const sample = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`

	got, err := solve(strings.NewReader(sample), 10)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	want := 36
	if got != want {
		t.Logf("part1(sample) = %d ; want %d", got, want)
		t.Fail()
	}
}
