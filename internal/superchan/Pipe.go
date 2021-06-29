package superchan

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
)

type Pipe struct {
	context    context.Context
	in         <-chan string
	out        chan<- string
	transform  func(s string) string
	started    int32
	stopper    chan struct{}
	wg         sync.WaitGroup
	operlock   sync.Mutex
	joinMarker bool
	autoClose  bool
}

func NewPipe(context context.Context, in <-chan string, out chan<- string, transform func(s string) string) *Pipe {
	return &Pipe{context: context, in: in, out: out, transform: transform, stopper: make(chan struct{},1)}
}

// на основе Pipe активный канал, который будет закрыт после разбора всего in или остановки контекста
func NewTransformChannel(context context.Context, in <-chan string, transform func(s string) string) <-chan string {
	out := make(chan string)
	pipe := NewPipe(context, in, out, transform)
	pipe.autoClose = true
	pipe.Start()
	return out
}

// стартует "трубу" - данные начинают перекачиваться с трансформацией
func (p *Pipe) Start() {
	p.operlock.Lock()
	defer p.operlock.Unlock()
	defer fmt.Println(p.started)
	requireStart := atomic.CompareAndSwapInt32(&p.started, 0, 1)
	if requireStart {
		p.wg.Add(1)
		go func() {
			defer atomic.StoreInt32(&p.started, 0)
			defer p.wg.Done()
			if p.autoClose {
				defer close(p.out)
			}
			for {
				select {
				case <-p.context.Done():
					return

				case <-p.stopper:
					return
				case s, ok := <-p.in:
					if ok {
						st := p.transform(s)
						p.out <- st
					} else {
						return
					}
				}
			}
		}()
	}
}

// останаваливает трубу - досрочно прерывает основную рутину
func (p *Pipe) Stop() {
	p.operlock.Lock()
	defer p.operlock.Unlock()
	// признак, что Start был вызван и при этом сама горутина еще не завершилась
	requireStop := atomic.CompareAndSwapInt32(&p.started, 1, 0)
	if requireStop {
		// возможно именно сейчас горутина закрывается и значение в stopper просто зависнет
		// в любом случае надо будет пересоздать канал
		defer func() {
			close(p.stopper)
			p.stopper = make(chan struct{}, 1)
		}()
		// теперь шлем сигнал и ждем завершения
		p.stopper <- struct{}{}
		p.wg.Wait()
	}
}

func (p *Pipe) Wait() {
	p.wg.Wait()
}
