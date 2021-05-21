package sortedintlist

type ISortedIntList interface {
	Insert(value int) bool
	Delete(value int, all bool) bool
	Size() int
	UniqueSize() int
	GetAll() []int
	GetUnique() []int
}
