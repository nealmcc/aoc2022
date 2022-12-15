package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	v "github.com/nealmcc/aoc2022/pkg/vector/twod"
)

const _sample = `Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3
`

func TestParseRow(t *testing.T) {
	t.Parallel()

	tt := []struct {
		in   string
		want Sensor
	}{
		{
			`Sensor at x=2, y=18: closest beacon is at x=-2, y=15`,
			Sensor{
				Center: v.Point{X: 2, Y: 18},
				Beacon: v.Point{X: -2, Y: 15},
			},
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.in, func(t *testing.T) {
			t.Parallel()

			got, err := parseRow([]byte(tc.in))
			if err != nil {
				t.Log("unexpected error", err)
				t.FailNow()
			}
			if got != tc.want {
				t.Logf("parseRow(%q) = %v; want %v", tc.in, got, tc.want)
				t.Fail()
			}
		})
	}
}

func TestParse(t *testing.T) {
	t.Parallel()

	sensors, err := read(strings.NewReader(_sample))
	if err != nil {
		t.Log("error reading sample:", err)
		t.FailNow()
	}

	got, want := len(sensors), 14
	if got != want {
		t.Logf("got %d sensors; want %d", got, want)
		t.Fail()
	}

	assert.Equal(t, sensors[3], Sensor{
		Center: v.Point{X: 12, Y: 14},
		Beacon: v.Point{X: 10, Y: 16},
	})
}

func TestPart1(t *testing.T) {
	t.Parallel()

	sensors, err := read(strings.NewReader(_sample))
	if err != nil {
		t.Log("error reading sample:", err)
		t.FailNow()
	}

	got, want := part1(sensors, 10), 26
	if got != want {
		t.Logf("part1() = %d; want %d", got, want)
		t.Fail()
	}
}

func TestPart2(t *testing.T) {
	t.Parallel()

	sensors, err := read(strings.NewReader(_sample))
	if err != nil {
		t.Log("error reading sample:", err)
		t.FailNow()
	}

	got, want := part2(sensors, 20), 56000011
	if got != want {
		t.Logf("part2() = %d; want %d", got, want)
		t.Fail()
	}
}
