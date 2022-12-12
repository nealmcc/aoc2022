package collection

import (
	"testing"

	v "github.com/nealmcc/aoc2022/pkg/vector/twod"
	"github.com/stretchr/testify/assert"
)

func TestSlice2D_Neighbours4(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name string
		size int
		in   v.Point
		want []v.Point
	}{
		{
			name: "one-square grid => no neighbours",
			size: 1,
			in:   v.Point{},
			want: []v.Point{},
		},
		{
			name: "four-square grid, bottom right => two neighbours",
			size: 2,
			in:   v.Point{X: 1, Y: 1},
			want: []v.Point{
				{X: 0, Y: 1},
				{X: 1, Y: 0},
			},
		},
		{
			name: "nine-square grid, center => four neighbours",
			size: 3,
			in:   v.Point{X: 1, Y: 1},
			want: []v.Point{
				{X: 1, Y: 0},
				{X: 0, Y: 1},
				{X: 2, Y: 1},
				{X: 1, Y: 2},
			},
		},
		{
			name: "nine-square grid, middle-right => three neighbours",
			size: 3,
			in:   v.Point{X: 2, Y: 1},
			want: []v.Point{
				{X: 2, Y: 0},
				{X: 1, Y: 1},
				{X: 2, Y: 2},
			},
		},
		{
			name: "nine-square grid, bottom-middle => three neighbours",
			size: 3,
			in:   v.Point{X: 1, Y: 2},
			want: []v.Point{
				{X: 1, Y: 1},
				{X: 0, Y: 2},
				{X: 2, Y: 2},
			},
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			grid := NewSlice2D[int](tc.size)
			got := grid.Neighbours4(tc.in)
			assert.ElementsMatch(t, tc.want, got)
		})
	}
}
