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
	op            inspectFunc
	mTrue, mFalse *Monkey
}

// Item represents an item as its current worry level.
type Item = int

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
		x = m.op(x) / 3
		if m.test(x) {
			m.mTrue.catch(x)
		} else {
			m.mFalse.catch(x)
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

// catch asks this monkey to catch the given item.
func (m *Monkey) catch(x Item) {
	m.items.Push(x)
}

// Troop is a group of Monkeys.
type Troop []Monkey

// Round processes one turn for each monkey in the Troop.
func (t Troop) Round(log LogFunc) {
	for _, m := range t {
		m.Turn(log)
	}
}

// times generates a monkey operation that multiples an item's worry level.
func times(n int) inspectFunc {
	return func(x Item) Item { return x * n }
}

// plus generates a monkey operation that adds to an item's worry level.
func plus(n int) inspectFunc {
	return func(x Item) Item { return x + n }
}

// power generates a monkey operation that exponentiates an item's worry level.
// Works with powers >= 0. Does not check for overflow.
func power(n int) inspectFunc {
	return func(x Item) Item {
		pow := 1
		for i := 0; i < n; i++ {
			pow *= x
		}
		return pow
	}
}

// divisible generates a monkey test function to see if the item's worry level
// is evenly divisible by n.
func divisible(n int) testFunc {
	return func(x Item) bool { return x%n == 0 }
}
