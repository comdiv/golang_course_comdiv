package sortedintlistgentest

import (
	"github.com/comdiv/golang_course_comdiv/internal/sortedintlist"
	"math/rand"
	"testing"
)

const DEFAULT_BENCH_DATA_SIZE int = 10000
const DEFAULT_BENCH_DATA_SEED int64 = 1234567890

func GenericBenchmarkSorted_InsertRandom(create func() sortedintlist.IIntInsert, b *testing.B) {
	for n := 0; n < b.N; n++ {
		list := create()
		for _, v := range defaultRandomSet {
			list.Insert(v)
		}
	}
}

func GenericBenchmarkSorted_InsertAscNoDups(create func() sortedintlist.IIntInsert, b *testing.B) {
	values := generateData(b, func(i int) int {
		return i
	})
	for n := 0; n < b.N; n++ {
		list := create()
		for _, v := range values {
			list.Insert(v)
		}
	}
}

func GenericBenchmarkSorted_InsertDescNoDups(create func() sortedintlist.IIntInsert, b *testing.B) {
	var values = generateData(b, func(i int) int {
		return DEFAULT_BENCH_DATA_SIZE - i
	})
	for n := 0; n < b.N; n++ {
		list := create()
		for _, v := range values {
			list.Insert(v)
		}
	}
}

func GenericBenchmarkSorted_InsertManyDups(create func() sortedintlist.IIntInsert, b *testing.B) {
	values := generateData(b, func(i int) int {
		return i % 20
	})
	for n := 0; n < b.N; n++ {
		list := create()
		for _, v := range values {
			list.Insert(v)
		}
	}
}

func GenericBenchmarkSorted_Delete(create func() sortedintlist.IIntCollectionMutable, b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		list := create()
		for _, v := range defaultRandomSet {
			list.Insert(v)
		}
		b.StartTimer()
		for _, v := range defaultRandomSet {
			list.Delete(v, true)
		}
	}
}

func GenericBenchmarkSorted_GetAll(create func() sortedintlist.IIntListMutable, b *testing.B) *[]int {
	list := create()
	for _, v := range defaultRandomSet {
		list.Insert(v)
	}
	b.ResetTimer()
	var catchResult []int
	for n := 0; n < b.N; n++ {
		catchResult = list.GetAll()
	}
	return &catchResult
}

func GenericBenchmarkSorted_GetUnique(create func() sortedintlist.IIntSetMutable, b *testing.B) *[]int {
	list := create()
	for _, v := range defaultRandomSet {
		list.Insert(v)
	}
	b.ResetTimer()
	var catchResult []int
	for n := 0; n < b.N; n++ {
		catchResult = list.GetUnique()
	}
	return &catchResult
}

// generateDataSRSeed - обобщенны генератор с рандомизатором
func generateDataSRSeed(b *testing.B, size int, generator func(i int, r *rand.Rand) int, seed int64) []int {
	defer func() {
		if nil != b {
			b.ResetTimer()
		}
	}()
	var randomizer = rand.New(rand.NewSource(seed))
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = generator(i, randomizer)
	}
	return data
}

/**
Различные перегруженные варианты генератора для более простых вариантов использования
*/

func generateData(b *testing.B, generator func(i int) int) []int {
	return generateDataSRSeed(b, DEFAULT_BENCH_DATA_SIZE, func(i int, _ *rand.Rand) int { return generator(i) }, DEFAULT_BENCH_DATA_SEED)
}

func generateDataS(b *testing.B, size int, generator func(i int) int) []int {
	return generateDataSRSeed(b, size, func(i int, _ *rand.Rand) int { return generator(i) }, DEFAULT_BENCH_DATA_SEED)
}

func generateDataR(b *testing.B, generator func(i int, r *rand.Rand) int) []int {
	return generateDataSRSeed(b, DEFAULT_BENCH_DATA_SIZE, generator, 1234567890)
}
func generateDataSR(b *testing.B, size int, generator func(i int, r *rand.Rand) int) []int {
	return generateDataSRSeed(b, size, generator, DEFAULT_BENCH_DATA_SEED)
}

var defaultRandomSet = generateDataR(nil, func(i int, r *rand.Rand) int {
	return r.Intn(DEFAULT_BENCH_DATA_SIZE / 2)
})
