package dynmerger

import (
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
}

func New(inputs []chan string, out chan string) *DynamicMerger {
	result := &DynamicMerger{selectsMap: make(map[int]reflect.SelectCase), selects: make([]reflect.SelectCase, 0), out: out}
	for _, ch := range inputs {
		_ = result.Register(ch)
	}
	result.Start()
	return result
}

// регистрирует новый канал и возвращает его токн
func (m *DynamicMerger) Register(ch chan string) int {
	m.locker.Lock()
	defer m.locker.Unlock()
	m.idgen++
	token := m.idgen
	selectCase := reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)}
	m.selectsMap[token] = selectCase
	m.selects = append(m.selects, selectCase)
	return token
}

// регистрирует новый канал и возвращает его токн
func (m *DynamicMerger) Unregister(token int) {
	m.locker.Lock()
	defer m.locker.Unlock()
	delete(m.selectsMap, token)
	newselects := make([]reflect.SelectCase,0)
	for _, c := range m.selectsMap {
		newselects = append(newselects, c)
	}

	m.selects = newselects
}


func (m *DynamicMerger) Start() {
	// тут не стану заморачиваться за синхронизацию как в Pipe - там просто хотелось это отработать
	go func() {
		for {
			m.locker.RLock()
			_, value, ok := reflect.Select(m.selects)
			if ok {
				m.out <- value.String()
			}
			m.locker.RUnlock()
		}
	}()
}
