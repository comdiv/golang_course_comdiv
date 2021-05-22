package slices

import "github.com/comdiv/golang_course_comdiv/internal/sortedintlist"

// SortedIntListSliced - реализация ISortedIntList
type SortedIntListSliced struct {
	data []int
}

// hack to force implementation of interface in goland and check that it implements it
func sortedintlistslicedImplements() sortedintlist.ISortedIntList {
	return &SortedIntListSliced{}
}

func NewSortedIntListSliced() *SortedIntListSliced {
	return NewSortedIntListSlicedWithData(nil)
}

func NewSortedIntListSlicedWithData(initialdata []int) *SortedIntListSliced {
	result := &SortedIntListSliced{data: []int{}}
	if nil != initialdata {
		for _, v := range initialdata {
			result.Insert(v)
		}
	}
	return result
}

func (s SortedIntListSliced) Insert(value int) bool {
	panic("implement me")
}

func (s SortedIntListSliced) Delete(value int, all bool) bool {
	panic("implement me")
}

func (s SortedIntListSliced) Size() int {
	panic("implement me")
}

func (s SortedIntListSliced) UniqueSize() int {
	panic("implement me")
}

func (s SortedIntListSliced) GetAll() []int {
	panic("implement me")
}

func (s SortedIntListSliced) GetUnique() []int {
	panic("implement me")
}
