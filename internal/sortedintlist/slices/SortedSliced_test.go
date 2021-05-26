package slices

import (
	"github.com/comdiv/golang_course_comdiv/internal/sortedintlist"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortedSliced_GetUnique(t *testing.T) {
	sortedintlist.GenericTestSorted_GetUnique(NewSortedIntListSliced(), t)
}

func TestSortedSliced_GetAll(t *testing.T) {
	l := NewSortedIntListSliced()
	sortedintlist.GenericTestSorted_GetAll(l, t)
}

func TestSortedSliced_Size(t *testing.T) {
	l := NewSortedIntListSliced()
	sortedintlist.GenericTestSorted_Size(l, t)
}

func TestSortedSliced_UniqueSize(t *testing.T) {
	l := NewSortedIntListSliced()
	sortedintlist.GenericTestSorted_UniqueSize(l, t)
}

func TestSortedSliced_InsertList(t *testing.T) {
	l := NewSortedIntListSliced()
	sortedintlist.GenericTestSorted_InsertList(l, t)
}

func TestSortedSliced_InsertSet(t *testing.T) {
	l := NewSortedIntListSliced()
	sortedintlist.GenericTestSorted_InsertSet(l, t)
}

func TestSortedSliced_DeleteList(t *testing.T) {
	l := NewSortedIntListSliced()
	sortedintlist.GenericTestSorted_DeleteList(l, t)
}

func TestSortedSliced_DeleteSet(t *testing.T) {
	l := NewSortedIntListSliced()
	sortedintlist.GenericTestSorted_DeleteSet(l, t)
}

func TestSortedSliced_MinMax(t *testing.T) {
	l := NewSortedIntListSliced()
	sortedintlist.GenericTestSorted_MinMax(l, t)
}

func TestLastIndexOfSortedSlice(t *testing.T) {
	cases := []struct {
		sorted   bool
		data     []int
		value    int
		expected int
		exlast   int
	}{
		{true, nil, 1, LAST_NOT_FOUND, 0},
		{true, []int{}, 1, LAST_NOT_FOUND, 0},
		{true, []int{1, 2, 3, 4}, 3, 2, 2},
		{true, []int{1, 2, 3, 4}, 5, LAST_INDEX_AFTER, 3},
		{true, []int{1, 2, 3, 3, 3, 4}, 5, LAST_INDEX_AFTER, 5},
		{true, []int{2, 3, 4}, 1, LAST_INDEX_BEFORE, 0},
		{true, []int{2, 3, 4, 6, 7}, 5, LAST_NOT_FOUND, 2},
		{true, []int{2, 3, 4, 6, 7, 17, 19, 27, 33, 87, 99, 108, 113, 167, 198, 234, 235, 455}, 113, 12, 12},
		{true, []int{2, 3, 4, 6, 7, 17, 19, 27, 33, 87, 99, 108, 113, 113, 113, 167, 198, 234, 235, 455}, 113, 14, 14},
		{true, []int{2, 3, 4, 6, 7, 17, 19, 27, 33, 87, 99, 108, 113, 167, 198, 234, 235, 455}, 111, LAST_NOT_FOUND, 11},
	}

	for _, c := range cases {
		// сначала проверим всегда без оптимизаций
		result, _ := LastIndexOf(c.data, c.value, false)
		assert.Equal(t, c.expected, result)

		if c.sorted {
			result, last := LastIndexOf(c.data, c.value, true)
			assert.Equal(t, c.expected, result)
			assert.Equal(t, c.exlast, last)
		}
	}
}
