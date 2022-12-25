package main

import (
	"fmt"
	"strings"
	"testing"
)

var _examples = []struct {
	snafu   string
	decimal int
}{
	{snafu: "1=-0-2", decimal: 1747},
	{snafu: " 12111", decimal: 906},
	{snafu: "  2=0=", decimal: 198},
	{snafu: "    21", decimal: 11},
	{snafu: "  2=01", decimal: 201},
	{snafu: "   111", decimal: 31},
	{snafu: " 20012", decimal: 1257},
	{snafu: "   112", decimal: 32},
	{snafu: " 1=-1=", decimal: 353},
	{snafu: "  1-12", decimal: 107},
	{snafu: "    12", decimal: 7},
	{snafu: "    1=", decimal: 3},
	{snafu: "   122", decimal: 37},
}

func TestSNAFUToInt(t *testing.T) {
	for _, tc := range _examples {
		tc := tc
		t.Run(fmt.Sprintf("%d", tc.decimal), func(t *testing.T) {
			t.Parallel()
			got := SNAFUToInt(tc.snafu)
			if got != tc.decimal {
				t.Logf("SNAFUToInt(%s) = %d ; want %d", tc.snafu, got, tc.decimal)
				t.Fail()
			}
		})
	}
}

func TestIntToSNAFU(t *testing.T) {
	for _, tc := range _examples {
		tc := tc
		t.Run(fmt.Sprintf("%d", tc.decimal), func(t *testing.T) {
			t.Parallel()
			got := IntToSNAFU(tc.decimal)
			if got != strings.TrimSpace(tc.snafu) {
				t.Logf("IntToSNAFU(%d) = %q ; want %q",
					tc.decimal, got, strings.TrimSpace(tc.snafu))
				t.Fail()
			}
		})
	}
}
