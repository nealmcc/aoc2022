package main

import (
	"testing"

	v "github.com/nealmcc/aoc2022/pkg/vector/twod"
)

func TestSensor_Radius(t *testing.T) {
	t.Parallel()
	tt := []struct {
		name string
		in   Sensor
		want int
	}{
		{
			name: "empty sensor has zero radius",
			want: 0,
		},
		{
			name: "sensor with a beacon on the center point has zero radius",
			in: Sensor{
				Center: v.Point{X: 1, Y: 1},
				Beacon: v.Point{X: 1, Y: 1},
			},
			want: 0,
		},
		{
			name: "sensor with a beacon along its x axis",
			in: Sensor{
				Center: v.Point{X: 1, Y: 1},
				Beacon: v.Point{X: 4, Y: 1},
			},
			want: 3,
		},
		{
			name: "sensor with a beacon along its y axis",
			in: Sensor{
				Center: v.Point{X: 1, Y: 1},
				Beacon: v.Point{X: 1, Y: 12},
			},
			want: 11,
		},
		{
			name: "sensor with a beacon on neither axis",
			in: Sensor{
				Center: v.Point{X: -3, Y: 2},
				Beacon: v.Point{X: 1, Y: 12},
			},
			want: 14,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := tc.in.Radius()
			if got != tc.want {
				t.Logf("%v.Radius() = %d ; want %d", tc.in, got, tc.want)
				t.Fail()
			}
		})
	}
}

func TestSensor_SegmentAt(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name   string
		circle Sensor
		y      int
		want   Segment
	}{
		{
			name:   "radius 1 sensor, sliced through the middle => length of 3",
			circle: Sensor{Beacon: v.Point{Y: 1}},
			y:      0,
			want:   Segment{From: -1, To: 1},
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, _ := tc.circle.SegmentAt(tc.y)
			if got != tc.want {
				t.Logf("SegmentAt(%d) = %+v; want %+v", tc.y, got, tc.want)
				t.Fail()
			}
		})
	}
}
