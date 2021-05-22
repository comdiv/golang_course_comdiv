package test

import (
	"github.com/comdiv/golang_course_comdiv/internal/sortedintlist"
	"github.com/comdiv/golang_course_comdiv/internal/sortedintlist/slices"
	"testing"
)

func BenchmarkSortedSliced_Insert(b *testing.B) {
	GenericBenchmarkSorted_Insert(func() sortedintlist.ISortedIntList {
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
