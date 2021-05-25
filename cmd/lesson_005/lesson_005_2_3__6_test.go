package lesson_005

import (
	"math/rand"
	"testing"
)

type exchangeCase struct {
	x1 int
	x2 int
}

func TestExchangePointer(t *testing.T) {
	var randomizer = rand.New(rand.NewSource(655646))
	exchangeCases := []*exchangeCase{
		{1, 2},
		{3, 4},
		{5, 6},
		{randomizer.Intn(100), randomizer.Intn(200)},
		{randomizer.Intn(100), randomizer.Intn(200)},
		{randomizer.Intn(100), randomizer.Intn(200)},
		{randomizer.Intn(100), randomizer.Intn(200)},
	}
	for _, et := range exchangeCases {
		var x1, x2 = et.x1, et.x2
		exchangePointer(&x1, &x2)
		if x1 != et.x2 || x2 != et.x1 {
			t.Errorf("Expected %d, %d but was %d, %d", et.x2, et.x1, x1, x2)
		}
	}
}
