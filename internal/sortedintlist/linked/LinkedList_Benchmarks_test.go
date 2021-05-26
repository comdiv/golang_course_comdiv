package linked_test

import (
	"github.com/comdiv/golang_course_comdiv/internal/sortedintlist"
	"github.com/comdiv/golang_course_comdiv/internal/sortedintlist/linked"
	"github.com/comdiv/golang_course_comdiv/internal/sortedintlistgentest"
	"testing"
)

func BenchmarkSortedLinkedList_InsertRandom(b *testing.B) {
	sortedintlistgentest.GenericBenchmarkSorted_InsertRandom(func() sortedintlist.IIntInsert {
		return linked.New()
	}, b)
}

func BenchmarkSortedLinkedList_InsertAscNoDups(b *testing.B) {
	sortedintlistgentest.GenericBenchmarkSorted_InsertAscNoDups(func() sortedintlist.IIntInsert {
		return linked.New()
	}, b)
}

func BenchmarkSortedLinkedList_InsertDescNoDups(b *testing.B) {
	sortedintlistgentest.GenericBenchmarkSorted_InsertDescNoDups(func() sortedintlist.IIntInsert {
		return linked.New()
	}, b)
}
func BenchmarkSortedLinkedList_InsertManyDups(b *testing.B) {
	sortedintlistgentest.GenericBenchmarkSorted_InsertManyDups(func() sortedintlist.IIntInsert {
		return linked.New()
	}, b)
}

func BenchmarkSortedLinkedList_Delete(b *testing.B) {
	sortedintlistgentest.GenericBenchmarkSorted_Delete(func() sortedintlist.IIntCollectionMutable {
		return linked.New()
	}, b)
}

func BenchmarkSortedLinkedList_GetAll(b *testing.B) {
	sortedintlistgentest.GenericBenchmarkSorted_GetAll(func() sortedintlist.IIntListMutable {
		return linked.New()
	}, b)
}
func BenchmarkSortedLinkedList_GetUnique(b *testing.B) {
	sortedintlistgentest.GenericBenchmarkSorted_GetUnique(func() sortedintlist.IIntSetMutable {
		return linked.New()
	}, b)
}

func BenchmarkLinkedFind(b *testing.B) {
	l := linked.New()
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
