package main

import (
	"testing"
)

func TestFromBase10(t *testing.T) {
	for _, tc := range []struct {
		num    int
		toBase int
		want   string
	}{
		{0, 0, "-1"},
		{0, 1, "-1"},
		{123, 10, "123"},
		{123, 5, "443"},
		{123, 0, "-1"},
		{123, 1, "-1"},
		{123, 37, "-1"},
		{123, -1, "-1"},
		{320, 10, "320"},
		{320, 16, "140"},
		{320, -1, "-1"},
		{320, -320, "-1"},
		{61, 2, "111101"},
		{61, 4, "331"},
		{43981, 16, "ABCD"},
		{84653, 20, "ABCD"},
		{65535, 5, "4044120"},
	} {
		t.Run(tc.want, func(t *testing.T) {
			if got := fromBase10(tc.num, tc.toBase); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
