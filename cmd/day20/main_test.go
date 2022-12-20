package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var _sample = []int{1, 2, -3, 3, -2, 0, 4}

func TestPart1(t *testing.T) {
	got, want := part1(_sample), 3
	if got != want {
		t.Logf("part1(sample) = %d ; want %d", got, want)
		t.Fail()
	}
}

func TestNewList(t *testing.T) {
	list := NewList(_sample, 0)

	if list.Len() != len(_sample) {
		t.Logf("got list.Len() = %d; want %d", list.Len(), len(_sample))
		t.Fail()
	}

	got := list.Root.Len()
	if got != len(_sample) {
		t.Logf("got list.Root.Len() = %d; want %d", got, len(_sample))
		t.Fail()
	}

	want := []int{0, 4, 1, 2, -3, 3, -2}
	assert.Equal(t, want, list.Normal())

	if list.offset != 5 {
		t.Logf("got list.offset = %d; want %d", list.offset, 5)
		t.Fail()
	}

	assert.Equal(t, _sample, list.Fixed())
}

func TestMix(t *testing.T) {
	list := NewList(_sample, 0)
	list.Mix()

	got := list.Normal()
	want := []int{0, 3, -2, 1, 2, -3, 4}
	assert.Equal(t, want, got)
}

func TestMix_edge_cases(t *testing.T) {
	tt := []struct {
		name string
		in   []int
		want []int
	}{
		{
			name: "postive number one less than the size of the input",
			in:   []int{0, 1, 2},
			want: []int{0, 2, 1},
		},
		{
			name: "postive number equal to the size of the input",
			in:   []int{0, 1, 1, 1, 1, 6},
			want: []int{0, 6, 1, 1, 1, 1},
		},
		{
			name: "postive number one larger than the size of the input",
			in:   []int{0, 1, 1, 5},
			want: []int{0, 1, 5, 1},
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			// t.Parallel()
			list := NewList(tc.in, 0)
			list.Mix()
			assert.Equal(t, tc.want, list.Fixed())
		})
	}
}
