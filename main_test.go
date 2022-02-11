package main

import (
	"reflect"
	"testing"
)

// I have  a lot of questions about this testing structure (maybe i can use it only on playground???)
/*func testCaseError(t *testing.T) {// here how i understand coded what to do when is error
	t.Error("forced error")
}

func testCaseOK(t *testing.T) {
	for _, tc := range []struct {	// here write what? variables of tested function or want
		name string
	}{// this is for tests to write} {// here i understand, itts similar to older tests
		t.Run(name, func(t *testing.T) {
			gotL, gotC := SomeFuncReturningLaureatesAndCounters(tc.awards)
			sort.Strings(gotL)
			if reflect.DeepEqual(gotL, tc.wantL) || gotC != tc.wantC {
				t.Errorf("got = (%v, %v), want = (%v, %v)", gotL, gotC, tc.wantL, tc.wantC)
			}
		})
	}
}

func main() {	// here is error about main redeclared?
	testing.Main(
		/* matchString // func(a, b string) (bool, error) { return a == b, nil }, // what this means?
		/* tests // []testing.InternalTest{
			{Name: "test error", F: testCaseError},
			{Name: "test OK", F: testCaseOK},
		},
		// benchmarks // nil /* examples //,nil)// and what this?
//}*/

func TestToBase10(t *testing.T) {
	for _, tc := range []struct {
		num      string
		fromBase int
		want     int
	}{
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
		{num: 123, toBase: 10, want: "123"},
		{num: 74667, toBase: 16, want: "123AB"},
		{num: 45896, toBase: 36, want: "ZEW"},
		{num: 13, toBase: 2, want: "1101"},
	} {
		if got := fromBase10(tc.num, tc.toBase); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("toBase10(%v, %v) = (%v), want = (%v)", tc.num, tc.toBase, got, tc.want)
		}
	}
}
