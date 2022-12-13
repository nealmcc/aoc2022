package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const _sample = `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]
`

func TestRead(t *testing.T) {
	t.Parallel()

	lines, err := read(strings.NewReader(_sample))
	if err != nil {
		t.Log("error reading sample", err)
		t.FailNow()
	}
	pairs := makePairs(lines)

	a := assert.New(t)
	a.Equal(8, len(pairs))
	a.Equal(data{
		data{1.0, 1.0, 3.0, 1.0, 1.0},
		data{1.0, 1.0, 5.0, 1.0, 1.0},
	}, pairs[0])
}

func TestPart1(t *testing.T) {
	t.Parallel()

	lines, err := read(strings.NewReader(_sample))
	if err != nil {
		t.Log("error reading sample", err)
		t.FailNow()
	}

	got, want := part1(lines), 13
	if got != want {
		t.Logf("part1() =  %d; want %d", got, want)
		t.Fail()
	}
}

func TestPart2(t *testing.T) {
	t.Parallel()

	lines, err := read(strings.NewReader(_sample))
	if err != nil {
		t.Log("error reading sample", err)
		t.FailNow()
	}

	got, err := part2(lines)
	if err != nil {
		t.Log("error in part 2", err)
		t.FailNow()
	}

	want := 140
	if got != want {
		t.Logf("part2() =  %d; want %d", got, want)
		t.Fail()
	}
}

// var _result int // prevent the compiler from optimising away the call.

// func BenchmarkPart2(b *testing.B) {
// 	file, err := os.Open("input.txt")
// 	if err != nil {
// 		b.Fatal(err)
// 	}
// 	hill, err := read(file)
// 	if err != nil {
// 		b.Fatal(err)
// 	}

// 	var result int
// 	for n := 0; n < b.N; n++ {
// 		result = part2(hill)
// 	}
// 	_result = result
// }
