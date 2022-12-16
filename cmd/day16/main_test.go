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

func TestParseRow(t *testing.T) {
	t.Parallel()

	tt := []struct {
		in   string
		want Valve
	}{
		{
			`Valve AA has flow rate=0; tunnels lead to valves DD, II, BB`,
			Valve{
				ID:         ID("AA"),
				Flow:       0,
				Neighbours: []ValveID{ID("DD"), ID("II"), ID("BB")},
			},
		},
		{
			`Valve JJ has flow rate=21; tunnel leads to valve II`,
			Valve{
				ID:         ID("JJ"),
				Flow:       21,
				Neighbours: []ValveID{ID("II")},
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
			assert.Equal(t, tc.want, got)
		})
	}
}

// func TestPart1(t *testing.T) {
// 	t.Parallel()

// 	valves, err := read(strings.NewReader(_sample))
// 	if err != nil {
// 		t.Log("error reading sample:", err)
// 		t.FailNow()
// 	}

// 	got, want := part1(valves, 30), 1651
// 	if got != want {
// 		t.Logf("part1() = %d; want %d", got, want)
// 		t.Fail()
// 	}
// }
