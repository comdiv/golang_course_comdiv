package linked

// SortedLinkedList - объект, который содержит данные и управляет сортировкой
type SortedLinkedList struct {
	itemCount int
	indexSize int
	head      *SortedLinkedListItem
	tail      *SortedLinkedListItem
}

func (l *SortedLinkedList) IsIntRangeInitialized() bool {
	panic("implement me")
}

func (l *SortedLinkedList) GetMin() (int, error) {
	panic("implement me")
}

func (l *SortedLinkedList) GetMax() (int, error) {
	panic("implement me")
}

// NewSortedLinkedList - конструктор для SortedLinkedList
// isUnique - признак, что значения должны быть уникальными
// lazySort - признак, что сортировка производится только в момент чтения
func NewSortedLinkedList() *SortedLinkedList {
	return &SortedLinkedList{}
}

// UniqueSize - количество узлов значений в индексе, по сути число уникальных элементов
func (l *SortedLinkedList) UniqueSize() int {
	return l.indexSize
}

// Size - количество всех чисел в индексе без учета уникальности
func (l *SortedLinkedList) Size() int {
	return l.itemCount
}

// Head - указатель на начальный элемент индекса
func (l *SortedLinkedList) Head() *SortedLinkedListItem {
	return l.head
}

// Tail - указатель на последний элемент индекса
func (l *SortedLinkedList) Tail() *SortedLinkedListItem {
	return l.tail
}

// GetUnique - получить срез только уникальных упорядоченных чисел из индекса
func (l *SortedLinkedList) GetUnique() []int {
	var result = make([]int, l.UniqueSize())
	if l.head != nil {
		for i, current := 0, l.head; current != nil; i, current = i+1, current.next {
			result[i] = current.Value()
		}
	}
	return result
}

func (l *SortedLinkedList) GetAll() []int {
	var result = make([]int, l.Size())
	if l.head != nil {
		for i, current := 0, l.head; current != nil; i, current = i+current.Count(), current.next {
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

func (l *SortedLinkedList) FindItemFor(value int) (*SortedLinkedListItem, int) { // FOUND_TYPE_*
	if l.head == nil {
		return nil, FOUND_TYPE_NOT_FOUND
	}
	if l.head.value > value {
		return l.head, FOUND_TYPE_NEXT
	}
	if l.tail.value < value {
		return l.tail, FOUND_TYPE_PREV
	}
	for current := l.head; current != nil; current = current.next {
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
func (l *SortedLinkedList) Insert(value int) bool { // true - INSERT, false - JUST INCREMENT
	l.itemCount++
	if l.head == nil { // первый элемент добавлям
		l.head = &SortedLinkedListItem{value: value, count: 1}
		l.tail = l.head
		l.indexSize++
		return true
	}
	target, targetType := l.FindItemFor(value)
	inserted := true
	switch targetType {
	case FOUND_TYPE_FOUND: // ничего не надо добавлять, только накрутим счетчик всех значений
		target.count++
		inserted = false
	case FOUND_TYPE_PREV: // присоединить справа
		target.InsertRight(value)
		if target == l.tail {
			l.tail = target.next
		}
		l.indexSize++
	case FOUND_TYPE_NEXT: // присоединить слева
		target.InsertLeft(value)
		if target == l.head {
			l.head = target.prev
		}
		l.indexSize++
	}
	return inserted
}

// InsertAll - вставка всех элементов из среза
func (l *SortedLinkedList) InsertAll(values []int) {
	for _, v := range values {
		l.Insert(v)
	}
}

// InsertAllVar - вставка всех элементов из variadic
func (l *SortedLinkedList) InsertAllVar(values ...int) {
	for _, v := range values {
		l.Insert(v)
	}
}

// Delete - удаление элемента из связанного списка
func (l *SortedLinkedList) Delete(value int, all bool) bool { // true - deleted, false - not found
	target, targetType := l.FindItemFor(value)
	if targetType == FOUND_TYPE_FOUND { // элемент найден
		if all || target.Count() == 1 {
			l.indexSize--
			l.itemCount -= target.count
			if target == l.head {
				l.head = target.next
			}
			if target == l.tail {
				l.tail = target.prev
			}
			target.Delete()
		} else {
			l.itemCount--
			target.count--
		}
		return true
	} else {
		return false
	}
}
