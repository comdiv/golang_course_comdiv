package sortedintlist

import (
	"math/rand"
	"testing"
)

func GenericBenchmarkSorted_InsertRandom(create func() ISortedIntList, b *testing.B) {
	var randomizer = rand.New(rand.NewSource(1234567890))
	var values [10000]int
	for i, _ := range values {
		values[i] = randomizer.Intn(5000)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		list := create()
		for _, v := range values {
			list.Insert(v)
		}
	}
}

func GenericBenchmarkSorted_InsertAscNoDups(create func() ISortedIntList, b *testing.B) {
	var values [10000]int
	for i, _ := range values {
		values[i] = i
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		list := create()
		for _, v := range values {
			list.Insert(v)
		}
	}
}

func GenericBenchmarkSorted_InsertDescNoDups(create func() ISortedIntList, b *testing.B) {
	var values [10000]int
	for i, _ := range values {
		values[i] = 10000 - i
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		list := create()
		for _, v := range values {
			list.Insert(v)
		}
	}
}

func GenericBenchmarkSorted_InsertManyDups(create func() ISortedIntList, b *testing.B) {
	var values [10000]int
	for i, _ := range values {
		values[i] = i % 20
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		list := create()
		for _, v := range values {
			list.Insert(v)
		}
	}
}

func GenericBenchmarkSorted_Delete(create func() ISortedIntList, b *testing.B) {
	var randomizer = rand.New(rand.NewSource(1234567890))
	var values [10000]int
	for i, _ := range values {
		values[i] = randomizer.Intn(5000)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		list := create()
		for _, v := range values {
			list.Insert(v)
		}
		b.StartTimer()
		for _, v := range values {
			list.Delete(v, true)
		}
	}
}

func GenericBenchmarkSorted_GetAll(create func() ISortedIntList, b *testing.B) *[]int {
	var randomizer = rand.New(rand.NewSource(1234567890))
	var values [10000]int
	for i, _ := range values {
		values[i] = randomizer.Intn(5000)
	}
	list := create()
	for _, v := range values {
		list.Insert(v)
	}
	b.ResetTimer()
	var catchResult []int
	for n := 0; n < b.N; n++ {
		catchResult = list.GetAll()
	}
	return &catchResult
}

func GenericBenchmarkSorted_GetUnique(create func() ISortedIntList, b *testing.B) *[]int {
	var randomizer = rand.New(rand.NewSource(1234567890))
	var values [10000]int
	for i, _ := range values {
		values[i] = randomizer.Intn(5000)
	}
	list := create()
	for _, v := range values {
		list.Insert(v)
	}
	b.ResetTimer()
	var catchResult []int
	for n := 0; n < b.N; n++ {
		catchResult = list.GetUnique()
	}
	return &catchResult
}
