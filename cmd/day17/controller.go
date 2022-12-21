package main

import (
	"context"
	"fmt"
)

// Controller is responsible for manipulating the board an notifying
// observers about changes.
type Controller struct {
	model       *Board
	numRocks    int // rocks that we have actually cemented in position
	extraRocks  int // rocks to add on, based on skipped rows
	extraHeight int // height to add on, based on skipped rows
	seq         uint64
	parity      uint8        // even means fall, odd means shift
	movegen     func() byte  // a generator function that produces movements
	shapegen    func() Shape // a generator function that produces shapes
}

const _movement = ">>><<<>>><><>>><<>>>><<>>>><<<<><<<<>>>><>>><<<>><<<>>><<<<>><<<<>><<<<>>><><<<<>><><<<<>>>><<<>><><<<><><<>><<<<>>>><<><<<<>>>><<<<>>><>><<>>><<>><<><>>>><<<>>><<<>>><<<>>>><>>><>><<<<>><<<<>><>>><><>><<<>>>><<<>>>><<>>>><<<>><<<>>><><<<>>><<<<>>><<<<>>><<<>><<><<<>><><<<>>>><>><<>><<><<<<><<<<>>><<<><<>>><>>><<><<><<>>><<<>><<<><>><<<>><<<><<<>>>><<><>><<<>>><<<>>><>>><<<<>>><<<<>>><>><>>><<>><<<>>>><<>>><<>>>><>>>><><<<>>><>><<>><<<<><<>>><<<><<>>><<<>>><<<<>>><<<<>>>><<<<>><<>><<<><<<<>><<<><>>><<<>>><<<><<<<><<<<>>><<>>>><<><<<><>>><<<<>>>><<<>><<<>><<<<><>><>><<<>><<<<><<><<<><<>>>><<<><<<>>>><<<<>>>><<>>><<><<>>><<>><<<>><>><<>>><<<<><<<>>><<<>>>><<<<><<<>>><<<>><<<>>>><<>>><<<>>>><<<<>>>><>>>><<>>><>><<>>><<<<><><<<<>>>><<<>><<<>>><<<<>>>><<<<>>><<>>>><>>>><<<<>>><<<>>>><<<<><<<<>>>><<<<>><<<><<<><<<>>><<<<>>>><<>><<>>>><<>>><<<<><>>>><<>>><<>><>>><<<<>>>><<<<>>><<<><<<>>><<><<<><<<><>>>><<<<>>>><<<<>><<<>>><><>>><<<<>><<<>>>><<<>>>><<>>>><<<<>><<><<<>><<<<><<<>>><<<<>><>>><<<>>>><<<<>>>><<<><><<<<>><<<>><<>>><<<<><>>><<<<>>>><<>>><<>><<><<>>><<>>>><<>>><>><>>>><>>>><>>><>>>><<>><<<<><>>><>>><>><<>>>><>><<>>>><<>>>><<<>><>><>><>>><><>>>><>>>><>><<><<<>>><<>><<<<>>><<<<>>><<<>>><<<>>><<<<>>>><>><<<<>><<<<>>><<<>>>><<<><>>>><<<<>>>><>><<>>><<<<>>><<<>><<>>>><<>><<<<><<<>><>><<<<>><<<<>>><<<<><<<>>><<>>><<><><<>>><<<>><<<>>><<<<>>><<>>><><<>>><<<<><<>>><<<<>>><<>><<<<>><<<>>>><<<>>>><<<<>><<>><<<>><<<>>><<<>>><<<>>>><<>><<>>><>><<<<>>><<<<><<>>><<<<>><<<<>>><<<<>>><<<<>>>><<<<><>>><>><<<>>><>>><<<<><<>><<>><<<<>><<><<<>><<<>>>><<<>>>><<<>>><<<<>><>><<<<>>><<>><<<<>>><<<>>><<>><<<>>><<>>>><<<>>>><<<>>>><><><<<<><<<<>><<>><>>>><<>><>>><>>>><<<<>>>><><<<>>>><>><>>>><<><<>><<><>>><<<<>>><>>>><<<><<>>>><<><<<>>>><<<>><<<><<>>>><<<><<<>>>><<>><>>>><<>>><<<<><<<>>><<>>><<<<><><<<<><>><><><<<<><><<>>>><<<>><<<<>><>><<<><<<<>>>><<>><<>>>><>>><<<><<<>><<<<>>><>><><><<>>>><><<<>><<>>><<<>>><<<<>>><<>>>><<<>><<<>>><<>><<>>>><<<<>>><<>>>><><<<>><>><<<>><<<><<><>>>><><<<>>>><>>><<>><<><<<<>>>><<>>><<<<>>>><>><><<>>><<<><>><<><>><>>>><>>>><<<><><<<>>><<>>>><<<>>><<<<>>>><<<>><<<<>>>><<<>>><<<<><<<><<<>>><><<><>><<<<>>><<<<>>><<<>>>><<<<>><<<>>><><<>><<><<>>>><<>><>>><<<<>><<<>>>><<<>>><<<><<><<><<<>>><<<<>>><>>>><<<>>>><>>>><<<>>>><<>>>><<>>><<<>><<<>>>><>><<>>><<<>>>><<>>><<<>>>><<><<>><<>>>><<<>><<>>>><<<>><>><>><<<<>>>><<<>>>><>>><>>>><<<>><<>><><<<<>>><<<<>>><<<<><<<>>><<<<><<>>>><<>><<>>><<<>>><<<<><<<>>><<><<<<><<>><><>><<>>><>>>><<<><>>>><<>><>><><<<<><<<<>>>><<<><<<><<<<>><<>>><<<<><<<>>><<<<>><<<><><<<<><<>>><<>>>><<>>>><<<<>><<<<><<>>><<>>><<<<>><<<<><<<>><<<>><>><>>><<><<<<><<>><>><<<><<>><<>>>><<<<>>><<<<>><<<><<<><<>>><<>>>><<>><>>>><<<>>>><<<><<<<><<<<>>><>>>><<<<>>>><<<>>><>>><<<>><<<<><>>>><>>><><<<>>>><>>><>><<><<>><<<<>><<<>><<<>><<<<>>>><<<<>>>><<>>><<<<>>>><<<<>><<>>><<>>><<<<><<>>><<<>>><>>>><<><<<<>><>><>>>><>>>><><<><<>>><<><<>>><>>>><<>>><<<<>><<<><<>><>>><>><<>>><<><<>>>><<<<>>><<<>>>><>>>><>>><<>>>><<<<>><<><<<<>>>><<<><<>>>><>>><<<><<>>><>>>><<><<>><<>><<<>>><<><><<<<>><<<<><>>>><<<>>>><>><<<<>><<<>>><<<>>>><<<<>>>><><<<>>><<>><><>>>><<<<>>>><<<<><>>><<<>>>><<><<<>><<<<><>>>><<>><<<>><<<>>><<><<<>>><<<><<<>><<><>><><<<<>>><<<<>>><<><<<<><>>>><<<<>>><<>>>><>>>><<<><>><>><<<<>><<>><<><<>>><<><<<<>><>><<<>>><<<>><<<><<<<>>><<<>>><<>>><>><<<>><<<>><<<>><<><<><><<>>>><<<<>><>>>><<<>>>><<<>>>><><<<<>>><<<<>><<>><<<>>>><>><<<<>>><<<>><<><>>><<<<><<<>>>><<<>>><<<<>>>><><<>>>><>><<<<><<<><><>><<<>>>><<<>><>>><<<>>>><<<<>>>><<<<><><>><<<<>><>><>><<<<>>><<<<>>><<<<>>>><<<><>>><>>><>><>>>><<<<>>>><<>><<>>><<<<><<<<>><<<<>><<<>><<<<>><<><<<>>>><>>><<><<<>><<<>><<><>>>><<<<>>><><><<<<>><>>><<><<>><<<>>>><>><<<><<<<>>><<>>><<<<><<<<>><<><<>><><<<<>>>><<<><<>>><<<<>>><>><<<<>>>><<>>>><<><<>>>><<<<>><<<>>>><<>>><<<><<>>>><><<>><>><<<><<>>>><<>>>><<><<><<<<>>><<<<>><<<<><>>>><<<><<><<><<<>><<<<><<<><<<<>>><<<<>>><<>>><>>>><<<>><<>><><<>>><>><<<><>><<<>>><<<<>>>><><<<<>>><<<>><>><<<<>>><<<><<<<>>><><<>>>><>>>><<<<>><<<>>><<<>>><<<<><<>><<>><<<<>>>><<>>>><<><<>>><<>>><<<>>><><>>>><>>><<>>><<<<>>>><>>>><>>><<<<>>><<<<>>><><<<>>><<<<>>>><>><><<>>><<<><>>><<<>>><<<>>><<<>>>><>>>><<<<>>>><<<<>>>><<<>>>><<><<>><<<<>>><<<>>>><<<>><<<<>>><<><<><<<<>>><>><<<><<<>>>><<>><<><<<>>>><<<><<<>>><<<<>>>><<>>>><<<<>>><><<><<<<><<<<>>><<>>><>>>><><><>>>><<<<>>><<<>>><<<><<<<>>>><>><><<<>><>><<<<>>><><<<<>>>><<<>>>><>><<<>>>><>><<<><<<<><><>>><>>><<<>>><<><<<>>>><>>>><<<>>><>><><<<<><<><<<>>>><<><<><<><<<>>>><<<><<<>>><<<<>><<<<>>>><>>>><><<<>>><<<<>>>><<><>>>><<<>><>>>><<><<<<>><<<>><<>>>><<><>><><<<<><<<<>>>><>>><<>><<><>>>><<<><<<>><<><>>>><<>>>><<<>>><<<>>>><<>>>><<<>><<<>>><<<>><<<<>>><<>>>><<<>>>><<<>><<<<>>><>>>><>>><<<<>>>><<<<>><>>><<>>>><<<><><><<>>><<<<>>><>>><<<<>>>><<>>>><<<<>>>><<<>>><<<<>><<>>>><>>><>>><<>>><<<<>>>><<<><><<<<>>>><<<>>>><<<<>>><<<<>>><<<>><<<>>>><<>>>><><<<<>>>><<<<>><>><<<><<<><<<<>>><<<<>>>><>>><>><>>>><<<<>>><<>><<<>>><<>><<>><<><<>>><<>>><<<>>>><<<><><<<><<<<>>><<<>><<<<>>>><<>><<<><>><<<><>>>><<<<>>><<<<>><<><>><<<>>><<>>><<<>>><<<<>>>><<<>><<<<>><<>><>>>><>>><>><<<>>><<<><<<<>><<>>><<>>>><<>>><><<>><<<>>>><<><<>>>><<<>>>><<<<>>>><<<<>><<<<>>><>><<<>><>>>><<<><<>><<><<<<><<>>><>>><>>>><<><<<<>><>>><<<<><<<<>><<><>>>><<>><<<>>><<<>>>><<<><><<<>><<<>>>><<<>>><<<<>><<<<>><>><>><<<<>>><<><<<<>><>>><<<>><<><<>><<>>>><>><<<<><<<<>>><>>><<<>><<<>><>><<<<><<<>>>><>><>>><<<>>>><<<<>>><<>>>><<<>><<>>>><<<<>>><<<<>>><<<<>>>><<<<>><<<<>>><>>><>>><<<<>>><<<<>>><<>>>><<<<>>><<<>>><<<>><><<<>>>><<>>><<><<<>>>><<<>><>><>>>><><<><<<>>><<>>>><<<<>><<<<>>>><>>><<><<<<><<<<>>><>>>><<<>><><>>>><<<>><>><<<>><<><<>><<<>>><>>><<>>><<><<<>>><<<>>>><>><>><<<<><<<<>>>><<<>>>><<<>>><<>><<>>><>><<<<>><<>>><<<<><>><<>>><<<>>><<<<>>><>>><<>>>><<<<><>>>><<<>><<<>>>><<<>>><<<<>>>><<>><<<<>>><<><>><>>>><>>>><<<>><<>>><<<>>>><<<>>>><<><<><<<>><>>>><<>>><<<<>>>><><<>>>><<>>>><<<>><<<>>><<<>>>><<<>><<>>>><><><<<><<>><>>>><<>>>><<>><<<><<<><><<<<>>><><<<><<><<<>><<<><<<>>>><<<>><<>>><<<<>>>><>>><>>>><<<<>>><<<>><<<<>>>><><>>><><>><<>><<>>>><><>>><<<><>>><<<>>>><<><<<<>>><>><<>>>><>><<<>>><<<<>>>><>><<>>><<<>>><<<<>><<>>><<<<>><<>>><<<<>>>><<>>><<<<>><><<<<>><<<>>><<<><<<<>>><<<<>>><<><>>>><<>>><<<>><>>>><<<<><><<<><<>>>><<><<><>>>><<<>>>><<<<>>><<><<>><<<>>>><<>>>><>>>><<><<<>><<><<>><<>><<<>>><>>>><<<<><<<>><<<<><<><<<<><<<<>><<<<>>>><<><<>>><>>><<<>>><>><<<>><<<>>><<<<><<<>><<<<>>>><<>>>><<<<>><<<<>>><<<>>><<<<>>>><<><<>><<><<<>>><<<<>>>><<><<>>>><>><<<<><<>>>><><<<<>>><<<>><<<<>>><<>>>><<<<>>><<<<><<<<>>><>>><>>><<<>><<<><>>>><><<<><>><<<<><<<><<<>>>><<<>><>><<<><>>>><<>>><>><<>>>><<>>>><<<>><<<<>>><<<>>>><<<<><<<<>><<<<>>><><<<<>>><<><>><<<<>><<>><>>><<<<>><<><<>>>><<<>>><<<<><<<<>>><<<>><<<><<<>><<<<><<<<><>>>><>>><<<>>>><>>><<>><<<<>><<<>>><<<<>><<<<>>><<<<><<<>>>><<>>><<>>><>>>><>><<>>>><<>><>>><<><<<<>>>><<<>>><<<>><<<><<<>>><>>>><<<><<<<><<<><<>>><<<>><<<<>><>><<>><<<<><><<<<>>><>>><<<<>>>><<>><<<<><<<<>>><<>>><<>><<<<>><><<<><<>>>><><<<><<<<>>>><>>><>><<<<>>><>>>><<<><<<<>>><<<<><<<><<<>>><<<<>>><<<>>><>>>><<<>><<<<>>><<>><<>>><>>><>>><<<><<><<<<><<<>><<<>><<<<>><>>>><<<<>>><<>>><>>>><<<<>><<<<><>>>><<<<>><>>><<<<>>>><<>>>><<<<>><<<<>><<<>>>><<<>>>><<>>><<<<>>>><>>>><<<>>><><<<><<<<>><<>><<<>><>>><>><<><<<<>>>><<<<>>>><>><<>><<<>>><>>>><>>>><><<<<><<<>>>><<<<><<<>>><>>><<>>><<>>><<<<>>>><>>>><>>><<<<>>>><<<<>><<<<><<<>><<<<><<>><<>>><>><<>><>>>><<><<<<>><><>>><<<<>><<<><<>>>><>>>><<<<><><<<<><<>><<<>><<<<>>><<<><<<>><<<>>><><>>><><>>><<<<><<>><<<>>>><<<<>>><<<<>>>><>>>><>><<<<>><<<>><<>>>><<>><<>>><>>>><><><<>>>><<<<>>><>>>><<<>>>><<<<>><>>><<<<>>>><>><<<>><<><<<<>>><<<<>>><<<>>>><<<<>><<>>><<>>>><<<<>>>><<<<>>><<>><<<<>><<<>>>><<<>>>><<<>><<<<>>><<>>><>>>><><<>>>><<<<>>><<>>><<<><<<<>>>><>>><<<<>><<<<>><<<>><>>><>><<<<>>><<><>>><<>><<>>><<><<>><><><>>><<<>>>><<<<>>>><>>><<<>>>><<<><<>>>><>><<<<><<<<>>>><<<>>>><<<<>>><<<><<>>><><<<><<<>><<<>>>><<>>><<<<>>>><<<<>><<>><<<<>>>><<>><<<>>><<<><<<>><>><<<>>><<><<<><<><<<><<<<>>><<<><<<<>><>>><<<>>><>><<>>><<<><<>><<<<>>>><<<<>><<<>>>><<><>><<>>>><<<><<>><<<<>>><<>>><<<<>>>><<<>>><<><<>>><<>>>><<<>><>>><<<<>>><<>>><><>><>>><<<>><<<><<>><<>><<<<>>>><<<<><>><<<>>><<>>><><<<<>><<<<>><<>>><<><<<<><<><>>><<<><>>>><>><<<>><>><<<<>>>><<<>>>><<<<>>><<<<>>>><<>>><<>><<<<><>>>><<<>>>><<>><>><<><<<<>>>><<<>>>><<>>>><>><<<>><<><<<>>><<<>><<><<><<>>><<<>>><<<>><<<>><<><<<<>>>><<<>>><<><<<<>><<<>>><<>>><<>>><>><<<>>>><<<<>>><<<<>><<<>><>>>><<<>>><<<<>>><<><<<><<>><>><<>>><<>>><<><<>>><<<<>><<>>>><<>>><<<>>>><<<>><<><<>><<<<>><<>>>><<<<>><<>>><>>><<<>><<<><<>>><>><<<<>>>><<<<>><<<<><<><<<>>><<<<>>>><<<>>>><<><<<>><<<<>>>><<<<><<><<<>>>><<>>><<<>>>><<<<><>>>><>>>><<<<>>><>>>><<<><>>><<>>><>>><<<><<<<>>><<>>><>>>><<>><>>>><<>>><<<<>>>><<<<>>>><<<<>>><<>><<<<>>>><<<<>>><<<>><><<<<>>><<>><>>><>>>><<<<>><>>>><<<<>>><<<<>>><<<<><>>><<<<>><<<>>><<<>><<<<>>>><<<><<<<>>><<>><<<<>>><<<><<<<>>><<<>><>>>><<<<>>>><>>>><<<>>>><>>><<<>>><<<>><<>>>><<>><<<<>><>><<<><<<>>><<<<>>><>><<<>><>><<<>>>><<<<>>>><<<<><>>>><<<<><<<<><<<<>>>><<<<>><<>>>><>><<<<>><<>>><<<<><<<<>><<<>><<>>><<<>>>><<<<>>><<>><<<<>><<<<>><<>><<<<>>>><<><>><>><<><<<>>>><<>>><<<<>><<<><>>>><<<<><<>>>><<<<>>>><<>>><<<><<<>><>>><>>>><>>><<<>>>><>><>><>><<>><>><>>><<><>>>><<<>>><>><>>>><<<<>>><<<>><<>>>><>>>><<<>><>><<<<>><<><<<>>><<>>><<>>><<<<>><>>><>>><<<>><<>>>><<>>>><<><<<<>>>><<<<>>><>>><<<<><<<>>>><>><<<<>>>><<<<><<>><<<>>><<<<>>><><<><<>>><<<><<<>><<<<><<<<>><<>>>><><<<>>>><<<<>>><<<>><>>>><<>>><<<<>>>><><<><<<><<<>>><<<<>><<<>>><<<<><<<<>>>><><>><<<<><<><<>><<>><><<>><<<<>><<<<>>>><<<>>>><<<<>>><>>><<<>>>><<<<>>><<<><<<<>><<<>><<<>><<<>><>>>><<<>><<>>><>>>><>>>><<<><<<<>>>><<>>>><<<>>><<<><<<>>>><>>>><<><>>><><><<>>><<>>><>><<<><>>><<>>><<><<<><<<>><<>><>>>><>><<<><<<><<<>><>><>>><<><<>><<<<>>><<<<>>>><<<<><<>>>><<>>>><<>><<<>>>><>>>><<><<<>>>><<<<>>>><<<<>>><<>>><<><<>><<<<>>>><>>><<<>><<<>>><>>>><<<>>><<<>><<<<><<<>><<<<><<<>>>><<>>>><<<<><>>><>>>><>>><<<>>><<<>>><<>>><<<<>>><<<>><<>><<>><<<<>>>><<<>><<<<><<<>>>><>>><<><<><<<<><<<<>>>><<<><><<<>><><<<>>><<<<>>><<<<>><<<><<<<><<<<>><<<<>><<<>>><<<<>><<<><>>>><<<<>>>><<<><<<>"

