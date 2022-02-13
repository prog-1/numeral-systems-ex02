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
		{"", 0, -1},
		{"", 1, -1},
		{"", 16, -1},
		{"1", 0, -1},
		{"321", 0, -1},
		{"013579", -1, -1},
		{"341B1", -195, -1},
		{"123", 10, 123},
		{"123", 3, -1},
		{"123", 4, 27},
		{"-123", 4, -1},
		{"111101", 2, 61},
		{"ABCD", 16, 43981},
		{"ABCD", 12, -1},
		{"1010111010", 1, -1},
		{"1010201100", 2, -1},
		{"01A2C", 14, 4744},
		{"19Z3", 36, 59583},
		{"0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ", 37, -1},
		{"123", 123, -1},
	} {
		t.Run(tc.num, func(t *testing.T) {
			if got := toBase10(tc.num, tc.fromBase); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
