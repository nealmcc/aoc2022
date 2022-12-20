package main

import (
	"io"
	"strings"
)

type (
	// Board is the data model for a game of 'almost Tetris'
	Board struct {
		height int   // the height of the rocks that have stopped.
		rows   []Row // rows are numbered from 0 at the bottom
		curr   *Rock // the piece that is currently falling
	}

	Rock struct {
		rows  Shape // the shape of the current piece, as modified by left and right movement.
		start int   // the index on the board where the bottom of this rock is
	}

	Row uint8 // The 1's bit is the rightmost edge. The MSB is ignored.

	// To shift a Shape left or right, just shift all 4 bytes left or right by one.
	// if any row of the Shape is odd, then we can no longer shift right.
	// if any row of the Shape is >= 0x40 (second MSB) then we can no longer shift left.
	Shape [4]Row
)

var (
	_dash   = Shape{30, 0, 0, 0}
	_plus   = Shape{8, 28, 8, 0}
	_corner = Shape{28, 4, 4, 0}
	_bar    = Shape{0x10, 0x10, 0x10, 0x10}
	_square = Shape{24, 24, 0, 0}
)

var _shapes = []Shape{_dash, _plus, _corner, _bar, _square}

// makeShape is a utility function that helps to define shapes like the above.
func makeShape(s string) Shape {
	rows := strings.Split(s, "\n")
	var sh Shape

	for y := 0; y < len(rows) && y < 4; y++ {
		row := rows[y]
		for x := 0; x < len(row) && x < 7; x++ {
			if row[x] == '#' || row[x] == 'x' {
				sh[y] |= 1 << (6 - x)
			}
		}
	}

	reverse(sh[:])
	return sh
}

func reverse[T any](items []T) {
	for i, j := 0, len(items)-1; i < j; i, j = i+1, j-1 {
		items[i], items[j] = items[j], items[i]
	}
}

// WriteTo implements io.WriterTo
func (r Row) WriteTo(w io.Writer) (n64 int64, err error) {
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

// WriteTo implements io.WriterTo
func (b Board) WriteTo(w io.Writer) (n64 int64, err error) {
	var sum int64

	for i := b.height; i >= 0 && i < len(b.rows); i-- {
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
