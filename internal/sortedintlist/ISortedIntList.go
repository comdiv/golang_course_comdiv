package sortedintlist

type ISortedIntList interface {
	Insert(value int) bool
	Delete(value int, all bool) bool
	Size() int
	UniqueSize() int
	GetAll() []int
	GetUnique() []int
}

func InsertAll(l ISortedIntList, data []int) {
	for _, v := range data {
		l.Insert(v)
	}
}

func InsertAllVar(l ISortedIntList, data ...int) {
	for _, v := range data {
		l.Insert(v)
	}
}
