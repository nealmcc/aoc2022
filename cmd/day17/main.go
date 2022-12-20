package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strconv"
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

	// p2 := part2(sensors, 4000000)
	// end := time.Now()

	fmt.Printf("part 1: %d in %s\n", p1, middle.Sub(start))
	// fmt.Printf("part 2: %d in %s\n", p2, end.Sub(middle))
}

// part1 solves part 1 of the puzzle
func part1(moves string, w io.Writer) int {
	ctrl := NewController(generator(0, []byte(moves)))

	tick := make(chan struct{})
	defer close(tick)

	go func() {
		for {
			tick <- struct{}{}
		}
	}()

	ch := ctrl.Run(context.Background(), tick, 2022)

	buf := &Buffer{}
	lastRowRendered := 0

	numrocks := 0
	height := 0

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
			numrocks++
			fmt.Fprintf(w, "%+v\n", ev)
			r := buf.Reader()
			r.Seek(0, io.SeekStart)
			save(r, fmt.Sprintf("part1-%04d.log", ev.Seq))

		case GameStoppedEvent:
			if ev.Error == nil {
				height, _ = strconv.Atoi(ev.Msg)
			}
			r := buf.Reader()
			r.Seek(0, io.SeekStart)
			io.Copy(w, r)
			r.Seek(0, io.SeekStart)
			save(r, "part1-final.log")
		}
	}
	return height
}

// animate1 animates part 1 of the puzzle
func animate1(ctx context.Context, moves string, w io.Writer) int {
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

	height := 0

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

		case GameStoppedEvent:
			if ev.Error == nil {
				height, _ = strconv.Atoi(ev.Msg)
			}

		default:
			fmt.Fprintf(w, "%+v\n", ev)
		}
		r := buf.Reader()
		r.Seek(0, io.SeekStart)
		io.Copy(w, r)
	}
	return height
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
