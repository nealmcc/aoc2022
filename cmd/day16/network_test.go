package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNetwork_TransitionMany(t *testing.T) {
	t.Parallel()
	valves, err := read(strings.NewReader(_sample))
	if err != nil {
		t.Log("error reading sample:", err)
		t.FailNow()
	}

	network := NewNetwork(valves)

	got := network.TransitionMany("DDBBJJHHEECC")

	want := state{
		curr:         ID("CC"),
		mins:         24, // one less minute has elapsed than the current 'minute number' in the example.
		totalFlow:    1165,
		missedFlow:   81*24 - 1165,
		currFlowRate: 81,
	}

	assert.Equal(t, want, got)
}
