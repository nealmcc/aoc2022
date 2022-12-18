package main

import (
	"io"
	"strings"
)

type (
	row uint8 // The 1's bit is the rightmost edge. The MSB is ignored.

	// To shift a shape left or right, just shift all 4 bytes left or right by one.
	// if any row of the shape is odd, then we can no longer shift right.
	// if any row of the shape is >= 0x40 (second MSB) then we can no longer shift left.
	shape [4]row

	board struct {
		top  int   // the tallest row that contains (part of) a piece
		rows []row // rows are numbered from 0 at the bottom
		// curr piece
	}

	piece struct {
		shape  shape
		height int // the index on the board for the top row of this shape
	}
)

// bit masks for an individual square on the board.
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

// WriteTo implements io.WriterTo
func (b board) WriteTo(w io.Writer) (n64 int64, err error) {
	var sum int64

	for i := b.top; i >= 0 && i < len(b.rows); i-- {
		n, err := b.rows[i].WriteTo(w)
		sum += n
		if err != nil {
			return sum, err
		}
	}

	n, err := w.Write([]byte("+-------+"))
	sum += int64(n)
	if err != nil {
		return sum, err
	}

	return sum, nil
}

// WriteTo implements io.WriterTo
func (r row) WriteTo(w io.Writer) (n64 int64, err error) {
	buf := make([]byte, 10)
	buf[0] = '|'
	buf[8] = '|'
	buf[9] = '\n'

	for i := 7; i > 0; i-- {
		if r%2 == 1 {
			buf[i] = '#'
		} else {
			buf[i] = '.'
		}
		r >>= 1
	}

	n, err := w.Write(buf)
	return int64(n), err
}
