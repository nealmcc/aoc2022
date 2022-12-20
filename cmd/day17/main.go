package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx := withInterrupt(context.Background())

	start := time.Now()
	p1 := part1(ctx, _movement, os.Stdout)
	middle := time.Now()

	// p2 := part2(sensors, 4000000)
	// end := time.Now()

	fmt.Printf("part 1: %d in %s\n", p1, middle.Sub(start))
	// fmt.Printf("part 2: %d in %s\n", p2, end.Sub(middle))
}

// part1 solves part 1 of the puzzle:
//
// Run the simulation 2022 times, and see how tall the tower will get.
func part1(ctx context.Context, moves string, w io.Writer) int {
	ctrl := NewController(generator(0, []byte(moves)), nil)

	ch := ctrl.Run(ctx, 2022)

	buf := &Buffer{}
	lastRowRendered := 0

	for ev := range ch {
		fmt.Fprintf(w, "%+v\n", ev)
		fmt.Fprintln(w, lastRowRendered)
		switch ev.Type {
		case GameStartedEvent:
			fmt.Fprintln(buf, "+-------+")
		case NewRockEvent, RockMovedEvent, RockStoppedEvent:
			if ev.RowsFrom != lastRowRendered {
				delta := lastRowRendered - ev.RowsFrom
				fmt.Fprintln(w, "delta", delta)
				lastRowRendered -= delta
				buf.Seek(int64(-10*delta), io.SeekEnd)
			}
			for i := 0; i < len(ev.Rows); i++ {
				ev.Rows[i].WriteTo(buf)
			}
			lastRowRendered += len(ev.Rows)
		}
		r := buf.Reader()
		r.Seek(0, io.SeekStart)
		io.Copy(w, r)
	}
	return 0
}

// withInterrupt wraps the given context, and will cancel it when the user
// presses CTRL+C or when the OS signals the program to end.
func withInterrupt(ctx context.Context) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		cancel()
	}()

	return ctx
}
