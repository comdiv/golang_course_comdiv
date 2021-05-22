package test

import (
	"github.com/comdiv/golang_course_comdiv/internal/sortedintlist"
	"github.com/comdiv/golang_course_comdiv/internal/sortedintlist/linked"
	"testing"
)

func BenchmarkSortedLinkedList_Insert(b *testing.B) {
	GenericBenchmarkSorted_Insert(func() sortedintlist.ISortedIntList {
		return linked.NewSortedLinkedList()
	}, b)
}
func BenchmarkSortedLinkedList_Delete(b *testing.B) {
	GenericBenchmarkSorted_Delete(func() sortedintlist.ISortedIntList {
		return linked.NewSortedLinkedList()
	}, b)
}

func BenchmarkSortedLinkedList_GetAll(b *testing.B) {
	GenericBenchmarkSorted_GetAll(func() sortedintlist.ISortedIntList {
		return linked.NewSortedLinkedList()
	}, b)
}
func BenchmarkSortedLinkedList_GetUnique(b *testing.B) {
	GenericBenchmarkSorted_GetAll(func() sortedintlist.ISortedIntList {
		return linked.NewSortedLinkedList()
	}, b)
}

func BenchmarkLinkedFind(b *testing.B) {
	l := linked.NewSortedLinkedList()
	for i := 0; i < 10000; i++ {
		l.Insert(i)
	}
	b.ResetTimer()
	var result *linked.SortedLinkedListItem
	for i := 0; i < b.N; i++ {
		result, _ = l.FindItemFor(9999)
	}
	_ = result
}