// NewController initializes a new Controller for a game of 'nearly Tetris'.
func NewController(movegen func() byte) *Controller {
	if movegen == nil {
		movegen = generator(0, []byte(_movement))
	}

	return &Controller{
		model: &Board{
			rows: make([]Row, 0, 64),
		},
		shapegen: generator(0, _shapes),
		movegen:  movegen,
	}
}

// Run begins the simulation, and allows it to continue until the given context
// is cancelled, or the given number of rocks have stopped falling.
func (c *Controller) Run(ctx context.Context, ticker <-chan struct{}, maxRocks ...int) <-chan GameEvent {
	ch := make(chan GameEvent)
	c.seq = 0

	go func() {
		defer close(ch)

		ch <- GameEvent{
			Seq:  c.seq,
			Type: GameStartedEvent,
			Msg:  GameStartedEvent.String(),
		}

		for {
			c.seq++

			select {
			case <-ctx.Done():
				err := ctx.Err()
				ch <- GameEvent{
					Seq:         c.seq,
					Type:        GameStoppedEvent,
					TotalRocks:  c.numRocks + c.extraRocks,
					TotalHeight: c.model.height + c.extraHeight,
					Msg:         err.Error(),
					Error:       err,
				}
				return

			case <-ticker:
				ev := c.tick()
				ch <- ev

				if ev.Type != RockStoppedEvent || len(maxRocks) == 0 {
					continue
				}

				if skip := c.extraRocks == 0 && c.numRocks > 2300 &&
					c.numRocks%_p2_period == 0; skip {
					x := (maxRocks[0]-c.numRocks)/_p2_period - 1
					fmt.Println("skipping", x, "periods")
					c.extraRocks = x * _p2_period
					c.extraHeight = x * _p2_height
				}

				if c.numRocks+c.extraRocks >= maxRocks[0] {
					ch <- GameEvent{
						Seq:         c.seq,
						Type:        GameStoppedEvent,
						TotalRocks:  c.numRocks + c.extraRocks,
						TotalHeight: c.model.height + c.extraHeight,
						Msg:         fmt.Sprintf("stopped after %d real rocks and %d fake ones", c.numRocks, c.extraRocks),
					}
					return
				}
			}
		}
	}()

	return ch
}

