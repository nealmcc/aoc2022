package main

import (
	"io"
	"strings"
)

type (
	// Board is the data model for a game of 'almost Tetris'
	Board struct {
		top  int   // the tallest row that contains solid rock.
		rows []Row // rows are numbered from 0 at the bottom
		curr *Rock // the piece that is currently falling
	}

	Rock struct {
		sh     Shape // the shape of the current piece, as modified by left and right movement.
		offset int   // number of rows below the last index of the board, where the top of this rock begins.
	}

	Row uint8 // The 1's bit is the rightmost edge. The MSB is ignored.

	// To shift a Shape left or right, just shift all 4 bytes left or right by one.
	// if any row of the Shape is odd, then we can no longer shift right.
	// if any row of the Shape is >= 0x40 (second MSB) then we can no longer shift left.
	Shape [4]Row
)

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

	return sh
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

// placeRock adds the given shape to this board and possibly
// adds more rows to the board, to accomodate it.
// returns the number of rows added.
func (b *Board) placeRock(sh Shape) int {
	n := 0
	for len(b.rows) < b.top+7 {
		b.rows = append(b.rows, 0)
		n++
	}
	b.curr = &Rock{sh, b.top + 7}
	return n
}
