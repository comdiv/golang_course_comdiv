package slices

import (
	"github.com/comdiv/golang_course_comdiv/internal/sortedintlist"
)

// SortedIntListSliced - реализация ISortedIntReplProvider
type SortedIntListSliced struct {
	data       []int
	dups       map[int]int
	uniqueSize int
	totalSize  int
}

func (s *SortedIntListSliced) IsIntRangeInitialized() bool {
	panic("implement me")
}

func (s *SortedIntListSliced) GetMin() (int, error) {
	panic("implement me")
}

func (s *SortedIntListSliced) GetMax() (int, error) {
	panic("implement me")
}

// hack to force implementation of interface in goland and check that it implements it
func sortedintlistslicedImplements() sortedintlist.ISortedIntReplProvider {
	return &SortedIntListSliced{}
}

func NewSortedIntListSliced() *SortedIntListSliced {
	return NewSortedIntListSlicedWithData(nil)
}

func NewSortedIntListSlicedWithData(initialdata []int) *SortedIntListSliced {
	result := &SortedIntListSliced{data: []int{}, dups: make(map[int]int)}
	if nil != initialdata {
		for _, v := range initialdata {
			result.Insert(v)
		}
	}
	return result
}

func (s *SortedIntListSliced) Insert(value int) bool {

	index, insertIndex := LastIndexOf(s.data, value, true)
	if index < 0 { // такого значения еще не было
		s.uniqueSize++
	}
	switch {
	case index == LAST_INDEX_AFTER:
		s.data = append(s.data, value)
	case index == LAST_INDEX_BEFORE:
		s.data = append([]int{value}, s.data...)
	case index >= 0:
		s.dups[value]++
	case s.totalSize == 0:
		s.data = append(s.data, value)
	default:
		if cap(s.data) > len(s.data) {
			s.data = append(s.data, 0)
			copy(s.data[insertIndex+2:], s.data[insertIndex+1:len(s.data)-1])
			s.data[insertIndex+1] = value
		} else {
			newdata := make([]int, len(s.data)+1)
			copy(newdata, s.data[:insertIndex+1])
			newdata[insertIndex+1] = value
			copy(newdata[insertIndex+2:], s.data[insertIndex+1:])
			s.data = newdata
		}
	}
	s.totalSize++
	return index < 0 // Только для совсем новых значений для конвенции с Linked мы возвращем true
}

func (s *SortedIntListSliced) Delete(value int, all bool) bool {
	index, _ := LastIndexOf(s.data, value, true)
	if index >= 0 {
		var duplicates = s.dups[value]
		if all || duplicates == 0 { // число полностью уходит
			s.uniqueSize--
			copy(s.data[index:], s.data[index+1:])
			s.data = s.data[:len(s.data)-1]
		} else {
			s.dups[value]--
		}
		s.totalSize--
		if all && duplicates > 0 {
			s.totalSize -= duplicates
		}

		return true
	}
	return false
}

const (
	LAST_NOT_FOUND    = -1
	LAST_INDEX_BEFORE = -2
	LAST_INDEX_AFTER  = -3
)

// LastIndexOf находит последний индекс указанного значения в слайсе
// раз решили все делать самостоятельно то и такая функция тоже своя
// спец значения -1 - не найден, -2 - не найден и меньше первого, -3 - не найден и больше последнего
func LastIndexOf(data []int, v int, isSorted bool) (int, int) { // index, and last current pos
	if len(data) == 0 {
		return LAST_NOT_FOUND, 0
	}
	if v > data[len(data)-1] {
		return LAST_INDEX_AFTER, len(data) - 1
	}
	if v < data[0] {
		return LAST_INDEX_BEFORE, 0
	}

	if !isSorted {
		var i, current int
		for i = len(data) - 1; i >= 0; i-- {
			current = data[i]
			if current == v {
				return i, i
			}

		}
		return LAST_NOT_FOUND, i
	}

	// sorted array search optimization
	var lowerPoint = 0
	var upperPoint = len(data) - 1
	var currentPoint = upperPoint / 2

	for {
		current := data[currentPoint]
		if current == v {
			var last int
			for last = currentPoint; last < len(data)-1 && data[last+1] == v; last++ {
			}
			return last, last
		} else if v > current && currentPoint < upperPoint {
			lowerPoint = currentPoint + 1
			if currentPoint == upperPoint-1 {
				currentPoint = upperPoint
			} else {
				currentPoint = currentPoint + ((upperPoint - currentPoint) / 2)
			}

		} else if v < current && currentPoint > lowerPoint {
			upperPoint = currentPoint - 1
			if currentPoint == lowerPoint+1 {
				currentPoint = lowerPoint
			} else {
				currentPoint = lowerPoint + ((currentPoint - lowerPoint) / 2)
			}
		} else {
			break
		}
	}

	if data[currentPoint] > v {
		currentPoint--
	}

	return LAST_NOT_FOUND, currentPoint
}

func (s *SortedIntListSliced) Size() int {
	return s.totalSize
}

func (s *SortedIntListSliced) UniqueSize() int {
	return s.uniqueSize
}

func (s *SortedIntListSliced) GetAll() []int {
	result := make([]int, s.totalSize)
	var targetindex int
	for _, v := range s.data {
		result[targetindex] = v
		targetindex++
		for d := 0; d < s.dups[v]; d++ {
			result[targetindex] = v
			targetindex++
		}
	}
	return result
}

func (s *SortedIntListSliced) GetUnique() []int {
	result := make([]int, s.uniqueSize)
	copy(result, s.data)
	return result
}
