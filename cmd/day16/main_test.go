package main

import (
	"os"
	"strings"
	"testing"
)

func TestPart1_sample(t *testing.T) {
	t.Parallel()

	valves, err := ReadValves(strings.NewReader(_sample))
	if err != nil {
		t.Log("error reading sample:", err)
		t.FailNow()
	}

	got, want := part1(valves), 1651
	if got != want {
		t.Logf("part1() = %d; want %d", got, want)
		t.Fail()
	}
}

func TestPart2_sample(t *testing.T) {
	t.Parallel()

	valves, err := ReadValves(strings.NewReader(_sample))
	if err != nil {
		t.Log("error reading sample:", err)
		t.FailNow()
	}

	got, want := part2(valves), 1707
	if got != want {
		t.Logf("part2() = %d; want %d", got, want)
		t.Fail()
	}
}

func TestPart1_actual(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}
	t.Parallel()

	file, _ := os.Open("input.txt")
	valves, _ := ReadValves(file)

	got, want := part1(valves), 1659
	if got != want {
		t.Logf("part1() = %d; want %d", got, want)
		t.Fail()
	}
}

func TestPart2_actual(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}
	t.Parallel()

	file, _ := os.Open("input.txt")
	valves, _ := ReadValves(file)

	got, want := part2(valves), 2382
	if got != want {
		t.Logf("part2() = %d; want %d", got, want)
		t.Fail()
	}
}

var _p1 int

func BenchmarkPart1(b *testing.B) {
	file, _ := os.Open("input.txt")
	valves, _ := ReadValves(file)
	b.ResetTimer()
	var p1 int
	for i := 0; i < b.N; i++ {
		p1 = part1(valves)
	}
	_p1 = p1
}

var _p2 int

func BenchmarkPart2(b *testing.B) {
	file, _ := os.Open("input.txt")
	valves, _ := ReadValves(file)
	b.ResetTimer()
	var p2 int
	for i := 0; i < b.N; i++ {
		p2 = part2(valves)
	}
	_p2 = p2
}
