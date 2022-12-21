package main

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	file, err := os.Open("sample.txt")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	tree, err := parsetree(file)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	got, err := part1(tree, "root")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	want := 152

	if got != want {
		t.Logf("part1(sample) = %d ; want %d", got, want)
		t.Fail()
	}
}

func TestPart2(t *testing.T) {
	file, err := os.Open("sample.txt")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	tree, err := parsetree(file)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	got, err := part2(tree)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	want := 301
	if got != want {
		t.Logf("part2(sample) = %d ; want %d", got, want)
		t.Fail()
	}
}
