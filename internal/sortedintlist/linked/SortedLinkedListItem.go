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

func (this *SortedLinkedListItem) Inc() {
	this.count++
}

func (this *SortedLinkedListItem) Dec() {
	if this.count == 1 {
		this.Delete()
	} else {
		this.count--
	}
}

// Value - значение числа в SortedLinkedListItem
func (this *SortedLinkedListItem) Value() int {
	return this.value
}

// Count - количество значений Value
func (this *SortedLinkedListItem) Count() int {
	return this.count
}

// Prev - указатель на предыдущий элемент индекса
func (this *SortedLinkedListItem) Prev() *SortedLinkedListItem {
	return this.prev
}

// Next - указатель на следующий элемент индекса
func (this *SortedLinkedListItem) Next() *SortedLinkedListItem {
	return this.next
}

// InsertRight - добавляет элемент справа
func (this *SortedLinkedListItem) InsertRight(value int) { // final and self INSERT_RESULT_TYPE)
	newItem := &SortedLinkedListItem{prev: this, next: this.next, value: value, count: 1}
	if this.next != nil {
		this.next.prev = newItem
	}
	this.next = newItem
}

// InsertLeft - добавляет элемент слева
func (this *SortedLinkedListItem) InsertLeft(value int) { // final and self INSERT_RESULT_TYPE)
	newItem := &SortedLinkedListItem{prev: this.prev, next: this, value: value, count: 1}
	if this.prev != nil {
		this.prev.next = newItem
	}
	this.prev = newItem
}

// Delete удаление элемента из списка
func (this *SortedLinkedListItem) Delete() { // DELETE_RESULT, deleted items size
	if this.next != nil {
		this.next.prev = this.prev
	}
	if this.prev != nil {
		this.prev.next = this.next
	}
}
