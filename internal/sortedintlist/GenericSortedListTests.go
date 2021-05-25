package sortedintlist

import (
	"reflect"
	"testing"
)

func GenericTestSorted_GetUnique(l IIntSetMutable, t *testing.T) {
	InsertAllVar(l, 8, 1, 2, 4, 5, 4, 4, 5, 6, 1)
	all := l.GetUnique()
	expected := []int{1, 2, 4, 5, 6, 8}
	if !reflect.DeepEqual(all, expected) {
		t.Errorf("Вернулись не те значения `%v` вместо `%v`", all, expected)
	}
}

func GenericTestSorted_GetAll(l IIntListMutable, t *testing.T) {
	InsertAllVar(l, 8, 1, 2, 4, 5, 4, 4, 5, 6, 1)
	all := l.GetAll()
	expected := []int{1, 1, 2, 4, 4, 4, 5, 5, 6, 8}
	if !reflect.DeepEqual(all, expected) {
		t.Errorf("Вернулись не те значения `%v` вместо `%v`", all, expected)
	}
}

func GenericTestSorted_Size(l IIntListMutable, t *testing.T) {
	InsertAllVar(l, 1, 2, 4, 4, 4)
	if l.Size() != 5 {
		t.Errorf("Expected 5 but was %d", l.Size())
	}
}

func GenericTestSorted_UniqueSize(l IIntSetMutable, t *testing.T) {
	InsertAllVar(l, 1, 2, 4, 4, 4)
	if l.UniqueSize() != 3 {
		t.Errorf("Expected 3 but was %d", l.UniqueSize())
	}
}

func GenericTestSorted_InsertList(l IIntListMutable, t *testing.T) {
	var inserted bool
	inserted = l.Insert(1)
	if !(inserted && l.Size() == 1) {
		t.Errorf("%v %v", inserted, l.Size())
	}
	inserted = l.Insert(10)
	if !(inserted && l.Size() == 2) {
		t.Errorf("%v %v", inserted, l.Size())
	}

	inserted = l.Insert(10)
	if !(!inserted && l.Size() == 3) {
		t.Errorf("%v %v", inserted, l.Size())
	}
}

func GenericTestSorted_InsertSet(l IIntSetMutable, t *testing.T) {
	var inserted bool
	inserted = l.Insert(1)
	if !(inserted && l.UniqueSize() == 1) {
		t.Errorf("%v %v", inserted, l.UniqueSize())
	}
	inserted = l.Insert(10)
	if !(inserted && l.UniqueSize() == 2) {
		t.Errorf("%v %v", inserted, l.UniqueSize())
	}

	inserted = l.Insert(10)
	if !(!inserted && l.UniqueSize() == 2) {
		t.Errorf("%v %v", inserted, l.UniqueSize())
	}
}

func GenericTestSorted_DeleteList(l IIntListMutable, t *testing.T) {
	l.Insert(1)
	l.Insert(10)
	l.Insert(11)
	l.Insert(12)
	l.Insert(12)
	l.Insert(12)
	if !(l.Size() == 6) {
		t.Errorf("%v", l.Size())
	}
	var deleted bool
	deleted = l.Delete(10, true)
	if !(deleted && l.Size() == 5) {
		t.Errorf("%v %v", deleted, l.Size())
	}
	deleted = l.Delete(77777, true)
	if !(!deleted && l.Size() == 5) {
		t.Errorf("%v %v", deleted, l.Size())
	}
	deleted = l.Delete(12, false)
	if !(deleted && l.Size() == 4) {
		t.Errorf("%v %v", deleted, l.Size())
	}
	deleted = l.Delete(12, true)
	if !(deleted && l.Size() == 2) {
		t.Errorf("%v %v", deleted, l.Size())
	}
}

func GenericTestSorted_DeleteSet(l IIntSetMutable, t *testing.T) {
	l.Insert(1)
	l.Insert(10)
	l.Insert(11)
	l.Insert(12)
	l.Insert(12)
	l.Insert(12)
	if !(l.UniqueSize() == 4) {
		t.Errorf("%v", l.UniqueSize())
	}
	var deleted bool
	deleted = l.Delete(10, true)
	if !(deleted && l.UniqueSize() == 3) {
		t.Errorf("%v %v", deleted, l.UniqueSize())
	}
	deleted = l.Delete(77777, true)
	if !(!deleted && l.UniqueSize() == 3) {
		t.Errorf("%v %v", deleted, l.UniqueSize())
	}
	deleted = l.Delete(12, false)
	if !(deleted && l.UniqueSize() == 3) {
		t.Errorf("%v %v", deleted, l.UniqueSize())
	}
	deleted = l.Delete(12, true)
	if !(deleted && l.UniqueSize() == 2) {
		t.Errorf("%v %v", deleted, l.UniqueSize())
	}
}
