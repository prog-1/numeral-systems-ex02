package main

import "testing"

func TestToBase10(t *testing.T) {
	for _, tc := range []struct {
		number   string
		fromBase int
		want     int
	}{
		{number: "0", fromBase: 3, want: 0},
		{number: "1", fromBase: 4, want: 1},
		{number: "123", fromBase: 10, want: 123},
		{number: "111101", fromBase: 2, want: 61},
		{number: "ABCD", fromBase: 16, want: 43981},
		{number: "443", fromBase: 5, want: 123},
		{number: "1252", fromBase: 6, want: 320},
		{number: "19C51", fromBase: 14, want: 65535},
		{number: "93B7E", fromBase: 23, want: 2561064},
		{number: "A12CE75XZ90", fromBase: 36, want: 36669757855281588},
	} {
		got := toBase10(tc.number, tc.fromBase)
		if got != tc.want {
			t.Errorf("ERR: toBase10(%s, %v): got: %v, want: %v", tc.number, tc.fromBase, got, tc.want)
		}
	}
}

func TestFromBase10(t *testing.T) {
	for _, tc := range []struct {
		number int
		toBase int
		want   string
	}{
		{number: 0, toBase: 6, want: "0"},
		{number: 1, toBase: 14, want: "1"},
		{number: 123, toBase: 10, want: "123"},
		{number: 61, toBase: 2, want: "111101"},
		{number: 43981, toBase: 16, want: "ABCD"},
		{number: 123, toBase: 5, want: "443"},
		{number: 320, toBase: 16, want: "140"},
		{number: 65535, toBase: 5, want: "4044120"},
		{number: 2561064, toBase: 7, want: "30524442"},
		{number: 36669757855281588, toBase: 20, want: "8J10C869FA3J8"},
	} {
		got := fromBase10(tc.number, tc.toBase)
		if got != tc.want {
			t.Errorf("ERR: fromBase10(%v, %v): got: %s, want: %s", tc.number, tc.toBase, got, tc.want)
		}
	}
}
