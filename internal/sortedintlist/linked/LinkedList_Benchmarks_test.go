package linked_test

import (
	"github.com/comdiv/golang_course_comdiv/internal/sortedintlist"
	"github.com/comdiv/golang_course_comdiv/internal/sortedintlist/linked"
	"testing"
)

func BenchmarkSortedLinkedList_InsertRandom(b *testing.B) {
	sortedintlist.GenericBenchmarkSorted_InsertRandom(func() sortedintlist.IIntInsert {
		return linked.NewSortedLinkedList()
	}, b)
}

func BenchmarkSortedLinkedList_InsertAscNoDups(b *testing.B) {
	sortedintlist.GenericBenchmarkSorted_InsertAscNoDups(func() sortedintlist.IIntInsert {
		return linked.NewSortedLinkedList()
	}, b)
}

func BenchmarkSortedLinkedList_InsertDescNoDups(b *testing.B) {
	sortedintlist.GenericBenchmarkSorted_InsertDescNoDups(func() sortedintlist.IIntInsert {
		return linked.NewSortedLinkedList()
	}, b)
}
func BenchmarkSortedLinkedList_InsertManyDups(b *testing.B) {
	sortedintlist.GenericBenchmarkSorted_InsertManyDups(func() sortedintlist.IIntInsert {
		return linked.NewSortedLinkedList()
	}, b)
}

func BenchmarkSortedLinkedList_Delete(b *testing.B) {
	sortedintlist.GenericBenchmarkSorted_Delete(func() sortedintlist.IIntCollectionMutable {
		return linked.NewSortedLinkedList()
	}, b)
}

func BenchmarkSortedLinkedList_GetAll(b *testing.B) {
	sortedintlist.GenericBenchmarkSorted_GetAll(func() sortedintlist.IIntListMutable {
		return linked.NewSortedLinkedList()
	}, b)
}
func BenchmarkSortedLinkedList_GetUnique(b *testing.B) {
	sortedintlist.GenericBenchmarkSorted_GetUnique(func() sortedintlist.IIntSetMutable {
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
