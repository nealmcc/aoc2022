package main

import (
	"os"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
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
