package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var _sample = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`

func TestReadElves(t *testing.T) {
	t.Parallel()

	input := strings.NewReader(_sample)

	got, err := readElves(input)
	require.NoError(t, err)

	want := []elf{
		{
			food:     []int{1000, 2000, 3000},
			calories: 6000,
		},
		{
			food:     []int{4000},
			calories: 4000,
		},
		{
			food:     []int{5000, 6000},
			calories: 11000,
		},
		{
			food:     []int{7000, 8000, 9000},
			calories: 24000,
		},
		{
			food:     []int{10000},
			calories: 10000,
		},
	}
	assert.Equal(t, want, got)
}

func TestReadElf(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name    string
		in      string
		want    elf
		wantOK  bool
		wantErr bool
	}{
		{
			name:   "empty string => no elf",
			wantOK: false,
		},
		{
			name: "one line => one elf",
			in:   "1000",
			want: elf{
				food:     []int{1000},
				calories: 1000,
			},
			wantOK: true,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			r := require.New(t)

			s := bufio.NewScanner(strings.NewReader(tc.in))
			got, ok, err := readElf(s)
			if tc.wantErr {
				r.Error(err)
			} else {
				r.NoError(err)
			}

			r.Equal(tc.wantOK, ok, "okay should match")
			r.ElementsMatch(tc.want.food, got.food, "elf's food should match")
			r.Equal(tc.want.calories, got.calories, "elf's calories should match")
		})
	}
}