// tick processes one cycle of game activity.
func (c *Controller) tick() GameEvent {
	if c.model.curr == nil {
		return c.newRock()
	}

	var ev GameEvent
	if c.parity%2 == 0 {
		ev = c.fall()
	} else {
		ev = c.shift()
	}

	c.parity++
	return ev
}

// newRock adds the next shape to the board.
//
// Each rock appears so that its left edge is two units away from the
// left wall and its bottom edge is three units above the highest rock
// in the room (or the floor, if there isn't one).
func (c *Controller) newRock() GameEvent {
	c.parity = 1 // after appearing, the next thing a rock should do is shift.
	m := c.model
	for len(m.rows) < m.height+7 {
		m.rows = append(m.rows, 0)
	}
	m.curr = &Rock{c.shapegen(), m.height + 3}

	rows, from := c.compose(m.height)
	return GameEvent{
		Seq:      c.seq,
		Type:     NewRockEvent,
		Msg:      NewRockEvent.String(),
		Rows:     rows,
		RowsFrom: from,
	}
}

func (c *Controller) fall() GameEvent {
	m, rock := c.model, c.model.curr

	start, size := c.rockIndices()
	next := start - 1
	if next < 0 {
		return c.cement("rock stopped on the floor")
	}

	for i, j := next, 0; j < size; i, j = i+1, j+1 {
		if isConflict(m.rows[i], rock.rows[j]) {
			return c.cement("rock stopped by another rock")
		}
	}

	m.curr.start--
	rows, from := c.compose(next)
	return GameEvent{
		Seq:      c.seq,
		Type:     RockMovedEvent,
		Msg:      "rock fell",
		Rows:     rows,
		RowsFrom: from,
	}
}

