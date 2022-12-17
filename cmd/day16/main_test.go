package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	t.Skip("not ready yet")
	t.Parallel()

	valves, err := read(strings.NewReader(_sample))
	if err != nil {
		t.Log("error reading sample:", err)
		t.FailNow()
	}

	got, want := part1(valves, 30), 1651
	if got != want {
		t.Logf("part1() = %d; want %d", got, want)
		t.Fail()
	}
}
