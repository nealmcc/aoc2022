package main

import (
	"os"
	"strings"
	"testing"
)

const _sample = `30373
25512
65332
33549
35390
`

func TestPart1(t *testing.T) {
	t.Parallel()

	trees, err := NewForest(strings.NewReader(_sample))
	if err != nil {
		t.Logf("error reading sample")
		t.FailNow()
	}

	got, want := trees.Visibility(), 21
	t.Logf("trees: \n%s", _sample)
	t.Logf("overall: \n%s", got)
	t.Logf("from the top: \n%s", got.Filter(top))
	t.Logf("from the right: \n%s", got.Filter(right))
	t.Logf("from the bottom: \n%s", got.Filter(bottom))
	t.Logf("from the left: \n%s", got.Filter(left))
	if len(got) != 21 {
		t.Logf("part1() = has length of %d; want %d", len(got), want)
		t.Fail()
	}
}

func TestSize(t *testing.T) {
	t.Parallel()

	trees, err := NewForest(strings.NewReader(_sample))
	if err != nil {
		t.Logf("error reading sample")
		t.FailNow()
	}

	w, h := trees.Size()
	if w != 5 || h != 5 {
		t.Logf("size() = %d, %d ; want %d, %d", w, h, 5, 5)
		t.Fail()
	}

	file, err := os.Open("input.txt")
	if err != nil {
		t.Fatal(err)
	}
	trees, err = NewForest(file)
	if err != nil {
		t.Logf("error reading input file")
		t.Fail()
	}

	w, h = trees.Size()
	if w != 99 || h != 99 {
		t.Logf("size() = %d, %d ; want %d, %d", w, h, 99, 99)
		t.Fail()
	}
}

func TestInputLines(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name string
		inpt string
		want string
	}{
		{
			name: "first line of input",
			inpt: "102212110110302230012132130441442043243242145112422525112333344240121120342411110001231112222021211",
			want: "d5d555555555d55555555555555d5555555555555555d555555557555555555555555555555755555555575555555555757",
		},
		{
			name: "last line of input",
			inpt: "011011030023132123022011022032214014533223521512511221453113132215230244324434210342210100311021222",
			want: "dd55555d555555555555555555555555d555d55555555555555555555555555557555555555555555575555555755555557",
		},
		{
			name: "line 71 of the input",
			inpt: `121244523334642663334357574375756475655847758985795997565974854847688754735474744655256243422233523`,
			want: `dd55d5d55555d5555555555d555555555555555d55555d55555555555755555555557555555555755555557555555555757`,
		},
		{
			name: "10th column of input",
			inpt: `1
3
1
0
3
2
2
1
3
0
0
3
4
0
3
4
2
0
2
0
1
5
4
1
2
1
4
3
1
5
5
4
3
3
1
5
3
6
5
2
6
5
2
4
4
4
2
6
5
3
3
5
2
2
6
3
6
2
3
6
6
4
4
3
1
4
1
1
4
5
3
3
3
2
4
2
2
4
1
1
0
3
4
4
0
4
2
4
4
2
4
3
1
2
2
1
2
2
0`,
			want: `b
b
a
a
a
a
a
a
a
a
a
a
b
a
a
a
a
a
a
a
a
b
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
b
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
e
a
a
a
a
a
a
a
a
e
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
a
e
e
a
a
a
a
a
e
e`,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			f, _ := NewForest(strings.NewReader(tc.inpt))
			got := f.Visibility()

			t.Logf("from the top:\n%s", got.Filter(top))
			t.Logf("from the right:\n%s", got.Filter(right))
			t.Logf("from the bottom:\n%s", got.Filter(bottom))
			t.Logf("from the left:\n%s", got.Filter(left))

			// strip off the trailing newline in the output format:
			text := got.String()
			text = text[:len(text)-1]
			if text != tc.want {
				t.Logf("\ninput:\t%s\n  got:\t%s\n want:\t%s", tc.inpt, text, tc.want)
				t.Fail()
			}
		})
	}
}
