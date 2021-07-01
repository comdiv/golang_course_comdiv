package dynmerger

import (
	"context"
	"github.com/comdiv/golang_course_comdiv/internal/superchan/pipe"
	"reflect"
	"sync"
)

// позволяет направлять несколько каналов
type DynamicMerger struct {
	selectsMap map[int]reflect.SelectCase // будем собирать не каналы, а селекторы, тут карта для регистрации
	selects    []reflect.SelectCase
	out        chan string
	idgen      int
	locker     sync.RWMutex
	wg         sync.WaitGroup
	swg        sync.WaitGroup
	stopper    chan struct{}
}

func New(inputs []chan string, out chan string) *DynamicMerger {
	result := &DynamicMerger{selectsMap: make(map[int]reflect.SelectCase), selects: make([]reflect.SelectCase, 0), out: out, stopper: make(chan struct{})}
	for _, ch := range inputs {
		result.Register(ch)
	}
	return result
}

// регистрирует новый канал и возвращает его токн
func (m *DynamicMerger) Register(ch chan string) {
	pipe.New(context.TODO(), ch, m.out, func(s string) string { return s }).Start()
}
