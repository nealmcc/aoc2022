package main

import (
	"fmt"

	"github.com/nealmcc/aoc2022/pkg/collection"
)

// Monkey models a monkey from Advent of Code (2022, day 11)
type Monkey struct {
	id            int
	items         *collection.Queue[Item]
	test          testFunc
	inspect       inspectFunc
	mTrue, mFalse *Monkey
}

// LogFunc is a function that records each time a monkey inspects an item.
type LogFunc func(id int, item Item)

// these functions allow us to customise the monkey behaviour.
type (
	inspectFunc func(Item) Item
	testFunc    func(Item) bool
)

// Turn processes a single turn for this monkey, using the given logger to
// record the monkey's inspections.
func (m *Monkey) Turn(log LogFunc) {
	for m.items.Len() > 0 {
		x, _ := m.items.Pop()
		log(m.id, x)
		x = m.inspect(x)
		if m.test(x) {
			m.mTrue.items.Push(x)
		} else {
			m.mFalse.items.Push(x)
		}
	}
}

// Format implements fmt.Formatter.
func (m *Monkey) Format(s fmt.State, _ rune) {
	fmt.Fprintf(s, "Monkey %d:\n", m.id)
	fmt.Fprintf(s, "  items: %v\n", *m.items)
}

// String implements fmt.Stringer.
func (m *Monkey) String() string {
	return fmt.Sprintf("%v", m)
}

// Troop is a group of Monkeys.
type Troop []Monkey

// Round processes one turn for each monkey in the Troop, using the given log
// function to record each time a monkey inspects an item.
func (t Troop) Round(log LogFunc) {
	for _, m := range t {
		m.Turn(log)
	}
}

// times generates a monkey operation that multiples an item's worry level.
func times(n int) inspectFunc {
	return func(x Item) Item {
		return reduce(Item{
			div: n * x.div,
			mod: n * x.mod,
		})
	}
}

// plus generates a monkey operation that adds to an item's worry level.
func plus(n int) inspectFunc {
	return func(x Item) Item {
		return reduce(Item{
			div: x.div,
			mod: x.mod + n,
		})
	}
}

// square is an inspectFunc that multiples the value of the item by itself.
func square(x Item) Item {
	return reduce(Item{
		div: x.div*x.div + 2*x.div*x.mod,
		mod: x.mod * x.mod,
	})
}

// divisible generates a monkey test function to see if the item's worry level
// is evenly divisible by n.
func divisible(n int) testFunc {
	return func(x Item) bool { return x.mod%n == 0 }
}
