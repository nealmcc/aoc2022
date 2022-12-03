package main

import (
	"strings"
	"testing"
)

var _sample = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

func TestPart1(t *testing.T) {
	t.Parallel()

	got, err := part1(strings.NewReader(_sample))
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	want := 157
	if got != want {
		t.Logf("part1() = %d ; want %d", got, want)
		t.Fail()
	}
}

func TestPart2(t *testing.T) {
	t.Parallel()

	got, err := part2(strings.NewReader(_sample))
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	want := 70
	if got != want {
		t.Logf("part2() = %d ; want %d", got, want)
		t.Fail()
	}
}

func TestPriority(t *testing.T) {
	t.Parallel()

	t.Run("lowercase", func(t *testing.T) {
		t.Parallel()

		var item item
		var want int
		for item, want = 'a', 1; item <= 'z'; item, want = item+1, want+1 {
			got := item.priority()
			if want != got {
				t.Logf("priority(%b) = %d ; want %d", item, got, want)
				t.Fail()
			}
		}
	})

	t.Run("uppercase", func(t *testing.T) {
		t.Parallel()

		var item item
		var want int
		for item, want = 'A', 27; item <= 'Z'; item, want = item+1, want+1 {
			got := item.priority()
			if got != want {
				t.Logf("priority(%c) = %d ; want %d", item, got, want)
				t.Fail()
			}
		}
	})
}

func TestFindCommonItem(t *testing.T) {
	t.Parallel()

	all := bag("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	tt := []struct {
		name string
		in   []bag
		want item
	}{
		{
			name: "lowercase a",
			in: []bag{
				all,
				bag("abc"),
				bag("cdZa"),
				bag("CdZa"),
			},
			want: 'a',
		},
		{
			name: "lowercase z",
			in: []bag{
				all,
				bag("bcz"),
				bag("cdzZa"),
				bag("CdzZa"),
			},
			want: 'z',
		},
		{
			name: "uppercase A",
			in: []bag{
				all,
				bag("Abcz"),
				bag("AdzZa"),
				bag("AdzZa"),
			},
			want: 'A',
		},
		{
			name: "uppercase Z",
			in: []bag{
				all,
				bag("abZz"),
				bag("dzZa"),
				bag("dzZa"),
			},
			want: 'Z',
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := findCommonItem(tc.in...)
			if err != nil {
				t.Log(err)
				t.FailNow()
			}

			if got != tc.want {
				t.Logf("findCommonItem() = %c ; want %c", got, tc.want)
				t.Fail()
			}
		})
	}
}
