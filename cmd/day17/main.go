package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var _animateFlag = flag.Bool("animate", false, "run the animation")

func main() {
	flag.Parse()

	if *_animateFlag {
		ctx := withInterrupt(context.Background())
		animate1(ctx, _movement, os.Stdout)
		return
	}

	start := time.Now()
	p1 := part1(_movement, os.Stdout)
	middle := time.Now()
	p2 := part2(_movement, os.Stdout)
	end := time.Now()

	fmt.Printf("part 1: %d in %s\n", p1, middle.Sub(start))
	fmt.Printf("part 2: %d in %s\n", p2, end.Sub(middle))
}

// part1 solves part 1 of the puzzle
func part1(moves string, w io.Writer) int {
	ctrl := NewController(generator(0, []byte(moves)))

	tick := make(chan struct{})

	go func() {
		defer close(tick)
		for {
			tick <- struct{}{}
		}
	}()

	ch := ctrl.Run(context.Background(), tick, 2022)

	buf := &Buffer{}
	lastRowRendered := 0

	var height int
	for ev := range ch {
		switch ev.Type {

		case NewRockEvent, RockMovedEvent:
			if ev.RowsFrom != lastRowRendered {
				delta := lastRowRendered - ev.RowsFrom
				lastRowRendered -= delta
				buf.Seek(int64(-10*delta), io.SeekEnd)
			}
			for i := 0; i < len(ev.Rows); i++ {
				ev.Rows[i].WriteTo(buf)
			}
			lastRowRendered += len(ev.Rows)

		case GameStoppedEvent:
			height = ev.TotalHeight
			r := buf.Reader()
			r.Seek(0, io.SeekStart)
			save(r, fmt.Sprintf("part1-final-%04d.log", ev.TotalRocks))
		}
	}
	return height
}

const (
	_p2_period = 1730 // we get a repeating pattern every 1730 rocks (after about rock 300)
	_p2_height = 2647 // 2647 height for every 1730 rocks, once the pattern starts
)

// part2 solves part 2 of the puzzle
func part2(moves string, w io.Writer) int {
	ctrl := NewController(generator(0, []byte(moves)))

	tick := make(chan struct{})

	go func() {
		defer close(tick)
		for {
			tick <- struct{}{}
		}
	}()

	ch := ctrl.Run(context.Background(), tick, 1000000000000)
	buf := &Buffer{}
	lastRowRendered := 0

	var height int
	for ev := range ch {
		switch ev.Type {

		case NewRockEvent, RockMovedEvent:
			if ev.RowsFrom != lastRowRendered {
				delta := lastRowRendered - ev.RowsFrom
				lastRowRendered -= delta
				buf.Seek(int64(-10*delta), io.SeekEnd)
			}
			for i := 0; i < len(ev.Rows); i++ {
				ev.Rows[i].WriteTo(buf)
			}
			lastRowRendered += len(ev.Rows)

		case RockStoppedEvent:
			if ev.TotalRocks%_p2_period == 0 {
				r := buf.Reader()
				r.Seek(0, io.SeekStart)
				save(r, fmt.Sprintf("part2-%06d.log", ev.TotalRocks))
			}

		case GameStoppedEvent:
			height = ev.TotalHeight
			r := buf.Reader()
			r.Seek(0, io.SeekStart)
			save(r, fmt.Sprintf("part2-final-%04d.log", ev.TotalRocks))
		}
	}
	return height
}

// animate1 animates part 1 of the puzzle
func animate1(ctx context.Context, moves string, w io.Writer) {
	ctrl := NewController(generator(0, []byte(moves)))

	tick := make(chan struct{})
	defer close(tick)

	go func() {
		t := time.NewTicker(100 * time.Millisecond)
		defer t.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
				tick <- struct{}{}
			}
		}
	}()

	ch := ctrl.Run(ctx, tick, 2022)

	buf := &Buffer{}
	lastRowRendered := 0

	for ev := range ch {
		fmt.Fprintf(w, "%+v\n", ev)
		fmt.Fprintln(w, lastRowRendered)
		switch ev.Type {
		case GameStartedEvent:
			fmt.Fprintln(buf, "+-------+")

		case NewRockEvent, RockMovedEvent:
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

func save(r io.Reader, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	io.Copy(file, r)
}
