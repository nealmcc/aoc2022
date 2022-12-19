package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// Factory holds a potential state of the factory.
// Factory is comparable so it can be used as a map key.
type Factory struct {
	bp   Blueprint
	bots [4]int
	mats [4]int
}

// CanAfford checks to see if the factory contains enough raw materials to
// build the given robot.
func (f Factory) CanAfford(r Robot) bool {
	for i, cost := range f.bp.cost(r) {
		if f.mats[i] < cost {
			return false
		}
	}
	return true
}

// Spend pays the cost of the given robot. The robot isn't completed yet.
func (f *Factory) Spend(r Robot) {
	for i, cost := range f.bp.cost(r) {
		f.mats[i] -= cost
	}
}

// Tick completes the current turn for the factory, including sending all bots
// out to gather materials, returning the new number of cracked geodes,
// and (optionally) adding a robot to the factory afterwards.
func (f *Factory) Tick(r ...Robot) int {
	if len(r) > 0 {
		f.Spend(r[0])
	}
	for i, n := range f.bots {
		(*f).mats[i] += n
	}
	sc := f.mats[Geodemat]
	if len(r) > 0 {
		(*f).bots[r[0]] += 1
	}
	return sc
}

// Blueprint defines the material costs for each type of robot.
type Blueprint [6]byte

func (bp Blueprint) cost(r Robot) [4]int {
	m := [4]int{int(bp[r])}

	if r == Obsbot {
		m[Claymat] = int(bp[r+2])
	}
	if r == Geodebot {
		m[Obsmat] = int(bp[r+2])
	}
	return m
}

var _re = regexp.MustCompile(`([0-9]+)`)

func ParseBlueprint(s string) (Blueprint, error) {
	m := _re.FindAllString(s, -1)
	if len(m) != 7 {
		return Blueprint{}, fmt.Errorf("parse %q: want 7 parts; got %d", s, len(m))
	}

	var bp Blueprint

	save := func(i, j int) error {
		n, err := strconv.ParseUint(m[i], 10, 8)
		if err != nil {
			return fmt.Errorf("parse %q: %w", m[i], err)
		}
		bp[j] = byte(n)
		return nil
	}

	if err := save(1, 0); err != nil {
		return Blueprint{}, err
	}
	if err := save(2, 1); err != nil {
		return Blueprint{}, err
	}
	if err := save(3, 2); err != nil {
		return Blueprint{}, err
	}
	if err := save(4, 4); err != nil {
		return Blueprint{}, err
	}
	if err := save(5, 3); err != nil {
		return Blueprint{}, err
	}
	if err := save(6, 5); err != nil {
		return Blueprint{}, err
	}

	return bp, nil
}

// Material is an enumeration for each type of Material
type Material int

const (
	Oremat Material = iota
	Claymat
	Obsmat
	Geodemat
)

// String implements fmt.Stringer()
func (m Material) String() string {
	return [...]string{
		"ore",
		"clay",
		"obsidian",
		"geode",
	}[m]
}

// Robot is a type of Robot. Each Robot gathers a different material.
// The Robot's integer value is also the index within the blueprint
// where we store its material cost.
type Robot int

const (
	Orebot Robot = iota
	Claybot
	Obsbot
	Geodebot
)

// String implements fmt.Stringer()
func (r Robot) String() string {
	return [...]string{
		"ore bot",
		"clay bot",
		"obsisidan bot",
		"Geode Kracker 9000",
	}[r]
}
