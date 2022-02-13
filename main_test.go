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
		{"", 0, 0},
		{"", 1, 0},
		{"", 16, 0},
		{"123", 10, 123},
		{"111101", 2, 61},
		{"ABCD", 16, 43981},
	} {
		t.Run(tc.num, func(t *testing.T) {
			if got := tobase10(tc.num, tc.fromBase); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
func TestFromBase10(t *testing.T) {
	for _, tc := range []struct {
		num      int
		fromBase int
		want     string
	}{
		{0, 0, ""},
		{0, 1, ""},
		{0, 16, ""},
		{123, 10, "123"},
		{61, 2, "111101"},
		{43981, 16, "ABCD"},
	} {
		t.Run(tc.want, func(t *testing.T) {
			if got := fromBase10(tc.num, tc.fromBase); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
