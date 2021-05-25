package sortedintlist

// IIntInsert контракт на поддержание вставки int значений
type IIntInsert interface {
	Insert(value int) bool
}

// IIntDelete контракт на поддержание удаления int значений
type IIntDelete interface {
	Delete(value int, all bool) bool
}

// IIntCollectionMutable - контракт на поддержания измененного состояния (вставка и удаления)
type IIntCollectionMutable interface {
	IIntInsert
	IIntDelete
}

// IIntList - контракт не изменяемого списка int
type IIntList interface {
	Size() int
	GetAll() []int
}

// IIntListMutable - контракт изменяемого списка int
type IIntListMutable interface {
	IIntCollectionMutable
	IIntList
}

// IIntSet - контракт не изменяемого набора int
type IIntSet interface {
	UniqueSize() int
	GetUnique() []int
}

// IIntSetMutable - контракт изменяемого набора int
type IIntSetMutable interface {
	IIntCollectionMutable
	IIntSet
}

// IIntMinMax - контракт на возврат максимального и минимального значения
type IIntMinMax interface {
	// IsIntRangeInitialized - определяет вообще определен ли минимакс-диапазон
	IsIntRangeInitialized() bool
	// GetMin  - вернуть минимальное значение или ошибку, если диапазон не инициализирован
	GetMin() (int, error)
	// GetMax  - вернуть максимальное значение или ошибку, если диапазон не инициализирован
	GetMax() (int, error)
}

func InsertAll(l IIntInsert, data []int) {
	for _, v := range data {
		l.Insert(v)
	}
}

func InsertAllVar(l IIntInsert, data ...int) {
	for _, v := range data {
		l.Insert(v)
	}
}
