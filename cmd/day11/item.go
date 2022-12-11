package main

// Item represents an item that the monkeys are tossing around.
// Each item's value is stored modulo 9699690 (2 * 3 * 5 * 7 * 11 * 13 * 17 * 19).
// The items integer value can by calculated as 9699690 * div + mod
type Item struct {
	div int
	mod int
}

const _lcm = 2 * 3 * 5 * 7 * 11 * 13 * 17 * 19 * 23

// NewItem converts an integer to an Item.
func NewItem(n int) Item {
	return Item{
		div: n / _lcm,
		mod: n % _lcm,
	}
}

// reduce returns a copy of the given item with the same value, expressed
// with the lowest positive modulo that it can.
func reduce(n Item) Item {
	return Item{
		div: n.div + n.mod/_lcm,
		mod: n.mod % _lcm,
	}
}
