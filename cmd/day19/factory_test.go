package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseBlueprint(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name string
		in   string
		want Blueprint
		cost map[Robot][4]int
	}{
		{
			name: "sample row 1",
			in:   `Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.`,
			want: Blueprint{4, 2, 3, 2, 14, 7},
			cost: map[Robot][4]int{
				Orebot:   {Oremat: 4},
				Claybot:  {Oremat: 2},
				Obsbot:   {Oremat: 3, Claymat: 14},
				Geodebot: {Oremat: 2, Obsmat: 7},
			},
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			bp, err := ParseBlueprint(tc.in)
			if err != nil {
				t.Log(err)
				t.FailNow()
			}
			if bp != tc.want {
				t.Logf("ParseBlueprint() = %v ; want %v", bp, tc.want)
				t.FailNow()
			}

			for bot, want := range tc.cost {
				got := bp.cost(bot)
				assert.Equalf(t, want, got, "%s", bot)
			}
		})
	}
}
