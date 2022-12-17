package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
)

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

var _re = regexp.MustCompile(`Valve ([A-Z][A-Z]) has flow rate=([\d]+); tunnels? leads? to valves? ((?:[A-Z][A-Z])(?:(?:, )?(?:[A-Z][A-Z]))*)`)

func ParseValve(b []byte) (Valve, error) {
	m := _re.FindAllSubmatch(b, -1)
	if len(m) != 1 {
		return Valve{}, fmt.Errorf("parse(%q): want 1 match got %d",
			b, len(m))
	}

	flow, err := strconv.Atoi(string(m[0][2]))
	if err != nil {
		return Valve{}, fmt.Errorf("parse flow %q: %w", m[0][2], err)
	}

	v := Valve{
		ID:         ID(string(m[0][1])),
		Flow:       flow,
		Neighbours: make([]ValveID, 0, 5),
	}

	parts := bytes.Split(m[0][3], []byte(", "))
	for _, x := range parts {
		v.Neighbours = append(v.Neighbours, ID(string(x)))
	}

	return v, nil
}
