package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"

	v "github.com/nealmcc/aoc2022/pkg/vector/twod"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	start := time.Now()

	sensors, err := read(file)
	if err != nil {
		log.Fatal(err.Error())
	}

	p1 := part1(sensors, 2000000)
	middle := time.Now()

	p2 := part2(sensors, 4000000)
	end := time.Now()

	fmt.Printf("part 1: %d in %s\n", p1, middle.Sub(start))
	fmt.Printf("part 2: %d in %s\n", p2, end.Sub(middle))
}

// read the lines from the given input.
func read(r io.Reader) ([]Sensor, error) {
	s := bufio.NewScanner(r)
	sensors := make([]Sensor, 0, 27)

	for s.Scan() {
		sensor, err := parseRow(s.Bytes())
		if err != nil {
			return nil, err
		}
		sensors = append(sensors, sensor)
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return sensors, nil
}

var _re = regexp.MustCompile(`(?:x=([-\d]+), y=([-\d]+))`)

func parseRow(b []byte) (Sensor, error) {
	halves := _re.FindAllSubmatch(b, -1)
	if len(halves) != 2 {
		return Sensor{}, fmt.Errorf("parse(%q): want 2 halves got %d",
			b, len(halves))
	}

	x1, err := strconv.Atoi(string(halves[0][1]))
	if err != nil {
		return Sensor{}, fmt.Errorf("parse(%q): %w", b, err)
	}
	y1, err := strconv.Atoi(string(halves[0][2]))
	if err != nil {
		return Sensor{}, fmt.Errorf("parse(%q): %w", b, err)
	}
	x2, err := strconv.Atoi(string(halves[1][1]))
	if err != nil {
		return Sensor{}, fmt.Errorf("parse(%q): %w", b, err)
	}
	y2, err := strconv.Atoi(string(halves[1][2]))
	if err != nil {
		return Sensor{}, fmt.Errorf("parse(%q): %w", b, err)
	}

	center := v.Point{X: x1, Y: y1}
	beacon := v.Point{X: x2, Y: y2}

	sensor := Sensor{
		Center: center,
		Beacon: beacon,
	}

	return sensor, nil
}

// part1 solves part 1 of the puzzle:
func part1(sensors []Sensor, y int) int {
	segments := segmentsAt(sensors, y)

	beacons := make(map[int]struct{}, 4)
	for _, s := range sensors {
		if s.Beacon.Y == y {
			beacons[s.Beacon.X] = struct{}{}
		}
	}

	sum := 0
	for _, seg := range segments {
		sum += seg.Length()
		for x := range beacons {
			if x >= seg.From && x <= seg.To {
				sum--
			}
		}
	}

	return sum
}

func part2(sensors []Sensor, limit int) int {
	beacons := make(map[v.Point]struct{}, 27)
	for _, s := range sensors {
		beacons[s.Beacon] = struct{}{}
	}

	var x, y int
	for y = limit; y >= 0; y-- {
		segments := Constrain(0, limit, segmentsAt(sensors, y))
		min, max := limit+1, -1
		for _, seg := range segments {
			if seg.From < min {
				min = seg.From
			}
			if seg.To > max {
				max = seg.To
			}
		}
		if min != 0 {
			x = 0
			break
		} else if max != limit {
			x = limit
			break
		} else if len(segments) == 2 {
			x = segments[0].To + 1
			break
		}
	}
	return y + x*4000000
}

func segmentsAt(sensors []Sensor, y int) []Segment {
	segments := make([]Segment, 0, len(sensors))
	for _, s := range sensors {
		if seg, ok := s.SegmentAt(y); ok {
			segments = append(segments, seg)
		}
	}
	return JoinSegments(segments...)
}
