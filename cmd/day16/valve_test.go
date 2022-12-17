package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseValve(t *testing.T) {
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

			got, err := ParseValve([]byte(tc.in))
			if err != nil {
				t.Log("unexpected error", err)
				t.FailNow()
			}
			assert.Equal(t, tc.want, got)
		})
	}
}
