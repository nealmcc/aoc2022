package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

// Valve is one vertex in the graph. Each Valve has a unique Key, as well as
// a sequential index, based on its alphabetical sequence compared to all other
// valves.
type Valve struct {
	ID         ValveID
	Flow       int
	Neighbours []ValveID
	ix         int
}

// ValveID is a numeric equivalent of the two letter string used to identify a valve.
type ValveID uint16

// ReadValves reads the input and parses the valves.
func ReadValves(r io.Reader) (map[ValveID]*Valve, error) {
	s := bufio.NewScanner(r)
	valves := make(map[ValveID]*Valve, 64)

	for s.Scan() {
		v, err := ParseValve(s.Text())
		if err != nil {
			return nil, err
		}
		valves[v.ID] = &v
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return valves, nil
}

var _re = regexp.MustCompile(`Valve ([A-Z][A-Z]) has flow rate=([\d]+); tunnels? leads? to valves? ((?:[A-Z][A-Z])(?:(?:, )?(?:[A-Z][A-Z]))*)`)

func ParseValve(s string) (Valve, error) {
	m := _re.FindAllStringSubmatch(s, -1)
	if len(m) != 1 {
		return Valve{}, fmt.Errorf("parse(%q): want 1 match got %d", s, len(m))
	}

	flow, err := strconv.Atoi(m[0][2])
	if err != nil {
		return Valve{}, fmt.Errorf("parse flow %q: %w", m[0][2], err)
	}

	v := Valve{
		ID:         K(m[0][1]),
		Flow:       flow,
		Neighbours: make([]ValveID, 0, 4),
	}

	parts := strings.Split(m[0][3], ", ")
	for _, x := range parts {
		v.Neighbours = append(v.Neighbours, K(string(x)))
	}

	return v, nil
}

// K converts the given two-character string to a valve Key.
// Input is assumed to be valid (not checked).
// Copied from binary.BigEndian.Uint16
func K(s string) ValveID {
	return ValveID(ValveID(s[0])<<8 | ValveID(s[1]))
}

// Format implements fmt.Formatter.
// The formatting verb is ignored.
func (k ValveID) Format(s fmt.State, verb rune) {
	first := byte((k & 0xFF00) >> 8)
	second := byte(k & 0x00FF)
	s.Write([]byte{first, second})
}

// String implements fmt.Stringer.
func (k ValveID) String() string {
	return fmt.Sprintf("%v", k)
}
