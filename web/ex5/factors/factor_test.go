package factors_test

import (
	"."
	"testing"
)

func Test7(t *testing.T) {
	list := factors.Factor(7)
	if len(list) != 0 {
		t.Fatal("7 should have no factors")
	}
}

func Test9(t *testing.T) {
	list := factors.Factor(9)
	if len(list) != 1 {
		t.Fatal("9 should have one factor")
	}
	if list[0] != 3 {
		t.Fatal("9 should have one factor: 3")
	}
}

func Test11(t *testing.T) {
	list := factors.Factor(10)
	if len(list) != 0 {
		t.Fatal("11 should have no factors")
	}
}
