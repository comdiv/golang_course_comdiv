package test

import (
	"github.com/comdiv/golang_course_comdiv/internal/sortedintlist/slices"
	"testing"
)

func TestSortedSliced_GetUnique(t *testing.T) {
	GenericTestSorted_GetUnique(slices.NewSortedIntListSliced(), t)
}

func TestSortedSliced_GetAll(t *testing.T) {
	l := slices.NewSortedIntListSliced()
	GenericTestSorted_GetAll(l, t)
}

func TestSortedSliced_Size(t *testing.T) {
	l := slices.NewSortedIntListSliced()
	GenericTestSorted_Size(l, t)
}

func TestSortedSliced_UniqueSize(t *testing.T) {
	l := slices.NewSortedIntListSliced()
	GenericTestSorted_UniqueSize(l, t)
}

func TestSortedSliced_Insert(t *testing.T) {
	l := slices.NewSortedIntListSliced()
	GenericTestSorted_Insert(l, t)
}

func TestSortedSliced_Delete(t *testing.T) {
	l := slices.NewSortedIntListSliced()
	GenericTestSorted_Delete(l, t)
}

func TestLastIndexOfSortedSlice(t *testing.T) {
	cases := []struct {
		sorted   bool
		data     []int
		value    int
		expected int
		exlast   int
	}{
		{true, nil, 1, slices.LAST_NOT_FOUND, 0},
		{true, []int{}, 1, slices.LAST_NOT_FOUND, 0},
		{true, []int{1, 2, 3, 4}, 3, 2, 2},
		{true, []int{1, 2, 3, 4}, 5, slices.LAST_INDEX_AFTER, 3},
		{true, []int{1, 2, 3, 3, 3, 4}, 5, slices.LAST_INDEX_AFTER, 5},
		{true, []int{2, 3, 4}, 1, slices.LAST_INDEX_BEFORE, 0},
		{true, []int{2, 3, 4, 6, 7}, 5, slices.LAST_NOT_FOUND, 2},
		{true, []int{2, 3, 4, 6, 7, 17, 19, 27, 33, 87, 99, 108, 113, 167, 198, 234, 235, 455}, 113, 12, 12},
		{true, []int{2, 3, 4, 6, 7, 17, 19, 27, 33, 87, 99, 108, 113, 113, 113, 167, 198, 234, 235, 455}, 113, 14, 14},
		{true, []int{2, 3, 4, 6, 7, 17, 19, 27, 33, 87, 99, 108, 113, 167, 198, 234, 235, 455}, 111, slices.LAST_NOT_FOUND, 11},
	}

	for _, c := range cases {
		// сначала проверим всегда без оптимизаций
		result, _ := slices.LastIndexOf(c.data, c.value, false)
		if result != c.expected {
			t.Errorf("In `%v` %d should have index %d but was %d", c.data, c.value, c.expected, result)
		}

		if c.sorted {
			result, last := slices.LastIndexOf(c.data, c.value, true)
			if result != c.expected {
				t.Errorf("In sorted `%v` %d should have index %d but was %d", c.data, c.value, c.expected, result)
			}
			if last != c.exlast {
				t.Errorf("In `%v` %d should have last index %d but was %d", c.data, c.value, c.exlast, last)
			}
		}
	}
}
