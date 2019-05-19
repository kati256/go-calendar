package calendar

import "testing"

func TestCenteringEqual(t *testing.T) {
	res := center("x", 1)
	if res != "x" {
		t.Errorf("Single character didn't center properly:\n$%s$", res)
	}

	res = center("xx", 2)
	if res != "xx" {
		t.Errorf("Didn't center properly:\n$%s$", res)
	}
}

func TestCenteringPair(t *testing.T) {
	res := center("xx", 4)
	if res != " xx " {
		t.Errorf("Pair strings didn't center properly with pair widths:\n$%s$",
			res)
	}
	res = center("xx", 5)
	if res != " xx  " {
		t.Errorf("Pair strings didn't center properly with odd widths:\n$%s$",
			res)
	}
}

func TestCenterintOdd(t *testing.T) {
	res := center("xxx", 4)
	if res != "xxx " {
		t.Errorf("Odd strings didn't center properly with even widths:\n$%s$",
			res)
	}
	res = center("xxx", 5)
	if res != " xxx " {
		t.Errorf("Odd strings didn't center properly with odd widths:\n$%s$",
			res)
	}
}
