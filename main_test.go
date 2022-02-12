package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestToBase10(t *testing.T) {
	for _, tc := range []struct {
		name       string
		number     string
		base       int
		wantNumber int
	}{
		{"empty", "", 16, 0},
		{"from base 10 to base 10", "123", 10, 123},
		{"base 2 to base 10", "111101", 2, 61},
		{"base 16 to base 10", "ABCD", 16, 43981},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := toBase10(tc.number, tc.base); got != tc.wantNumber {
				t.Errorf("toBase10(num = %v, fromBase = %v) = %v, want %v", tc.number, tc.base, got, tc.wantNumber)
			}
		})
	}
}

func TestFromBase10(t *testing.T) {
	for _, tc := range []struct {
		name       string
		number     int
		base       int
		wantNumber string
	}{
		{"zero", 0, 16, "0"},
		{"from base 10 to base 10", 123, 10, "123"},
		{"from base 10 to base 2", 61, 2, "111101"},
		{"from base 10 to base 16", 43981, 16, "ABCD"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := fromBase10(tc.number, tc.base); got != tc.wantNumber {
				t.Errorf("fromBase10(num = %v, toBase = %v) = %v, want %v", tc.number, tc.base, got, tc.wantNumber)
			}
		})
	}
}

func TestBothRandom(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 10; i >= 0; i-- {
		base := rand.Intn(35) + 2
		num := rand.Intn(1000000)
		t.Run("Test using random numbers", func(t *testing.T) {
			func1res := fromBase10(num, base)
			func2res := toBase10(func1res, base)
			if func2res != num {
				t.Errorf(`
fromBase10(num = %v, toBase = %v) returned %v, then
toBase10(num = %v, fromBase = %v) returned %v
%v!=%v
`, num, base, func1res, func1res, base, func2res, num, func2res)
			}
		})
	}
}
