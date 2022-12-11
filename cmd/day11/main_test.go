package main

import (
	"testing"

	"github.com/nealmcc/aoc2022/pkg/collection"
)

func TestSolve(t *testing.T) {
	got := solve(sample())

	want := int64(2713310158)
	if got != want {
		t.Logf("part1(sample) = %d ; want %d", got, want)
		t.Fail()
	}
}

func sample() Troop {
	x := make(Troop, 4)
	for id := 0; id < 4; id++ {
		x[id] = Monkey{id: id}
	}

	x[0].items = collection.NewQueue(NewItem(79), NewItem(98))
	x[0].inspect = times(19)
	x[0].test = divisible(23)
	x[0].mTrue = &x[2]
	x[0].mFalse = &x[3]

	x[1].items = collection.NewQueue(NewItem(54), NewItem(65), NewItem(75), NewItem(74))
	x[1].inspect = plus(6)
	x[1].test = divisible(19)
	x[1].mTrue = &x[2]
	x[1].mFalse = &x[0]

	x[2].items = collection.NewQueue(NewItem(79), NewItem(60), NewItem(97))
	x[2].inspect = square
	x[2].test = divisible(13)
	x[2].mTrue = &x[1]
	x[2].mFalse = &x[3]

	x[3].items = collection.NewQueue(NewItem(74))
	x[3].inspect = plus(3)
	x[3].test = divisible(17)
	x[3].mTrue = &x[0]
	x[3].mFalse = &x[1]

	return x
}
