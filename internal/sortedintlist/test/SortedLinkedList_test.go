package test

import (
	"github.com/comdiv/golang_course_comdiv/internal/sortedintlist/linked"
	"math/rand"
	"reflect"
	"testing"
)

func TestSortedLinkedList_GetDistinct(t *testing.T) {
	l := linked.NewSortedLinkedList()
	l.InsertAllVar(8, 1, 2, 4, 5, 4, 4, 5, 6, 1)
	all := l.GetUnique()
	expected := []int{1, 2, 4, 5, 6, 8}
	if !reflect.DeepEqual(all, expected) {
		t.Errorf("Вернулись не те значения `%v`", all)
	}
}

func TestSortedLinkedList_GetAll(t *testing.T) {
	l := linked.NewSortedLinkedList()
	l.InsertAllVar(8, 1, 2, 4, 5, 4, 4, 5, 6, 1)
	all := l.GetAll()
	expected := []int{1, 1, 2, 4, 4, 4, 5, 5, 6, 8}
	if !reflect.DeepEqual(all, expected) {
		t.Errorf("Вернулись не те значения `%v`", all)
	}
}

func TestSortedLinkedList_ItemCount(t *testing.T) {
	l := linked.NewSortedLinkedList()
	l.InsertAllVar(1, 2, 4, 4, 4)
	if l.Size() != 5 {
		t.Errorf("Expected 5 but was %d", l.Size())
	}
}

func TestSortedLinkedList_IndexSize(t *testing.T) {
	l := linked.NewSortedLinkedList()
	l.InsertAllVar(1, 2, 4, 4, 4)
	if l.UniqueSize() != 3 {
		t.Errorf("Expected 3 but was %d", l.UniqueSize())
	}
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
		if !(trg.Value() == x && tp == linked.FOUND_TYPE_FOUND) {
			t.Errorf("%v %v", trg, tp)
		}
	}
	trg, tp := l.FindItemFor(items[0] - 100)
	if !(trg.Value() == items[0] && tp == linked.FOUND_TYPE_NEXT) {
		t.Errorf("%v %v", trg, tp)
	}
	trg, tp = l.FindItemFor(items[2] + 100)
	if !(trg.Value() == items[2] && tp == linked.FOUND_TYPE_PREV) {
		t.Errorf("%v %v", trg, tp)
	}
	trg, tp = l.FindItemFor(items[1] + 10)
	if !(trg.Value() == items[1] && tp == linked.FOUND_TYPE_PREV) {
		t.Errorf("%v %v", trg, tp)
	}
}

func TestSortedLinkedList_Insert(t *testing.T) {
	l := linked.NewSortedLinkedList()
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

func TestSortedLinkedList_Delete(t *testing.T) {
	l := linked.NewSortedLinkedList()
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

func TestSortedLinkedList_Tail(t *testing.T) {
	val := rand.Intn(100)
	l := linked.NewSortedLinkedList()
	l.Insert(val)
	if l.Tail() == nil {
		t.Errorf("При вставке первого элемента Tail не должен был остаться nil!")
	}
	initalTail := l.Tail()
	if l.Tail().Value() != val {
		t.Errorf("Единственный элемент он же последний и значение ожидалось %d, а на деле %d", val, l.Tail().Value())
	}
	l.Insert(val + 1)
	if l.Tail().Value() != val+1 {
		t.Errorf("Мы добавили значение в хвост и Tail должен был измениться на %d, а на деле %d", val+1, l.Tail().Value())
	}
	if initalTail.Next() != l.Tail() || l.Tail().Prev() != initalTail {
		t.Errorf("При добавлении в хвост связи между элементами не были сформированы")
	}
	l.Delete(val+1, true)
	if l.Tail() != initalTail {
		t.Errorf("После удаления последнего элемента Tail должен был откатиться на элемент назад")
	}
	if initalTail.Next() != nil {
		t.Errorf("После удаления последнего элемента Tail не может ссылаться на следующий элемент")
	}
}
func TestSortedLinkedList_Head(t *testing.T) {
	val := rand.Intn(100)
	l := linked.NewSortedLinkedList()
	l.Insert(val)
	if l.Head() == nil {
		t.Errorf("При вставке первого элемента Tail не должен был остаться nil!")
	}
	initialHead := l.Head()
	if l.Head().Value() != val {
		t.Errorf("Единственный элемент он же последний и значение ожидалось %d, а на деле %d", val, l.Head().Value())
	}
	l.Insert(val - 1)
	if l.Head().Value() != val-1 {
		t.Errorf("Мы добавили значение в начало и Head должен был измениться на %d, а на деле %d", val-1, l.Head().Value())
	}
	if initialHead.Prev() != l.Head() || l.Head().Next() != initialHead {
		t.Errorf("При добавлении в начало связи между элементами не были сформированы")
	}
	l.Delete(val-1, true)
	if l.Head() != initialHead {
		t.Errorf("После удаления первого элемента Head должен был откатиться на элемент назад")
	}
	if initialHead.Prev() != nil {
		t.Errorf("После удаления первого элемента Head не может ссылаться на следующий элемент")
	}
}

func TestNewSortedLinkedList(t *testing.T) {
	var list_1 *linked.SortedLinkedList = linked.NewSortedLinkedList()

	if list_1.Size() != 0 {
		t.Errorf("Size должен быть 0, а он %d", list_1.Size())
	}
	if list_1.UniqueSize() != 0 {
		t.Errorf("UniqueSize должен быть 0, а он %d", list_1.UniqueSize())
	}

	if list_1.Head() != nil {
		t.Errorf("Head должен быть nil, а он %p", list_1.Head())
	}

	if list_1.Tail() != nil {
		t.Errorf("Tail должен быть nil, а он %p", list_1.Tail())
	}
	var list_2 *linked.SortedLinkedList = linked.NewSortedLinkedList()
	if list_1 == list_2 {
		t.Errorf("NewSortedLinkedList shoud generate distinct, not singleton lists")
	}
}

func TestSortedLinkedListItem_Count(t *testing.T) {
	expected := rand.Intn(100)
	item := linked.NewSortedLinkedListItemC(1, expected)
	if item.Count() != expected {
		t.Errorf("Expected %d but was %d", expected, item.Count())
	}
}

func TestSortedLinkedListItem_Value(t *testing.T) {
	expected := rand.Intn(100)
	item := linked.NewSortedLinkedListItem(expected)
	if item.Value() != expected {
		t.Errorf("Expected %d but was %d", expected, item.Value())
	}
}

func BenchmarkSortedLinkedList_InsertAndDelete10000_5000(b *testing.B) {
	var values [10000]int
	for i, _ := range values {
		values[i] = rand.Intn(5000)
	}
	for n := 0; n < b.N; n++ {
		list := linked.NewSortedLinkedList()
		for _, v := range values {
			list.Insert(v)
		}
		for _, v := range values {
			list.Delete(v, true)
		}
	}
}

func BenchmarkSortedLinkedList_AllAndUnique10000_5000(b *testing.B) {
	var values [10000]int
	for i, _ := range values {
		values[i] = rand.Intn(5000)
	}
	list := linked.NewSortedLinkedList()
	for _, v := range values {
		list.Insert(v)
	}
	for n := 0; n < b.N; n++ {
		list.GetAll()
		list.GetUnique()
	}
}
