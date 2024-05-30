package service

import "testing"

func TestFactorialSuccess(t *testing.T) {
	wanted := 120
	res := factorial(5)
	if res != wanted {
		t.Errorf("Calc factorial fail: fact %d, want %d", res, wanted)
	}

}
