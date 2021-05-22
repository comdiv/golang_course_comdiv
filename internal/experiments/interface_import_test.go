package experiments

import (
	"fmt"
	"testing"
)
import "github.com/stretchr/testify/assert"

// определим интерфейс
type ISomeIFace interface {
	Hello() string
	IsGood() bool
}

/* -- вот так делать нельзя,
func (s *ISomeIFace) Hello() string {
	return "world"
}
*/
// и это при том, что так конечно можно:

func Hello2(s *ISomeIFace) string {
	return "world2"
}

type FaceNoImport struct {
}

func (s *FaceNoImport) Hello() string {
	return "world"
}

func (s *FaceNoImport) IsGood() bool {
	return true
}

// пока не определил обоих методова у FaceNoImport не мог выполнить этой строчки
var i1 ISomeIFace = new(FaceNoImport)

// но
type FaceImport struct {
	ISomeIFace // типа унаследовался
}

// и вуаля - типа я соответствую интерфейсу, хотя на самом деле ничего не определил
var i2 ISomeIFace = new(FaceImport)

func TestI1(t *testing.T) {
	assert.Equal(t, "world", i1.Hello())
	assert.Equal(t, true, i1.IsGood())
}

func TestI2(t *testing.T) {
	// И вот оказываается я только в рантайме получу нечто на редкость невнятное
	// вот просто нафига так было делать с этими интерфейсами???
	// и это при том, что понятия дефолтная реализация интерфейса отсутствует
	defer func() {
		if r := recover(); r != nil {
			errString := fmt.Sprint(r)
			assert.Equal(t, "runtime error: invalid memory address or nil pointer dereference", errString)
		}
	}()
	assert.Equal(t, "", i2.Hello())
	assert.Equal(t, false, i2.IsGood())
}
