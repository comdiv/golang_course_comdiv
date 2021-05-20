package lesson_005

import "testing"

func TestExchangePointer(t *testing.T) {
	var x1, x2 = 2, 4
	exchangePointer(&x1, &x2)
	if x1 != 4 || x2 != 2 {
		t.Errorf("Expected 4, 2 but was %d, %d", x1, x2)
	}
}