// cement the rock in position
func (c *Controller) cement(msg string) GameEvent {
	m, rock := c.model, c.model.curr

	start, size := c.rockIndices()
	for i, j := start, 0; j < size; i, j = i+1, j+1 {
		row := m.rows[i] | rock.rows[j]
		m.rows[i] = row
		if row > 0 && m.height < i+1 {
			m.height = i + 1
		}
	}
	c.model.curr = nil

	rows := make([]Row, len(m.rows)-start)
	copy(rows, m.rows[start:])

	c.numRocks++
	return GameEvent{
		Seq:         c.seq,
		Type:        RockStoppedEvent,
		Msg:         msg,
		TotalRocks:  c.numRocks + c.extraRocks,
		TotalHeight: c.model.height + c.extraHeight,
		Rows:        rows,
		RowsFrom:    start,
	}
}

// try to shift the rock either left or right
func (c *Controller) shift() GameEvent {
	m, rock := c.model, c.model.curr
	var msg string

	start, _ := c.rockIndices()

switch1:
	switch dir := c.movegen(); dir {
	case '<':
		msg = "rock pushed left"
		for i, row := range rock.rows {
			if row >= 0x40 {
				msg += "; but blocked by the wall"
				break switch1
			}
			if isConflict(row<<1, m.rows[start+i]) {
				msg += "; but blocked by another rock"
				break switch1
			}
		}
		for i, row := range rock.rows {
			(*rock).rows[i] = row << 1
		}

	case '>':
		msg = "rock pushed right"
		for i, row := range rock.rows {
			if row%2 == 1 {
				msg += "; but blocked by the wall"
				break switch1
			}
			if isConflict(row>>1, m.rows[start+i]) {
				msg += "; but blocked by another rock"
				break switch1
			}
		}
		for i, row := range rock.rows {
			(*rock).rows[i] = row >> 1
		}
	}

	rows, from := c.compose(start)
	return GameEvent{
		Seq:      c.seq,
		Type:     RockMovedEvent,
		Msg:      msg,
		Rows:     rows,
		RowsFrom: from,
	}
}

// compose returns a copy of top rows of the board.
// assumes that the rock is contained within this section, if it exists.
func (c *Controller) compose(from int) (rows []Row, start int) {
	m := c.model
	if m.curr == nil {
		return m.rows, 0
	}

	rows = make([]Row, len(m.rows)-from)
	copy(rows, m.rows[from:])

	i, size := c.rockIndices()
	i -= from

	for j := 0; j < size; i, j = i+1, j+1 {
		rows[i] |= m.curr.rows[j]
	}

	return rows, from
}

func (c *Controller) rockIndices() (lower, size int) {
	rock := c.model.curr
	if rock == nil {
		return 0, 0
	}

	lower = rock.start
	size = len(rock.rows)
	return
}

func isConflict(a, b Row) bool {
	return a&b > 0
}

func generator[T any](i int, source []T) func() T {
	return func() (val T) {
		if i >= len(source) {
			i -= len(source)
		}
		val, i = source[i], i+1
		return val
	}
}
