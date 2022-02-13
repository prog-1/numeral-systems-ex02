package main

import (
	"reflect"
	"testing"
)

func TestToBase10(t *testing.T) {
	for _, tc := range []struct {
		num      string
		fromBase int
		want     int
	}{
		{num: "", fromBase: 10, want: 0},
		{num: "123", fromBase: 10, want: 123},
		{num: "123AB", fromBase: 16, want: 74667},
		{num: "ZEW", fromBase: 36, want: 45896},
		{num: "1101", fromBase: 2, want: 13},
	} {
		if got := toBase10(tc.num, tc.fromBase); got != tc.want {
			t.Errorf("toBase10(%v, %v) = (%v), want = (%v)", tc.num, tc.fromBase, got, tc.want)
		}
	}
}
func TestFromBase10(t *testing.T) {
	for _, tc := range []struct {
		num    int
		toBase int
		want   string
	}{
		{num: 0, toBase: 10, want: "0"},
		{num: 123, toBase: 10, want: "123"},
		{num: 74667, toBase: 16, want: "123ab"},
		{num: 45896, toBase: 36, want: "zew"},
		{num: 13, toBase: 2, want: "1101"},
	} {
		if got := fromBase10(tc.num, tc.toBase); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("toBase10(%v, %v) = (%v), want = (%v)", tc.num, tc.toBase, got, tc.want)
		}
	}
}
