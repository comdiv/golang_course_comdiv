package sortedintlistgentest

import (
	"fmt"
	"github.com/comdiv/golang_course_comdiv/internal/sortedintlist"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func GenericTestSorted_GetUnique(l sortedintlist.IIntSetMutable, t *testing.T) {
	sortedintlist.InsertAllVar(l, 8, 1, 2, 4, 5, 4, 4, 5, 6, 1)
	unique := l.GetUnique()
	expected := []int{1, 2, 4, 5, 6, 8}
	assert.ElementsMatch(t, expected, unique)
}

func GenericTestSorted_GetAll(l sortedintlist.IIntListMutable, t *testing.T) {
	sortedintlist.InsertAllVar(l, 8, 1, 2, 4, 5, 4, 4, 5, 6, 1)
	all := l.GetAll()
	expected := []int{1, 1, 2, 4, 4, 4, 5, 5, 6, 8}
	assert.ElementsMatch(t, expected, all)
}

func GenericTestSorted_Size(l sortedintlist.IIntListMutable, t *testing.T) {
	sortedintlist.InsertAllVar(l, 1, 2, 4, 4, 4)
	assert.Equal(t, 5, l.Size())
}

func GenericTestSorted_UniqueSize(l sortedintlist.IIntSetMutable, t *testing.T) {
	sortedintlist.InsertAllVar(l, 1, 2, 4, 4, 4)
	assert.Equal(t, 3, l.UniqueSize())
}

func GenericTestSorted_InsertList(l sortedintlist.IIntListMutable, t *testing.T) {
	var inserted bool
	inserted = l.Insert(1)
	assert.True(t, inserted)
	assert.Equal(t, 1, l.Size())

	inserted = l.Insert(10)
	assert.True(t, inserted)
	assert.Equal(t, 2, l.Size())

	inserted = l.Insert(10)
	assert.False(t, inserted)
	assert.Equal(t, 3, l.Size())
}

func GenericTestSorted_InsertSet(l sortedintlist.IIntSetMutable, t *testing.T) {
	var inserted bool
	inserted = l.Insert(1)
	assert.True(t, inserted)
	assert.Equal(t, 1, l.UniqueSize())

	inserted = l.Insert(10)
	assert.True(t, inserted)
	assert.Equal(t, 2, l.UniqueSize())

	inserted = l.Insert(10)
	assert.False(t, inserted)
	assert.Equal(t, 2, l.UniqueSize())
}

func GenericTestSorted_DeleteList(l sortedintlist.IIntListMutable, t *testing.T) {
	l.Insert(1)
	l.Insert(10)
	l.Insert(11)
	l.Insert(12)
	l.Insert(12)
	l.Insert(12)
	if !(l.Size() == 6) {
		t.Errorf("%v", l.Size())
	}
	var deleted bool
	deleted = l.Delete(10, true)
	assert.True(t, deleted)
	assert.Equal(t, 5, l.Size())

	deleted = l.Delete(77777, true)
	assert.False(t, deleted)
	assert.Equal(t, 5, l.Size())

	deleted = l.Delete(12, false)
	assert.True(t, deleted)
	assert.Equal(t, 4, l.Size())

	deleted = l.Delete(12, true)
	assert.True(t, deleted)
	assert.Equal(t, 2, l.Size())
}

func GenericTestSorted_DeleteSet(l sortedintlist.IIntSetMutable, t *testing.T) {
	l.Insert(1)
	l.Insert(10)
	l.Insert(11)
	l.Insert(12)
	l.Insert(12)
	l.Insert(12)
	if !(l.UniqueSize() == 4) {
		t.Errorf("%v", l.UniqueSize())
	}
	var deleted bool
	deleted = l.Delete(10, true)
	assert.True(t, deleted)
	assert.Equal(t, 3, l.UniqueSize())

	deleted = l.Delete(77777, true)
	assert.False(t, deleted)
	assert.Equal(t, 3, l.UniqueSize())

	deleted = l.Delete(12, false)
	assert.True(t, deleted)
	assert.Equal(t, 3, l.UniqueSize())

	deleted = l.Delete(12, true)
	assert.True(t, deleted)
	assert.Equal(t, 2, l.UniqueSize())

}

func GenericTestSorted_MinMax(minmax sortedintlist.IIntMinMax, t *testing.T) {

	l, ok := minmax.(sortedintlist.IIntListMutable)
	if !ok {
		panic(fmt.Sprintf("Not l list given for test! %v", minmax))
	}
	assert.False(t, minmax.IsIntRangeInitialized(), "Should not be initialized at start")
	_, err := minmax.GetMin()
	assert.NotNil(t, err, "Should be error to ask min from empty list")
	_, err = minmax.GetMax()
	assert.NotNil(t, err, "Should be error to ask max from empty list")

	r := rand.New(rand.NewSource(DEFAULT_BENCH_DATA_SEED))
	basevalue := r.Intn(DEFAULT_BENCH_DATA_SIZE)
	l.Insert(basevalue)
	assert.True(t, minmax.IsIntRangeInitialized(), "Should be initialized if has values")
	min, err := minmax.GetMin()
	assert.Nil(t, err, "Min should be working after initialization")
	assert.Equal(t, basevalue, min)
	max, err := minmax.GetMax()
	assert.Nil(t, err, "Min should be working after initialization")
	assert.Equal(t, basevalue, max)

	delata := r.Intn(1000) + 500
	expectedmin := basevalue - delata
	l.Insert(expectedmin)
	expectedmax := basevalue + delata
	l.Insert(expectedmax)

	min, err = minmax.GetMin()
	assert.Nil(t, err, "Min should be working after initialization")
	assert.Equal(t, expectedmin, min)

	max, err = minmax.GetMax()
	assert.Nil(t, err, "Min should be working after initialization")
	assert.Equal(t, expectedmax, max)
}
