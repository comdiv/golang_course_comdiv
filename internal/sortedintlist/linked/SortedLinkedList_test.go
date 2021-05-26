package linked_test

import (
	"github.com/comdiv/golang_course_comdiv/internal/sortedintlist"
	"github.com/comdiv/golang_course_comdiv/internal/sortedintlist/linked"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestSortedLinkedList_GetUnique(t *testing.T) {
	sortedintlist.GenericTestSorted_GetUnique(linked.NewSortedLinkedList(), t)
}

func TestSortedLinkedList_GetAll(t *testing.T) {
	l := linked.NewSortedLinkedList()
	sortedintlist.GenericTestSorted_GetAll(l, t)
}

func TestSortedLinkedList_Size(t *testing.T) {
	l := linked.NewSortedLinkedList()
	sortedintlist.GenericTestSorted_Size(l, t)
}

func TestSortedLinkedList_UniqueSize(t *testing.T) {
	l := linked.NewSortedLinkedList()
	sortedintlist.GenericTestSorted_UniqueSize(l, t)
}

func TestSortedLinkedList_InsertList(t *testing.T) {
	l := linked.NewSortedLinkedList()
	sortedintlist.GenericTestSorted_InsertList(l, t)
}

func TestSortedLinkedList_InsertSet(t *testing.T) {
	l := linked.NewSortedLinkedList()
	sortedintlist.GenericTestSorted_InsertSet(l, t)
}

func TestSortedLinkedList_DeleteList(t *testing.T) {
	l := linked.NewSortedLinkedList()
	sortedintlist.GenericTestSorted_DeleteList(l, t)
}

func TestSortedLinkedList_DeleteSet(t *testing.T) {
	l := linked.NewSortedLinkedList()
	sortedintlist.GenericTestSorted_DeleteSet(l, t)
}

func TestSortedLinkedList_MinMax(t *testing.T) {
	l := linked.NewSortedLinkedList()
	sortedintlist.GenericTestSorted_MinMax(l, t)
}

func TestSortedLinkedList_FindItemFor(t *testing.T) {
	l := linked.NewSortedLinkedList()
	var items = [3]int{5, 10, 100}
	l.Insert(items[0])
	l.Insert(items[1])
	l.Insert(items[2])

	// можем все найти прямо
	for _, x := range items {
		trg, tp := l.FindItemFor(x)
		assert.Equal(t, x, trg.Value())
		assert.Equal(t, linked.FOUND_TYPE_FOUND, tp)
	}
	trg, tp := l.FindItemFor(items[0] - 100)
	assert.Equal(t, items[0], trg.Value())
	assert.Equal(t, linked.FOUND_TYPE_NEXT, tp)

	trg, tp = l.FindItemFor(items[2] + 100)
	assert.Equal(t, items[2], trg.Value())
	assert.Equal(t, linked.FOUND_TYPE_PREV, tp)

	trg, tp = l.FindItemFor(items[1] + 10)
	assert.Equal(t, items[1], trg.Value())
	assert.Equal(t, linked.FOUND_TYPE_PREV, tp)

}

func TestSortedLinkedList_Tail(t *testing.T) {
	var randomizer = rand.New(rand.NewSource(123445455))
	val := randomizer.Intn(100)
	l := linked.NewSortedLinkedList()
	l.Insert(val)
	assert.NotNil(t, l.Tail())

	initalTail := l.Tail()
	assert.Equal(t, val, l.Tail().Value())

	l.Insert(val + 1)
	assert.Equal(t, val+1, l.Tail().Value())
	assert.Equal(t, l.Tail(), initalTail.Next())
	assert.Equal(t, initalTail, l.Tail().Prev())

	l.Delete(val+1, true)
	assert.Equal(t, initalTail, l.Tail())
	assert.Nil(t, initalTail.Next())
}
func TestSortedLinkedList_Head(t *testing.T) {
	var randomizer = rand.New(rand.NewSource(123445455))
	val := randomizer.Intn(100)
	l := linked.NewSortedLinkedList()
	l.Insert(val)
	assert.NotNil(t, l.Head())

	initialHead := l.Head()
	assert.Equal(t, val, l.Head().Value())

	l.Insert(val - 1)
	assert.Equal(t, val-1, l.Head().Value())
	assert.Equal(t, l.Head(), initialHead.Prev())

	l.Delete(val-1, true)
	assert.Equal(t, initialHead, l.Head())
	assert.Nil(t, initialHead.Prev())
}

func TestNewSortedLinkedList(t *testing.T) {
	var l = linked.NewSortedLinkedList()
	assert.Equal(t, 0, l.Size())
	assert.Equal(t, 0, l.UniqueSize())
	assert.Nil(t, l.Head())
	assert.Nil(t, l.Tail())
	var l2 = linked.NewSortedLinkedList()
	assert.NotSame(t, l, l2)
}

func TestSortedLinkedListItem_Count(t *testing.T) {
	var randomizer = rand.New(rand.NewSource(123445455))
	expected := randomizer.Intn(100)
	item := linked.NewSortedLinkedListItemC(1, expected)
	assert.Equal(t, expected, item.Count())
}

func TestSortedLinkedListItem_Value(t *testing.T) {
	var randomizer = rand.New(rand.NewSource(123445455))
	expected := randomizer.Intn(100)
	item := linked.NewSortedLinkedListItem(expected)
	assert.Equal(t, expected, item.Value())
}
