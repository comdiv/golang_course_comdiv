package pipe

import (
	"context"
	"github.com/comdiv/golang_course_comdiv/internal/superchan"
	"sync"
)

type Pipe struct {
	in        <-chan string
	out       chan<- string
	transform func(s string) string
	autoClose bool
}

func New(in <-chan string, out chan<- string, transform func(s string) string) *Pipe {
	return &Pipe{in: in, out: out, transform: transform}
}

// на основе Pipe активный канал, который будет закрыт после разбора всего in или остановки контекста
func PipeChannel(ctx context.Context, in <-chan string, transform func(s string) string) <-chan string {
	out := make(chan string)
	pipe := New(in, out, transform)
	pipe.autoClose = true
	go pipe.Start(ctx)
	return out
}

// стартует трубу с заданным контекстом, возвращает функицю для ожидания завершения
func (p *Pipe) StartAsync(ctx context.Context) superchan.Job {
	var wg sync.WaitGroup
	wg.Add(1)
	innerCtx, cancel := context.WithCancel(ctx)
	go func() {
		defer wg.Done()
		p.Start(innerCtx)
	}()
	return superchan.Job{
		Id:     superchan.NextJobId(),
		Cancel: cancel,
		Wait: func() {
			wg.Wait()
		},
	}
}

// стартует "трубу" - данные начинают перекачиваться с трансформацией
func (p *Pipe) Start(ctx context.Context) { // returns waiter function to join started job
	if p.autoClose {
		defer close(p.out)
	}
	for {
		select {
		case <-ctx.Done():
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
}
