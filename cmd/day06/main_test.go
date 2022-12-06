package main

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	t.Parallel()
	tt := []struct {
		in   string
		want int
	}{
		{
			in:   "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			want: 7,
		},
		{
			in:   "bvwbjplbgvbhsrlpgdmjqwftvncz",
			want: 5,
		},
		{
			in:   "nppdvjthqldpwncqszvftbrmjlhg",
			want: 6,
		},
		{
			in:   "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			want: 10,
		},
		{
			in:   "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			want: 11,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.in, func(t *testing.T) {
			t.Parallel()

			got, err := part1([]byte(tc.in))
			if err != nil {
				t.Log(err)
				t.FailNow()
			}
			if got != tc.want {
				t.Logf("part1(%s) = %d ; want %d", tc.in, got, tc.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	t.Parallel()
	tt := []struct {
		in   string
		want int
	}{
		{
			in:   "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			want: 19,
		},
		{
			in:   "bvwbjplbgvbhsrlpgdmjqwftvncz",
			want: 23,
		},
		{
			in:   "nppdvjthqldpwncqszvftbrmjlhg",
			want: 23,
		},
		{
			in:   "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			want: 29,
		},
		{
			in:   "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			want: 26,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.in, func(t *testing.T) {
			t.Parallel()

			got, err := part2([]byte(tc.in))
			if err != nil {
				t.Log(err)
				t.FailNow()
			}
			if got != tc.want {
				t.Logf("part1(%s) = %d ; want %d", tc.in, got, tc.want)
			}
		})
	}
}

var _result int

func BenchmarkPart1(b *testing.B) {
	data, _ := os.ReadFile("input.txt")
	var p1 int
	for n := 0; n < b.N; n++ {
		// always record the result of part1 to prevent
		// the compiler eliminating the function call.
		p1, _ = part1(data)
	}
	// storing the result in a package-level variable prevents the
	// compiler from eliminating the Benchmark
	_result = p1
}

func BenchmarkPart2(b *testing.B) {
	data, _ := os.ReadFile("input.txt")
	var p2 int
	for n := 0; n < b.N; n++ {
		p2, _ = part2(data)
	}
	_result = p2
}
