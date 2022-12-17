package collection

import (
	"strings"
	"testing"
)

func TestSet_Add(t *testing.T) {
	s := make(Set[int])
	s.Add(1)
	s.Add(2)
	s.Add(3)
	got, want := len(s), 3
	if got != want {
		t.Logf("len = %d; want %d", got, want)
		t.Fail()
	}
}

func TestSet_Remove(t *testing.T) {
	s := make(Set[byte])
	s.Add('a')
	s.Add('b')
	s.Add('c')
	s.Remove('b')
	got, want := len(s), 2
	if got != want {
		t.Logf("len = %d; want %d", got, want)
		t.Fail()
	}
}

func TestSet_Contains(t *testing.T) {
	t.Run("empty set", func(t *testing.T) {
		var s Set[int]

		got, want := s.Contains(3), false
		if got != want {
			t.Logf("len = %t; want %t", got, want)
			t.Fail()
		}
	})

	t.Run("non-empty set", func(t *testing.T) {
		s := make(Set[int])
		s.Add(42)
		got, want := s.Contains(42), true
		if got != want {
			t.Logf("len = %t; want %t", got, want)
			t.Fail()
		}
	})
}

func TestSet_Init(t *testing.T) {
	var s Set[string]
	s.Init(strings.Split("The quick brown fox jumps over the lazy dog", " "))

	got, want := len(s), 9
	if got != want {
		t.Logf("len = %d; want %d", got, want)
		t.Fail()
	}

	if !s.Contains("dog") {
		t.Logf("%v.Contains(%q) = false ; want true", s, "dog")
		t.Fail()
	}
}

func TestSet_String(t *testing.T) {
	t.Parallel()
	tt := []struct {
		name string
		in   Set[int]
		want string
	}{
		{
			name: "nil set",
			in:   nil,
			want: "[]",
		},
		{
			name: "empty set",
			in:   Set[int]{},
			want: "[]",
		},
		{
			name: "one item",
			in:   Set[int]{1: struct{}{}},
			want: "[1]",
		},
		{
			name: "three items",
			in: Set[int]{
				1: struct{}{},
				2: struct{}{},
				3: struct{}{},
			},
			want: "[1, 2, 3]",
		},
		{
			name: "arbitrary order",
			in: Set[int]{
				3: struct{}{},
				1: struct{}{},
				2: struct{}{},
			},
			want: "[1, 2, 3]",
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := tc.in.String()
			if got != tc.want {
				t.Logf("got %s; want %s", got, tc.want)
				t.Fail()
			}
		})
	}
}
