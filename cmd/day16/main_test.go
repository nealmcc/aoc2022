package main

import (
	"strings"
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

func TestDist(t *testing.T) {
	valves, err := read(strings.NewReader(_sample))
	if err != nil {
		t.Log("error reading sample:", err)
		t.FailNow()
	}

	got := distances(valves)
	want := map[ValveID]map[ValveID]int{
		ID("AA"): {
			ID("AA"): 0, ID("BB"): 1, ID("CC"): 2, ID("DD"): 1, ID("EE"): 2,
			ID("FF"): 3, ID("GG"): 4, ID("HH"): 5, ID("II"): 1, ID("JJ"): 2,
		},
		ID("BB"): {
			ID("AA"): 1, ID("BB"): 0, ID("CC"): 1, ID("DD"): 2, ID("EE"): 3,
			ID("FF"): 4, ID("GG"): 5, ID("HH"): 6, ID("II"): 2, ID("JJ"): 3,
		},
		ID("CC"): {
			ID("AA"): 2, ID("BB"): 1, ID("CC"): 0, ID("DD"): 1, ID("EE"): 2,
			ID("FF"): 3, ID("GG"): 4, ID("HH"): 5, ID("II"): 3, ID("JJ"): 4,
		},
		ID("DD"): {
			ID("AA"): 1, ID("BB"): 2, ID("CC"): 1, ID("DD"): 0, ID("EE"): 1,
			ID("FF"): 2, ID("GG"): 3, ID("HH"): 4, ID("II"): 2, ID("JJ"): 3,
		},
		ID("EE"): {
			ID("AA"): 2, ID("BB"): 3, ID("CC"): 2, ID("DD"): 1, ID("EE"): 0,
			ID("FF"): 1, ID("GG"): 2, ID("HH"): 3, ID("II"): 3, ID("JJ"): 4,
		},
		ID("FF"): {
			ID("AA"): 3, ID("BB"): 4, ID("CC"): 3, ID("DD"): 2, ID("EE"): 1,
			ID("FF"): 0, ID("GG"): 1, ID("HH"): 2, ID("II"): 4, ID("JJ"): 5,
		},
		ID("GG"): {
			ID("AA"): 4, ID("BB"): 5, ID("CC"): 4, ID("DD"): 3, ID("EE"): 2,
			ID("FF"): 1, ID("GG"): 0, ID("HH"): 1, ID("II"): 5, ID("JJ"): 6,
		},
		ID("HH"): {
			ID("AA"): 5, ID("BB"): 6, ID("CC"): 5, ID("DD"): 4, ID("EE"): 3,
			ID("FF"): 2, ID("GG"): 1, ID("HH"): 0, ID("II"): 6, ID("JJ"): 7,
		},
		ID("II"): {
			ID("AA"): 1, ID("BB"): 2, ID("CC"): 3, ID("DD"): 2, ID("EE"): 3,
			ID("FF"): 4, ID("GG"): 5, ID("HH"): 6, ID("II"): 0, ID("JJ"): 1,
		},
		ID("JJ"): {
			ID("AA"): 2, ID("BB"): 3, ID("CC"): 4, ID("DD"): 3, ID("EE"): 4,
			ID("FF"): 5, ID("GG"): 6, ID("HH"): 7, ID("II"): 1, ID("JJ"): 0,
		},
	}
	assert.Equal(t, want, got)
}

// 	got, want := part1(valves, 30), 1651
// 	if got != want {
// 		t.Logf("part1() = %d; want %d", got, want)
// 		t.Fail()
// 	}
// }
