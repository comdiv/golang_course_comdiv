package linked

// SortedLinkedListItem - внутренний элемент индекса, хранит в себе значение,
// количество копий значения и указатель на соседние элементы индекса с меньшими
// и большими значениями соответственно
type SortedLinkedListItem struct {
	value int
	count int
	prev  *SortedLinkedListItem
	next  *SortedLinkedListItem
}

func NewSortedLinkedListItem(value int) *SortedLinkedListItem {
	return NewSortedLinkedListItemC(value, 1)
}
func NewSortedLinkedListItemC(value int, count int) *SortedLinkedListItem {
	return &SortedLinkedListItem{value: value, count: count}
}

func (l *SortedLinkedListItem) Inc() {
	l.count++
}

func (l *SortedLinkedListItem) Dec() {
	if l.count == 1 {
		l.Delete()
	} else {
		l.count--
	}
}

// Value - значение числа в SortedLinkedListItem
func (l *SortedLinkedListItem) Value() int {
	return l.value
}

// Count - количество значений Value
func (l *SortedLinkedListItem) Count() int {
	return l.count
}

// Prev - указатель на предыдущий элемент индекса
func (l *SortedLinkedListItem) Prev() *SortedLinkedListItem {
	return l.prev
}

// Next - указатель на следующий элемент индекса
func (l *SortedLinkedListItem) Next() *SortedLinkedListItem {
	return l.next
}

// InsertRight - добавляет элемент справа
func (l *SortedLinkedListItem) InsertRight(value int) { // final and self INSERT_RESULT_TYPE)
	newItem := &SortedLinkedListItem{prev: l, next: l.next, value: value, count: 1}
	if l.next != nil {
		l.next.prev = newItem
	}
	l.next = newItem
}

// InsertLeft - добавляет элемент слева
func (l *SortedLinkedListItem) InsertLeft(value int) { // final and self INSERT_RESULT_TYPE)
	newItem := &SortedLinkedListItem{prev: l.prev, next: l, value: value, count: 1}
	if l.prev != nil {
		l.prev.next = newItem
	}
	l.prev = newItem
}

// Delete удаление элемента из списка
func (l *SortedLinkedListItem) Delete() { // DELETE_RESULT, deleted items size
	if l.next != nil {
		l.next.prev = l.prev
	}
	if l.prev != nil {
		l.prev.next = l.next
	}
}
