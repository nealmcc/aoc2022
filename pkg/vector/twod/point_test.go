package twod

import (
	"fmt"
	"testing"
)

func ExamplePoint_Reduce() {
	a := Point{X: -12, Y: -3}
	b, scale := a.Reduce()
	fmt.Println(b, scale)
	// Output: {-4 -1} 3
}

func TestManhattanLength(t *testing.T) {
	tt := []struct {
		name string
		in   Point
		want int
	}{
		{"zero vector has zero length", Point{}, 0},
		{"vector in quadrant 1", Point{3, 3}, 6},
		{"vector in quadrant 2", Point{-4, 4}, 8},
		{"vector in quadrant 3", Point{-5, -5}, 10},
		{"vector in quadrant 4", Point{6, -6}, 12},
		{"vertical, positive", Point{0, 7}, 7},
		{"vertical, negative", Point{0, -7}, 7},
		{"horizontal, positive", Point{8, 0}, 8},
		{"horizontal, negative", Point{-8, 0}, 8},
		{"beacon from day 15", Point{2, 10}.Sub(Point{8, 7}), 9},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := ManhattanLength(tc.in)
			if got != tc.want {
				t.Logf("ManhattanLength(%v) = %d; want %d", tc.in, got, tc.want)
				t.Fail()
			}
		})
	}
}

func TestRot90(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name string
		in   Point
		want Point
	}{
		{
			name: "0 degrees becomes 90 degrees",
			in:   Point{X: 1},
			want: Point{Y: 1},
		},
		{
			name: "90 degrees becomes 180 degrees",
			in:   Point{Y: 1},
			want: Point{X: -1},
		},
		{
			name: "180 degrees becomes 270 degrees",
			in:   Point{X: -1},
			want: Point{Y: -1},
		},
		{
			name: "270 degrees becomes 0 degrees",
			in:   Point{Y: -1},
			want: Point{X: 1},
		},
		{
			name: "Rotating the sum of both unit vectors",
			in:   Point{X: 1, Y: 1},
			want: Point{X: -1, Y: 1},
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := tc.in.Rot90()
			if got != tc.want {
				t.Logf("%v.Rot90() = %v ; want %v", tc.in, got, tc.want)
				t.Fail()
			}
		})
	}
}

func TestRot270(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name string
		in   Point
		want Point
	}{
		{
			name: "0 degrees becomes 270 degrees",
			in:   Point{X: 1},
			want: Point{Y: -1},
		},
		{
			name: "90 degrees becomes 0 degrees",
			in:   Point{Y: 1},
			want: Point{X: 1},
		},
		{
			name: "180 degrees becomes 90 degrees",
			in:   Point{X: -1},
			want: Point{Y: 1},
		},
		{
			name: "270 degrees becomes 180 degrees",
			in:   Point{Y: -1},
			want: Point{X: -1},
		},
		{
			name: "Rotating the sum of both unit vectors",
			in:   Point{X: 1, Y: 1},
			want: Point{X: 1, Y: -1},
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := tc.in.Rot270()
			if got != tc.want {
				t.Logf("%v.Rot270() = %v ; want %v", tc.in, got, tc.want)
				t.Fail()
			}
		})
	}
}
