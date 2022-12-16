package main

import "fmt"

// Valve is one node in a network. Each Valve has a unique ID, some neighbours,
// and a flow rate of "pressure per minute" that it can release.
type Valve struct {
	ID         ValveID
	Flow       int
	Neighbours []ValveID
}

// ValveID is an alias for a 16-bit unsigned int.
// It is just the two bytes of the string stored side by side in a single value.
type ValveID uint16

// ID converts the given two-character string to a valve ID.
// Input is assumed to be valid (not checked).
func ID(s string) ValveID {
	sum := ValveID(s[0])
	sum = sum << 8
	sum += ValveID(s[1])
	return sum
}

// Format implements fmt.Formatter.
func (id ValveID) Format(s fmt.State, verb rune) {
	first := byte((id & 0xFF00) >> 8)
	second := byte(id & 0x00FF)
	s.Write([]byte{first, second})
}

// String implements fmt.Stringer.
func (id ValveID) String() string {
	return fmt.Sprintf("%v", id)
}
