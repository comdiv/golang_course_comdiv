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
	wg         sync.WaitGroup
	stopper    chan struct{}
}

func New(inputs []chan string, out chan string) *DynamicMerger {
	result := &DynamicMerger{selectsMap: make(map[int]reflect.SelectCase), selects: make([]reflect.SelectCase, 0), out: out, stopper: make(chan struct{})}
	for _, ch := range inputs {
		result.directRegister(ch)
	}
	result.Start()
	return result
}

func (m *DynamicMerger) directRegister(ch chan string){
	m.idgen++
	selectCase := reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)}
	m.selectsMap[m.idgen] = selectCase
	m.selects = append(m.selects, selectCase)
}

// регистрирует новый канал и возвращает его токн
func (m *DynamicMerger) Register(ch chan string) int {
	m.locker.Lock()
	defer m.locker.Unlock()
	m.stopper <- struct{}{}
	m.wg.Wait()
	m.directRegister(ch)
	m.Start()
	return m.idgen
}

// регистрирует новый канал и возвращает его токн
func (m *DynamicMerger) Unregister(token int) {
	m.locker.Lock()
	defer m.locker.Unlock()
	m.stopper <- struct{}{}
	m.wg.Wait()
	delete(m.selectsMap, token)
	newselects := make([]reflect.SelectCase, 0)
	for _, c := range m.selectsMap {
		newselects = append(newselects, c)
	}

	m.selects = newselects
	m.Start()
}

func (m *DynamicMerger) Start() {
	m.wg.Add(1)
	// тут не стану заморачиваться за синхронизацию как в Pipe - там просто хотелось это отработать
	go func(){
		defer m.wg.Done()
		for {
			select {
				case <- m.stopper:
					return
			default:
				break
			}
			_, value, ok := reflect.Select(m.selects)
			if ok {
				m.out <- value.String()
			}
		}
	}()
}
