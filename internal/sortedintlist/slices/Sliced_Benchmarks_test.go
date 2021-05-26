package slices

import (
	"github.com/comdiv/golang_course_comdiv/internal/sortedintlist"
	"github.com/comdiv/golang_course_comdiv/internal/sortedintlistgentest"
	"testing"
)

func BenchmarkSortedSliced_InsertRandom(b *testing.B) {
	sortedintlistgentest.GenericBenchmarkSorted_InsertRandom(func() sortedintlist.IIntInsert {
		return New()
	}, b)
}

func BenchmarkSortedSliced_InsertAscNoDups(b *testing.B) {
	sortedintlistgentest.GenericBenchmarkSorted_InsertAscNoDups(func() sortedintlist.IIntInsert {
		return New()
	}, b)
}

func BenchmarkSortedSliced_InsertDescNoDups(b *testing.B) {
	sortedintlistgentest.GenericBenchmarkSorted_InsertDescNoDups(func() sortedintlist.IIntInsert {
		return New()
	}, b)
}
func BenchmarkSortedSliced_InsertManyDups(b *testing.B) {
	sortedintlistgentest.GenericBenchmarkSorted_InsertManyDups(func() sortedintlist.IIntInsert {
		return New()
	}, b)
}

func BenchmarkSortedSliced_Delete(b *testing.B) {
	sortedintlistgentest.GenericBenchmarkSorted_Delete(func() sortedintlist.IIntCollectionMutable {
		return New()
	}, b)
}

func BenchmarkSortedSliced_GetAll(b *testing.B) {
	sortedintlistgentest.GenericBenchmarkSorted_GetAll(func() sortedintlist.IIntListMutable {
		return New()
	}, b)
}
func BenchmarkSortedSliced_GetUnique(b *testing.B) {
	sortedintlistgentest.GenericBenchmarkSorted_GetUnique(func() sortedintlist.IIntSetMutable {
		return New()
	}, b)
}

func GenericLastIndexOfBenchmark(size int, sortedMode bool, b *testing.B) {
	data := make([]int, size)
	for i := range data {
		data[i] = i
	}
	b.ResetTimer()
	var result int
	for i := 0; i < b.N; i++ {
		result, _ = LastIndexOf(data, data[1], sortedMode)
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
