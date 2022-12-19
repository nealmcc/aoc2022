package main

import (
	"strings"
	"testing"
)

const _sample = `Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.
Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian. `

func TestPart1_sample(t *testing.T) {
	t.Parallel()

	model, err := readInput(strings.NewReader(_sample))
	if err != nil {
		t.Log("error reading sample:", err)
		t.FailNow()
	}

	got, want := part1(model), 33
	if got != want {
		t.Logf("part1() = %d; want %d", got, want)
		t.Fail()
	}
}

var _p1 int

func BenchmarkPart1(b *testing.B) {
	// file, _ := os.Open("input.txt")
	model, _ := readInput(strings.NewReader(_sample))
	b.ResetTimer()
	var p1 int
	for i := 0; i < b.N; i++ {
		p1 = part1(model)
	}
	_p1 = p1
}
