package main

import (
	"strings"
	"testing"
)

const _sample = `30373
25512
65332
33549
35390
`

func TestPart1(t *testing.T) {
	t.Parallel()

	trees, err := NewForest(strings.NewReader(_sample))
	if err != nil {
		t.Logf("error reading sample")
		t.FailNow()
	}

	got, want := trees.Visibility(), 21
	t.Logf("trees: \n%s", _sample)
	t.Logf("overall: \n%s", got)
	t.Logf("from the top: \n%s", got.Filter(top))
	t.Logf("from the right: \n%s", got.Filter(right))
	t.Logf("from the bottom: \n%s", got.Filter(bottom))
	t.Logf("from the left: \n%s", got.Filter(left))
	if len(got) != 21 {
		t.Logf("part1() = has length of %d; want %d", len(got), want)
		t.Fail()
	}
}

func TestPart2(t *testing.T) {
	t.Parallel()

	trees, err := NewForest(strings.NewReader(_sample))
	if err != nil {
		t.Logf("error reading sample")
		t.FailNow()
	}

	pos, sc := trees.SceneScore()
	wantPos := Pos{Row: 3, Col: 2}
	want := 8
	if pos != wantPos || sc != want {
		t.Logf("SceneScore() = %v, %d ; want %v, %d", pos, sc, wantPos, want)
		t.Fail()
	}
}
