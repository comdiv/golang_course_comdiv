package test

import (
	"github.com/comdiv/golang_course_comdiv/internal/sortedintlist"
	"reflect"
	"testing"
)

func GenericTestSorted_GetUnique(l sortedintlist.ISortedIntList, t *testing.T) {
	sortedintlist.InsertAllVar(l, 8, 1, 2, 4, 5, 4, 4, 5, 6, 1)
	all := l.GetUnique()
	expected := []int{1, 2, 4, 5, 6, 8}
	if !reflect.DeepEqual(all, expected) {
		t.Errorf("Вернулись не те значения `%v`", all)
	}
}

func GenericTestSorted_GetAll(l sortedintlist.ISortedIntList, t *testing.T) {
	sortedintlist.InsertAllVar(l, 8, 1, 2, 4, 5, 4, 4, 5, 6, 1)
	all := l.GetAll()
	expected := []int{1, 1, 2, 4, 4, 4, 5, 5, 6, 8}
	if !reflect.DeepEqual(all, expected) {
		t.Errorf("Вернулись не те значения `%v`", all)
	}
}

func GenericTestSorted_Size(l sortedintlist.ISortedIntList, t *testing.T) {
	sortedintlist.InsertAllVar(l, 1, 2, 4, 4, 4)
	if l.Size() != 5 {
		t.Errorf("Expected 5 but was %d", l.Size())
	}
}

func GenericTestSorted_UniqueSize(l sortedintlist.ISortedIntList, t *testing.T) {
	sortedintlist.InsertAllVar(l, 1, 2, 4, 4, 4)
	if l.UniqueSize() != 3 {
		t.Errorf("Expected 3 but was %d", l.UniqueSize())
	}
}

func GenericTestSorted_Insert(l sortedintlist.ISortedIntList, t *testing.T) {
	var inserted bool
	inserted = l.Insert(1)
	if !(inserted && l.UniqueSize() == 1 && l.Size() == 1) {
		t.Errorf("%v %v %v", inserted, l.UniqueSize(), l.Size())
	}
	inserted = l.Insert(10)
	if !(inserted && l.UniqueSize() == 2 && l.Size() == 2) {
		t.Errorf("%v %v %v", inserted, l.UniqueSize(), l.Size())
	}

	inserted = l.Insert(10)
	if !(!inserted && l.UniqueSize() == 2 && l.Size() == 3) {
		t.Errorf("%v %v %v", inserted, l.UniqueSize(), l.Size())
	}
}

func GenericTestSorted_Delete(l sortedintlist.ISortedIntList, t *testing.T) {
	l.Insert(1)
	l.Insert(10)
	l.Insert(11)
	l.Insert(12)
	l.Insert(12)
	l.Insert(12)
	if !(l.UniqueSize() == 4 && l.Size() == 6) {
		t.Errorf("%v %v", l.UniqueSize(), l.Size())
	}
	var deleted bool
	deleted = l.Delete(10, true)
	if !(deleted && l.UniqueSize() == 3 && l.Size() == 5) {
		t.Errorf("%v %v %v", deleted, l.UniqueSize(), l.Size())
	}
	deleted = l.Delete(77777, true)
	if !(!deleted && l.UniqueSize() == 3 && l.Size() == 5) {
		t.Errorf("%v %v %v", deleted, l.UniqueSize(), l.Size())
	}
	deleted = l.Delete(12, false)
	if !(deleted && l.UniqueSize() == 3 && l.Size() == 4) {
		t.Errorf("%v %v %v", deleted, l.UniqueSize(), l.Size())
	}
	deleted = l.Delete(12, true)
	if !(deleted && l.UniqueSize() == 2 && l.Size() == 2) {
		t.Errorf("%v %v %v", deleted, l.UniqueSize(), l.Size())
	}
}
