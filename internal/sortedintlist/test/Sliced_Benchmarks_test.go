package test

import (
	"github.com/comdiv/golang_course_comdiv/internal/sortedintlist"
	"github.com/comdiv/golang_course_comdiv/internal/sortedintlist/slices"
	"testing"
)

func BenchmarkSortedSliced_InsertRandom(b *testing.B) {
	GenericBenchmarkSorted_InsertRandom(func() sortedintlist.ISortedIntList {
		return slices.NewSortedIntListSliced()
	}, b)
}

func BenchmarkSortedSliced_InsertAscNoDups(b *testing.B) {
	GenericBenchmarkSorted_InsertAscNoDups(func() sortedintlist.ISortedIntList {
		return slices.NewSortedIntListSliced()
	}, b)
}

func BenchmarkSortedSliced_InsertDescNoDups(b *testing.B) {
	GenericBenchmarkSorted_InsertDescNoDups(func() sortedintlist.ISortedIntList {
		return slices.NewSortedIntListSliced()
	}, b)
}
func BenchmarkSortedSliced_InsertManyDups(b *testing.B) {
	GenericBenchmarkSorted_InsertManyDups(func() sortedintlist.ISortedIntList {
		return slices.NewSortedIntListSliced()
	}, b)
}

func BenchmarkSortedSliced_Delete(b *testing.B) {
	GenericBenchmarkSorted_Delete(func() sortedintlist.ISortedIntList {
		return slices.NewSortedIntListSliced()
	}, b)
}

func BenchmarkSortedSliced_GetAll(b *testing.B) {
	GenericBenchmarkSorted_GetAll(func() sortedintlist.ISortedIntList {
		return slices.NewSortedIntListSliced()
	}, b)
}
func BenchmarkSortedSliced_GetUnique(b *testing.B) {
	GenericBenchmarkSorted_GetAll(func() sortedintlist.ISortedIntList {
		return slices.NewSortedIntListSliced()
	}, b)
}

func GenericLastIndexOfBenchmark(size int, sortedMode bool, b *testing.B) {
	data := make([]int, size)
	for i, _ := range data {
		data[i] = i
	}
	b.ResetTimer()
	var result int
	for i := 0; i < b.N; i++ {
		result, _ = slices.LastIndexOf(data, data[1], sortedMode)
	}
	_ = result
}

func BenchmarkLastIndexOf_5_sorted(b *testing.B) {
	GenericLastIndexOfBenchmark(5, true, b)
}
func BenchmarkLastIndexOf_5_non_sorted(b *testing.B) {
	GenericLastIndexOfBenchmark(5, false, b)
}

func BenchmarkLastIndexOf_10_sorted(b *testing.B) {
	GenericLastIndexOfBenchmark(10, true, b)
}
func BenchmarkLastIndexOf_10_non_sorted(b *testing.B) {
	GenericLastIndexOfBenchmark(10, false, b)
}

func BenchmarkLastIndexOf_20_sorted(b *testing.B) {
	GenericLastIndexOfBenchmark(20, true, b)
}
func BenchmarkLastIndexOf_20_non_sorted(b *testing.B) {
	GenericLastIndexOfBenchmark(20, false, b)
}

func BenchmarkLastIndexOf_100_sorted(b *testing.B) {
	GenericLastIndexOfBenchmark(100, true, b)
}
func BenchmarkLastIndexOf_100_non_sorted(b *testing.B) {
	GenericLastIndexOfBenchmark(100, false, b)
}
func BenchmarkLastIndexOf_1000_sorted(b *testing.B) {
	GenericLastIndexOfBenchmark(1000, true, b)
}
func BenchmarkLastIndexOf_1000_non_sorted(b *testing.B) {
	GenericLastIndexOfBenchmark(1000, false, b)
}

func BenchmarkLastIndexOf_10000_sorted(b *testing.B) {
	GenericLastIndexOfBenchmark(10000, true, b)
}
func BenchmarkLastIndexOf_10000_non_sorted(b *testing.B) {
	GenericLastIndexOfBenchmark(10000, false, b)
}
