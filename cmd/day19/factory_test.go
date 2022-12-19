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
		want blueprint
		cost map[robot]map[material]int
	}{
		{
			name: "sample row 1",
			in:   `Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.`,
			want: blueprint{4, 2, 3, 14, 2, 7},
			cost: map[robot]map[material]int{
				OreBot:   {OreMat: 4},
				ClayBot:  {OreMat: 2},
				ObsBot:   {OreMat: 3, ClayMat: 14},
				GeodeBot: {OreMat: 2, ObsMat: 7},
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
				got := bp.Cost(bot)
				assert.Equalf(t, want, got, "%s", bot)
			}
		})
	}
}
