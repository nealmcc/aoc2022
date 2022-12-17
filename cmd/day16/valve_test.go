package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const _sample = `Valve AA has flow rate=0; tunnels lead to valves DD, II, BB
Valve BB has flow rate=13; tunnels lead to valves CC, AA
Valve CC has flow rate=2; tunnels lead to valves DD, BB
Valve DD has flow rate=20; tunnels lead to valves CC, AA, EE
Valve EE has flow rate=3; tunnels lead to valves FF, DD
Valve FF has flow rate=0; tunnels lead to valves EE, GG
Valve GG has flow rate=0; tunnels lead to valves FF, HH
Valve HH has flow rate=22; tunnel leads to valve GG
Valve II has flow rate=0; tunnels lead to valves AA, JJ
Valve JJ has flow rate=21; tunnel leads to valve II
`

/*

   AA --- BB --- CC
    0     13      2
   |  \        /
   |   +- DD -+- EE --- FF --- GG --- HH
   |      20      3      0      0     22
   |
   II --- JJ
    0     21

*/

func TestParseLine(t *testing.T) {
	t.Parallel()

	tt := []struct {
		in   string
		want Valve
	}{
		{
			`Valve AA has flow rate=0; tunnels lead to valves DD, II, BB`,
			Valve{
				ID:         K("AA"),
				Flow:       0,
				Neighbours: []ValveID{K("DD"), K("II"), K("BB")},
			},
		},
		{
			`Valve JJ has flow rate=21; tunnel leads to valve II`,
			Valve{
				ID:         K("JJ"),
				Flow:       21,
				Neighbours: []ValveID{K("II")},
			},
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.in, func(t *testing.T) {
			t.Parallel()

			got, err := ParseValve(tc.in)
			if err != nil {
				t.Log("unexpected error", err)
				t.FailNow()
			}
			assert.Equal(t, tc.want, got)
		})
	}
}
