package SortedLinkedList

// SortedLinkedList - объект, который содержит данные и управляет сортировкой
type SortedLinkedList struct {
	itemCount int
	indexSize int
	head      *SortedLinkedListItem
	tail      *SortedLinkedListItem
}

// NewSortedLinkedList - конструктор для SortedLinkedList
// isUnique - признак, что значения должны быть уникальными
// lazySort - признак, что сортировка производится только в момент чтения
func NewSortedLinkedList() *SortedLinkedList {
	return &SortedLinkedList{}
}

// IndexSize - количество узлов значений в индексе, по сути число уникальных элементов
func (this *SortedLinkedList) IndexSize() int {
	return this.indexSize
}

// ItemCount - количество всех чисел в индексе без учета уникальности
func (this *SortedLinkedList) ItemCount() int {
	return this.itemCount
}

// Head - указатель на начальный элемент индекса
func (this *SortedLinkedList) Head() *SortedLinkedListItem {
	return this.head
}

// Tail - указатель на последний элемент индекса
func (this *SortedLinkedList) Tail() *SortedLinkedListItem {
	return this.tail
}

// GetDistinct - получить срез только уникальных упорядоченных чисел из индекса
func (this *SortedLinkedList) GetDistinct() []int {
	var result = make([]int, this.IndexSize())
	if this.head != nil {
		for i, current := 0, this.head; current != nil; i, current = i+1, current.next {
			result[i] = current.Value()
		}
	}
	return result
}

func (this *SortedLinkedList) GetAll() []int {
	var result = make([]int, this.ItemCount())
	if this.head != nil {
		for i, current := 0, this.head; current != nil; i, current = i+current.Count(), current.next {
			for j := i; j < i+current.Count(); j++ {
				result[j] = current.Value()
			}
		}
	}
	return result
}

const (
	FOUND_TYPE_NOT_FOUND = 0
	FOUND_TYPE_FOUND     = 1
	FOUND_TYPE_PREV      = -2
	FOUND_TYPE_NEXT      = 2
)

func (this *SortedLinkedList) FindItemFor(value int) (*SortedLinkedListItem, int) { // FOUND_TYPE_*
	if this.head == nil {
		return nil, FOUND_TYPE_NOT_FOUND
	}
	if this.head.value > value {
		return this.head, FOUND_TYPE_NEXT
	}
	if this.tail.value < value {
		return this.tail, FOUND_TYPE_PREV
	}
	for current := this.head; current != nil; current = current.next {
		if value == current.Value() {
			return current, FOUND_TYPE_FOUND
		}
		if value > current.Value() && current.Next() != nil && current.Next().Value() > value {
			return current, FOUND_TYPE_PREV
		}
	}
	panic("По идее алгоритм должен был уже сойтись без вариантов")
}

// Insert - добавляет элемент в SortedLinkedList, всегда в сортированном порядке
// если массив - уникальный, то если число уже есть - оно не добавляется, иначе всегда добавляется
func (this *SortedLinkedList) Insert(value int) bool { // true - INSERT, false - JUST INCREMENT
	this.itemCount++
	if this.head == nil { // первый элемент добавлям
		this.head = &SortedLinkedListItem{value: value, count: 1}
		this.tail = this.head
		this.indexSize++
		return true
	}
	target, targetType := this.FindItemFor(value)
	inserted := true
	switch targetType {
	case FOUND_TYPE_FOUND: // ничего не надо добавлять, только накрутим счетчик всех значений
		target.count++
		inserted = false
	case FOUND_TYPE_PREV: // присоединить справа
		target.InsertRight(value)
		if target == this.tail {
			this.tail = target.next
		}
		this.indexSize++
	case FOUND_TYPE_NEXT: // присоединить слева
		target.InsertLeft(value)
		if target == this.head {
			this.head = target.prev
		}
		this.indexSize++
	}
	return inserted
}

// InsertAll - вставка всех элементов из среза
func (this *SortedLinkedList) InsertAll(values []int) {
	for _, v := range values {
		this.Insert(v)
	}
}

// InsertAllVar - вставка всех элементов из variadic
func (this *SortedLinkedList) InsertAllVar(values ...int) {
	for _, v := range values {
		this.Insert(v)
	}
}

// Delete - удаление элемента из связанного списка
func (this *SortedLinkedList) Delete(value int) bool { // true - deleted, false - not found
	target, targetType := this.FindItemFor(value)
	if targetType == FOUND_TYPE_FOUND { // элемент найден
		this.indexSize--
		this.itemCount -= target.count
		if target == this.head {
			this.head = target.next
		}
		if target == this.tail {
			this.tail = target.prev
		}
		target.Delete()
		return true
	} else {
		return false
	}
}

// SortedLinkedListItem - внутренний элемент индекса, хранит в себе значение,
// количество копий значения и указатель на соседние элементы индекса с меньшими
// и большими значениями соответственно
type SortedLinkedListItem struct {
	value int
	count int
	prev  *SortedLinkedListItem
	next  *SortedLinkedListItem
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
