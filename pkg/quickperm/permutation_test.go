package quickperm

import (
	"testing"
)

func TestPermutations(t *testing.T) {
	unique := make(map[string]struct{})
	for val := range Permutations([]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}) {
		s := string(val)
		t.Log(s)
		unique[s] = struct{}{}
	}

	want := 8 * 7 * 6 * 5 * 4 * 3 * 2
	if len(unique) != want {
		t.Logf("got %d unique items; want %d", len(unique), want)
		t.Fail()
	}
}
