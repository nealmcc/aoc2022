package main

import (
	"os"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	file, err := os.Open("sample2.txt")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	storm, err := parse(file)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	got, want := part1(storm), 18

	if got != want {
		t.Logf("part1(sample2) = %d ; want %d", got, want)
		t.Fail()
	}
}

func TestPart2(t *testing.T) {
	t.Skip("not ready")
	file, err := os.Open("sample2.txt")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	storm, err := parse(file)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	got, want := part2(storm), -1

	if got != want {
		t.Logf("part2(sample2) = %d ; want %d", got, want)
		t.Fail()
	}
}

func TestParse(t *testing.T) {
	tt := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "small sample",
			in: `#.#####
#.....#
#>....#
#.....#
#...v.#
#.....#
#####.#`,
			want: `#.#####
#.....#
#>....#
#.....#
#...v.#
#.....#
#####.#`,
		},
		{
			name: "larger sample",
			in: `#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#`,
			want: `#.######
#>>.<^<#
#.<..<<#
#>v.><>#
#<^v^^>#
######.#`,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			storm, err := parse(strings.NewReader(tc.in))
			if err != nil {
				t.Log(err)
				t.FailNow()
			}

			got := storm.String()
			if got != tc.want {
				t.Logf("parsed storm: \n%v\ngot:\n%v\nwant:\n%v\n",
					tc.in, got, tc.want)
				t.Fail()
			}
		})
	}
}
