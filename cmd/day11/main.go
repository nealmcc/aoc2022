package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/nealmcc/aoc2022/pkg/collection"
)

func main() {
	start := time.Now()
	answer := solve(barrel())
	end := time.Now()
	fmt.Printf("solution: %d in %s\n", answer, end.Sub(start))
}

func solve(t Troop) int64 {
	active := make([]int, len(t))
	logFn := func(id int, _ Item) {
		active[id]++
	}

	for i := 0; i < 10000; i++ {
		t.Round(logFn)
	}

	sort.Slice(active, func(i, j int) bool {
		return active[i] > active[j]
	})
	return int64(active[0] * active[1])
}

// barrel initializes the barrel of monkeys.
// I've hard-coded this instead of parsing the input, since there are only 8 monkeys,
// and my solution is based on the specific prime divisors that are in my input.
//
// Now that I know what part 2 is, I *could* write a parser to read the input,
// and calculate a dynamic LCM for the items, but... I've already solved the
// problem, and it doesn't feel worth it.
func barrel() Troop {
	x := make(Troop, 8)
	for id := 0; id < 8; id++ {
		x[id] = Monkey{id: id}
	}

	// Monkey 0:
	x[0].items = collection.NewQueue(NewItem(72), NewItem(97))
	x[0].inspect = times(13)
	x[0].test = divisible(19)
	x[0].mTrue = &x[5]
	x[0].mFalse = &x[6]

	// Monkey 1:
	x[1].items = collection.NewQueue(NewItem(55), NewItem(70), NewItem(90), NewItem(74), NewItem(95))
	x[1].inspect = square
	x[1].test = divisible(7)
	x[1].mTrue = &x[5]
	x[1].mFalse = &x[0]

	// Monkey 2:
	x[2].items = collection.NewQueue(NewItem(74), NewItem(97), NewItem(66), NewItem(57))
	x[2].inspect = plus(6)
	x[2].test = divisible(17)
	x[2].mTrue = &x[1]
	x[2].mFalse = &x[0]

	// Monkey 3:
	x[3].items = collection.NewQueue(NewItem(86), NewItem(54), NewItem(53))
	x[3].inspect = plus(2)
	x[3].test = divisible(13)
	x[3].mTrue = &x[1]
	x[3].mFalse = &x[2]

	// Monkey 4:
	x[4].items = collection.NewQueue(NewItem(50), NewItem(65), NewItem(78), NewItem(50), NewItem(62), NewItem(99))
	x[4].inspect = plus(3)
	x[4].test = divisible(11)
	x[4].mTrue = &x[3]
	x[4].mFalse = &x[7]

	// Monkey 5:
	x[5].items = collection.NewQueue(NewItem(90))
	x[5].inspect = plus(4)
	x[5].test = divisible(2)
	x[5].mTrue = &x[4]
	x[5].mFalse = &x[6]

	// Monkey 6:
	x[6].items = collection.NewQueue(NewItem(88), NewItem(92), NewItem(63), NewItem(94), NewItem(96), NewItem(82), NewItem(53), NewItem(53))
	x[6].inspect = plus(8)
	x[6].test = divisible(5)
	x[6].mTrue = &x[4]
	x[6].mFalse = &x[7]

	// Monkey 7:
	x[7].items = collection.NewQueue(NewItem(70), NewItem(60), NewItem(71), NewItem(69), NewItem(77), NewItem(70), NewItem(98))
	x[7].inspect = times(7)
	x[7].test = divisible(3)
	x[7].mTrue = &x[2]
	x[7].mFalse = &x[3]

	return x
}
