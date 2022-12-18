package main

import (
	"strings"

	v "github.com/nealmcc/aoc2022/pkg/vector/twod"
)

type (
	row   byte
	shape [4]byte
	piece struct {
		pos   v.Point // row and x position of the bottom-left corner of the shape
		shape shape
	}

	board struct {
		top  int
		rows []row
		curr shape
	}
)

const (
	air  = 0
	rock = 1
)

// make Shape converts a string into a slice of bytes - does not check its input.
func makeShape(s string) shape {
	rows := strings.Split(s, "\n")
	var sh shape

	for y := 0; y < len(rows) && y < 4; y++ {
		row := rows[y]
		for x := 0; x < len(row) && x < 7; x++ {
			if row[x] == '#' || row[x] == 'x' {
				sh[y] |= rock << (6 - x)
			}
		}
	}

	return sh
}
