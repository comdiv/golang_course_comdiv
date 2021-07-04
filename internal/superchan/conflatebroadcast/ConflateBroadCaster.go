package dynbroadcaster

import (
	"context"
	"github.com/comdiv/golang_course_comdiv/internal/superchan"
	"github.com/comdiv/golang_course_comdiv/internal/superchan/dynmerger"
	"sync"
	"sync/atomic"
)

// ConflateBroadCaster увязывает множество входов (поставщиков) и выходов (слушателей) в конфлейт режиме
// все входы активно читаются, чтение же ленивое текущего элемента в том темпе в котором это может себе позволить
// слушатель
type ConflateBroadCaster struct {
	// позволяет объединить сколько угодно входов в один выход
	merger *dynmerger.DynamicMerger
	// выходной канал для мержера, входной для конфлейтера
	in chan string
	// текущее значение
	current        string
	defaultContext context.Context
	listeners      []func(s string)
	cond           *sync.Cond
	messageId      int64
}

func (c *ConflateBroadCaster) WaitNew(last int64) <-chan struct{} {
	result := make(chan struct{})
	go func() {
		c.cond.L.Lock()
		defer c.cond.L.Unlock()
		if c.messageId <= last {
			c.cond.Wait()
		}
		close(result)
	}()
	return result
}

func New(ctx context.Context) *ConflateBroadCaster {
	if ctx == nil || ctx == context.TODO() {
		ctx = context.Background()
	}
	in := make(chan string)
	result := &ConflateBroadCaster{
		merger:         dynmerger.New(ctx, []chan string{}, in),
		in:             in,
		current:        "",
		defaultContext: ctx,
		listeners:      make([]func(s string), 0),
		cond:           sync.NewCond(&sync.Mutex{}),
	}
	return result
}

func (c *ConflateBroadCaster) StartAsync(ctx context.Context) superchan.Job {
	var wg sync.WaitGroup
	wg.Add(1)
	innerCtx, cancel := context.WithCancel(ctx)
	go func() {
		defer wg.Done()
		c.Start(innerCtx)
	}()
	return superchan.Job{
		Id:     superchan.NextJobId(),
		Cancel: cancel,
		Wait: func() {
			wg.Wait()
		},
	}
}

// синхронный старт рабочего цикла, который собственно просто собирает все с канала мержера
func (c *ConflateBroadCaster) Start(ctx context.Context) {
	for {
		select {
		case next, ok := <-c.in:
			if ok {
				c.cond.L.Lock()
				atomic.AddInt64(&c.messageId, 1)
				c.current = next
				c.cond.Broadcast()
				c.cond.L.Unlock()
			} else {
				return
			}
		case <-ctx.Done():
			return
		}
	}
}

// добавляет поставщика данных из которого производится чтение
func (c *ConflateBroadCaster) Publish(ctx context.Context, ch <-chan string) superchan.Job {
	return c.merger.Bind(ctx, ch)
}

type memo struct {
	c     *ConflateBroadCaster
	read  bool
	value string
}

func (m *memo) Get() string {
	if !m.read {
		m.value = m.c.current
		m.read = true
	}
	return m.value
}

func (c *ConflateBroadCaster) Listen(ctx context.Context, ch chan<- (func() string)) superchan.Job {
	if ctx == nil || ctx == context.TODO() {
		ctx = c.defaultContext
	}
	innerContext, cancel := context.WithCancel(ctx)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		var last int64 = 0
		defer wg.Done()
		for {
			select {
			case <-innerContext.Done():
				return
			case <-c.WaitNew(last):
				next := c.messageId
				m := memo{
					c: c,
				}
				select {
				case ch <- m.Get:
					last = next
				default:
					break
				}
			}
		}
	}()

	job := superchan.Job{
		Id:     superchan.NextJobId(),
		Cancel: cancel,
		Wait:   func() { wg.Wait() },
	}

	return job
}
