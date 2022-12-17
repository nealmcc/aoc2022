package main

import (
	"strings"
	"testing"
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

func TestNewGraph(t *testing.T) {
	t.Parallel()
	valves, err := read(strings.NewReader(_sample))
	if err != nil {
		t.Log("error reading sample:", err)
		t.FailNow()
	}

	routes := NewGraph(valves)
	// here we don't check the 'previous' value, because in the case of a cycle,
	// there is more than one valid route.
	want := Graph{
		ID("AA"): map[ValveID]Marker{
			ID("BB"): {Dist: 1}, ID("CC"): {Dist: 2}, ID("DD"): {Dist: 1},
			ID("EE"): {Dist: 2}, ID("FF"): {Dist: 3}, ID("GG"): {Dist: 4},
			ID("HH"): {Dist: 5}, ID("II"): {Dist: 1}, ID("JJ"): {Dist: 2},
		},
		ID("BB"): map[ValveID]Marker{
			ID("AA"): {Dist: 1}, ID("CC"): {Dist: 1}, ID("DD"): {Dist: 2},
			ID("EE"): {Dist: 3}, ID("FF"): {Dist: 4}, ID("GG"): {Dist: 5},
			ID("HH"): {Dist: 6}, ID("II"): {Dist: 2}, ID("JJ"): {Dist: 3},
		},
		ID("CC"): map[ValveID]Marker{
			ID("AA"): {Dist: 2}, ID("BB"): {Dist: 1}, ID("DD"): {Dist: 1},
			ID("EE"): {Dist: 2}, ID("FF"): {Dist: 3}, ID("GG"): {Dist: 4},
			ID("HH"): {Dist: 5}, ID("II"): {Dist: 3}, ID("JJ"): {Dist: 4},
		},
		ID("DD"): map[ValveID]Marker{
			ID("AA"): {Dist: 1}, ID("BB"): {Dist: 2}, ID("CC"): {Dist: 1},
			ID("EE"): {Dist: 1}, ID("FF"): {Dist: 2}, ID("GG"): {Dist: 3},
			ID("HH"): {Dist: 4}, ID("II"): {Dist: 2}, ID("JJ"): {Dist: 3},
		},
		ID("EE"): map[ValveID]Marker{
			ID("AA"): {Dist: 2}, ID("BB"): {Dist: 3}, ID("CC"): {Dist: 2},
			ID("DD"): {Dist: 1}, ID("FF"): {Dist: 1}, ID("GG"): {Dist: 2},
			ID("HH"): {Dist: 3}, ID("II"): {Dist: 3}, ID("JJ"): {Dist: 4},
		},
		ID("FF"): map[ValveID]Marker{
			ID("AA"): {Dist: 3}, ID("BB"): {Dist: 4}, ID("CC"): {Dist: 3},
			ID("DD"): {Dist: 2}, ID("EE"): {Dist: 1}, ID("GG"): {Dist: 1},
			ID("HH"): {Dist: 2}, ID("II"): {Dist: 4}, ID("JJ"): {Dist: 5},
		},
		ID("GG"): map[ValveID]Marker{
			ID("AA"): {Dist: 4}, ID("BB"): {Dist: 5}, ID("CC"): {Dist: 4},
			ID("DD"): {Dist: 3}, ID("EE"): {Dist: 2}, ID("FF"): {Dist: 1},
			ID("HH"): {Dist: 1}, ID("II"): {Dist: 5}, ID("JJ"): {Dist: 6},
		},
		ID("HH"): map[ValveID]Marker{
			ID("AA"): {Dist: 5}, ID("BB"): {Dist: 6}, ID("CC"): {Dist: 5},
			ID("DD"): {Dist: 4}, ID("EE"): {Dist: 3}, ID("FF"): {Dist: 2},
			ID("GG"): {Dist: 1}, ID("II"): {Dist: 6}, ID("JJ"): {Dist: 7},
		},
		ID("II"): map[ValveID]Marker{
			ID("AA"): {Dist: 1}, ID("BB"): {Dist: 2}, ID("CC"): {Dist: 3},
			ID("DD"): {Dist: 2}, ID("EE"): {Dist: 3}, ID("FF"): {Dist: 4},
			ID("GG"): {Dist: 5}, ID("HH"): {Dist: 6}, ID("JJ"): {Dist: 1},
		},
		ID("JJ"): map[ValveID]Marker{
			ID("AA"): {Dist: 2}, ID("BB"): {Dist: 3}, ID("CC"): {Dist: 4},
			ID("DD"): {Dist: 3}, ID("EE"): {Dist: 4}, ID("FF"): {Dist: 5},
			ID("GG"): {Dist: 6}, ID("HH"): {Dist: 7}, ID("II"): {Dist: 1},
		},
	}

	for k1, v1 := range want {
		for k2, v2 := range v1 {
			got := routes[k1][k2].Dist
			if got != v2.Dist {
				t.Logf("route[%v][%v].Dist = %d; want %d", k1, k2, got, v2.Dist)
				t.Fail()
			}
		}
	}
}
