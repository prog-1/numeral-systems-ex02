package main

import "testing"

func TestToBase10(t *testing.T) {
	for _, tc := range []struct {
		testn int
		num   string
		base  int
		want  int
	}{
		{testn: 1, num: "", base: 0, want: 0},
		{testn: 2, num: "1000111", base: 2, want: 71},
		{testn: 3, num: "3245607", base: 8, want: 871303},
		{testn: 4, num: "ABCDEF", base: 16, want: 11259375},
		{testn: 5, num: "Z", base: 36, want: 35},
	} {
		if got := toBase10(tc.num, tc.base); got != tc.want {
			t.Errorf("Error: test toBase10 number %v; got = %v, want = %v", tc.testn, got, tc.want)
		}
	}
}
func TestFromBase10(t *testing.T) {
	for _, tc := range []struct {
		testn int
		num   int
		base  int
		want  string
	}{
		{testn: 1, num: 0, base: 2, want: "0"},
		{testn: 2, num: 21, base: 3, want: "210"},
		{testn: 3, num: 133292, base: 16, want: "208AC"},
		{testn: 4, num: 8132, base: 8, want: "17704"},
		{testn: 5, num: 322, base: 2, want: "101000010"},
	} {
		if got := fromBase10(tc.num, tc.base); got != tc.want {
			t.Errorf("Error: test fromBase10 number %v; got = %v, want = %v", tc.testn, got, tc.want)
		}
	}
}
