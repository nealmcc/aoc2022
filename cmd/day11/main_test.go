package main

import (
	"fmt"
	"testing"

	"github.com/nealmcc/aoc2022/pkg/collection"
)

func TestPart1(t *testing.T) {
	got := part1(sample())

	want := int64(10605)
	if got != want {
		t.Logf("part1(sample) = %d ; want %d", got, want)
		t.Fail()
	}
}

func TestPower(t *testing.T) {
	tt := []struct {
		base, pow, want int
	}{
		{base: 0, pow: 0, want: 1},
		{base: 0, pow: 1, want: 0},
		{base: 0, pow: 2, want: 0},
		{base: 1, pow: 0, want: 1},
		{base: 3, pow: 0, want: 1},
		{base: 3, pow: 1, want: 3},
		{base: 3, pow: 2, want: 9},
		{base: 3, pow: 4, want: 81},
	}

	for _, tc := range tt {
		tc := tc
		name := fmt.Sprintf("%d ^ %d = %d", tc.base, tc.pow, tc.want)
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			fn := power(tc.pow)
			if got := fn(tc.base); got != tc.want {
				t.Logf("power(%d) of %d = %d ; want %d",
					tc.pow, tc.base, got, tc.want)
				t.Fail()
			}
		})
	}
}

func sample() Troop {
	x := make(Troop, 4)
	for id := 0; id < 4; id++ {
		x[id] = Monkey{id: id}
	}

	x[0].items = collection.NewQueue(79, 98)
	x[0].op = times(19)
	x[0].test = divisible(23)
	x[0].mTrue = &x[2]
	x[0].mFalse = &x[3]

	x[1].items = collection.NewQueue(54, 65, 75, 74)
	x[1].op = plus(6)
	x[1].test = divisible(19)
	x[1].mTrue = &x[2]
	x[1].mFalse = &x[0]

	x[2].items = collection.NewQueue(79, 60, 97)
	x[2].op = power(2)
	x[2].test = divisible(13)
	x[2].mTrue = &x[1]
	x[2].mFalse = &x[3]

	x[3].items = collection.NewQueue(74)
	x[3].op = plus(3)
	x[3].test = divisible(17)
	x[3].mTrue = &x[0]
	x[3].mFalse = &x[1]

	return x
}
