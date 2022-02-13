package main

import (
	"testing"
)

func TestToBase10(t *testing.T) {
	for _, tc := range []struct {
		num      string
		fromBase int
		want     int
	}{
		{num: "", fromBase: 0, want: 0},
		{num: "", fromBase: 1, want: 0},
		{num: "1", fromBase: 0, want: 1},
		{num: "123", fromBase: 10, want: 123},
		{num: "111101", fromBase: 2, want: 61},
		{num: "ABCD", fromBase: 16, want: 43981},
		{num: "331", fromBase: 4, want: 61},
		{num: "ABCD", fromBase: 20, want: 84653},
		{num: "1666", fromBase: 12, want: 2670},
		{num: "Z3", fromBase: 36, want: 1263},
	} {
		t.Run(tc.num, func(t *testing.T) {
			if got := toBase10(tc.num, tc.fromBase); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}

func TestFromBase10(t *testing.T) {
	for _, tc := range []struct {
		num    int
		toBase int
		want   string
	}{
		{num: 0, toBase: 6, want: "0"},
		{num: 123, toBase: 10, want: "123"},
		{num: 61, toBase: 2, want: "111101"},
		{num: 43981, toBase: 16, want: "ABCD"},
		{num: 159, toBase: 16, want: "9F"},
		{num: 12222100012, toBase: 8, want: "133037471054"},
		{num: 777, toBase: 2, want: "1100001001"},
		{num: 112345, toBase: 32, want: "3DMP"},
		{num: 230, toBase: 19, want: "C2"},
	} {
		got := fromBase10(tc.num, tc.toBase)
		if got != tc.want {
			t.Errorf("ERR: fromBase10(%v, %v): got: %s, want: %s", tc.num, tc.toBase, got, tc.want)
		}
	}
}
