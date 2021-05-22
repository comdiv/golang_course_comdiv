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
