package main

import (
	"strings"
	"testing"

	v "github.com/nealmcc/aoc2022/pkg/vector/twod"
)

func TestStorm_IceAt(t *testing.T) {
	t.Parallel()

	storm, err := parse(strings.NewReader(`#.#####
#.....#
#>....#
#.....#
#...v.#
#.....#
#####.#`))
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	tt := []struct {
		name   string
		pos    v.Point
		t      int
		want   Ice
		wantOk bool
	}{
		{
			name:   "time 0, ice A",
			pos:    v.Point{X: 0, Y: 1},
			t:      0,
			want:   East,
			wantOk: true,
		},
		{
			name:   "time 0, ice B",
			pos:    v.Point{X: 3, Y: 3},
			t:      0,
			want:   South,
			wantOk: true,
		},
		{
			name:   "time 1, ice A",
			pos:    v.Point{X: 1, Y: 1},
			t:      1,
			want:   East,
			wantOk: true,
		},
		{
			name:   "time 1, ice B",
			pos:    v.Point{X: 3, Y: 4},
			t:      1,
			want:   South,
			wantOk: true,
		},
		{
			name:   "time 2, ice B wrapped around",
			pos:    v.Point{X: 3, Y: 0},
			t:      2,
			want:   South,
			wantOk: true,
		},
		{
			name:   "time 3, ice A and B merged",
			pos:    v.Point{X: 3, Y: 1},
			t:      3,
			want:   South + East,
			wantOk: true,
		},
		{
			name:   "time 4, ice A at the right edge",
			pos:    v.Point{X: 4, Y: 1},
			t:      4,
			want:   East,
			wantOk: true,
		},
		{
			name:   "time 5, ice A wrapped around",
			pos:    v.Point{X: 0, Y: 1},
			t:      5,
			want:   East,
			wantOk: true,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, ok := storm.IceAt(tc.pos, tc.t)
			if got != tc.want || ok != tc.wantOk {
				t.Logf("storm.IceAt(%v, %v) = %d, %t ; want %d, %t",
					tc.pos, tc.t, got, ok, tc.want, tc.wantOk)
				t.Fail()
			}
		})
	}
}
