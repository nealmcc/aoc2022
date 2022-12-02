package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddScores(t *testing.T) {
	t.Parallel()

	input := `A Y
B X
C Z`

	p1, err := addScores(strings.NewReader(input), part1)
	require.NoError(t, err)
	assert.Equal(t, 15, p1)

	p2, err := addScores(strings.NewReader(input), part2)
	require.NoError(t, err)
	assert.Equal(t, 12, p2)
}
