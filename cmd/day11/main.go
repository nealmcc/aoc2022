package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/nealmcc/aoc2022/pkg/collection"
)

func main() {
	start := time.Now()

	p1 := part1(barrel())
	middle := time.Now()

	// p2, err := part2()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// end := time.Now()

	fmt.Printf("part 1: %d in %s\n", p1, middle.Sub(start))
	// fmt.Printf("part 2: %d in %s\n", p2, end.Sub(middle))
}

func part1(t Troop) int64 {
	active := make([]int, len(t))
	logFn := func(id int, x Item) {
		active[id]++
	}

	for i := 0; i < 20; i++ {
		t.Round(logFn)
	}

	sort.Slice(active, func(i, j int) bool {
		return active[i] > active[j]
	})
	fmt.Println(active)
	return int64(active[0] * active[1])
}

func barrel() Troop {
	x := make(Troop, 8)
	for id := 0; id < 8; id++ {
		x[id] = Monkey{id: id}
	}

	// Monkey 0:
	x[0].items = collection.NewQueue(Item(72), Item(97))
	x[0].op = times(13)
	x[0].test = divisible(19)
	x[0].mTrue = &x[5]
	x[0].mFalse = &x[6]

	// Monkey 1:
	x[1].items = collection.NewQueue(Item(55), Item(70), Item(90), Item(74), Item(95))
	x[1].op = power(2)
	x[1].test = divisible(7)
	x[1].mTrue = &x[5]
	x[1].mFalse = &x[0]

	// Monkey 2:
	x[2].items = collection.NewQueue(Item(74), Item(97), Item(66), Item(57))
	x[2].op = plus(6)
	x[2].test = divisible(17)
	x[2].mTrue = &x[1]
	x[2].mFalse = &x[0]

	// Monkey 3:
	x[3].items = collection.NewQueue(Item(86), Item(54), Item(53))
	x[3].op = plus(2)
	x[3].test = divisible(13)
	x[3].mTrue = &x[1]
	x[3].mFalse = &x[2]

	// Monkey 4:
	x[4].items = collection.NewQueue(Item(50), Item(65), Item(78), Item(50), Item(62), Item(99))
	x[4].op = plus(3)
	x[4].test = divisible(11)
	x[4].mTrue = &x[3]
	x[4].mFalse = &x[7]

	// Monkey 5:
	x[5].items = collection.NewQueue(Item(90))
	x[5].op = plus(4)
	x[5].test = divisible(2)
	x[5].mTrue = &x[4]
	x[5].mFalse = &x[6]

	// Monkey 6:
	x[6].items = collection.NewQueue(Item(88), Item(92), Item(63), Item(94), Item(96), Item(82), Item(53), Item(53))
	x[6].op = plus(8)
	x[6].test = divisible(5)
	x[6].mTrue = &x[4]
	x[6].mFalse = &x[7]

	// Monkey 7:
	x[7].items = collection.NewQueue(Item(70), Item(60), Item(71), Item(69), Item(77), Item(70), Item(98))
	x[7].op = times(7)
	x[7].test = divisible(3)
	x[7].mTrue = &x[2]
	x[7].mFalse = &x[3]

	return x
}
