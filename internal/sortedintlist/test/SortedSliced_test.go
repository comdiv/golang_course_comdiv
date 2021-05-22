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
