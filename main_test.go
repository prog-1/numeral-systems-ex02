package main

import (
	"testing"
)

func TesttoBase10(t *testing.T) {
	for _, tc := range []struct {
		name     string
		num      string
		fromBase int
		want     int
	}{
		{"empty string", "", 8, 0},
		{"from README", "123", 10, 123},
		{"from README", "111101", 2, 61},
		{"from README", "abcd", 16, 43981},
		{"from base 9", "328", 9, 100001101},
		{"from base 4", "03", 4, 11},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := toBase10(tc.num, tc.fromBase); got != tc.want {
				t.Errorf("toBase10(%v, %v) = %v, want = %v", tc.num, tc.fromBase, got, tc.want)
			}
		})
	}
}

func TestfromBase10(t *testing.T) {
	for _, tc := range []struct {
		name   string
		num    int
		toBase int
		want   string
	}{
		{"from README", 123, 10, "123"},
		{"from README", 61, 2, "111101"},
		{"from README", 43981, 16, "abcd"},
		{"to base 5", 871, 5, "11441"},
		{"to base 16", 124313, 16, "1e599"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := fromBase10(tc.num, tc.toBase); got != tc.want {
				t.Errorf("fromBase10(%v, %v) = %v, want = %v", tc.num, tc.toBase, got, tc.want)
			}
		})
	}
}
