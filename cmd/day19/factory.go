package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// material is an enumeration for each type of material
type material int

const (
	OreMat material = iota
	ClayMat
	ObsMat
	_
	GeodeMat
)

// String implements fmt.Stringer()
func (m material) String() string {
	return [...]string{
		"ore",
		"clay",
		"obsidian",
		"",
		"geode",
	}[m]
}

// robot is a type of robot. Each robot gathers a different material.
// The robot's integer value is also the index within the blueprint
// where we store its material cost.
type robot int

const (
	OreBot robot = iota
	ClayBot
	ObsBot
	_
	GeodeBot
)

// Gather returns the material type that this robot gathers.
func (r robot) Gather() material {
	return material(r)
}

// AllBots lists all of the robots.
func AllBots() []robot {
	return []robot{OreBot, ClayBot, ObsBot, GeodeBot}
}

// String implements fmt.Stringer()
func (r robot) String() string {
	return [...]string{
		"ore bot",
		"clay bot",
		"obsisidan bot",
		"",
		"Geode Kracker 9000",
	}[r]
}

// blueprint defines the material costs for each type of robot.
type blueprint [6]byte

func (bp blueprint) Cost(r robot) map[material]int {
	m := map[material]int{
		OreMat: int(bp[r]),
	}
	if r == ObsBot {
		m[ClayMat] = int(bp[r+1])
	}
	if r == GeodeBot {
		m[ObsMat] = int(bp[r+1])
	}
	return m
}

var _re = regexp.MustCompile(`([0-9]+)`)

func ParseBlueprint(s string) (blueprint, error) {
	m := _re.FindAllString(s, -1)
	if len(m) != 7 {
		return blueprint{}, fmt.Errorf("parse %q: want 7 parts; got %d", s, len(m))
	}

	var bp blueprint

	for i := 0; i < 6; i++ {
		n, err := strconv.ParseUint(m[i+1], 10, 8)
		if err != nil {
			return blueprint{}, fmt.Errorf("parse %q: %w", m[i], err)
		}
		bp[i] = byte(n)
	}

	return bp, nil
}
