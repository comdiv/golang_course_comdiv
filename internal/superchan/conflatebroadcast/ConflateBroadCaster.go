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
	messageId      int64
}

func New(ctx context.Context) *ConflateBroadCaster {
	if ctx == nil || ctx == context.TODO() {
		ctx = context.Background()
	}
	in := make(chan string)
	return &ConflateBroadCaster{
		merger:         dynmerger.New(ctx, []chan string{}, in),
		in:             in,
		current:        "",
		defaultContext: ctx,
		listeners:      make([]func(s string), 0),
	}
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
				atomic.AddInt64(&c.messageId, 1)
				c.current = next
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

func (c *ConflateBroadCaster) Listen(ctx context.Context, ch chan<- string) superchan.Job {

	if ctx == nil || ctx == context.TODO() {
		ctx = c.defaultContext
	}
	innerContext, cancel := context.WithCancel(ctx)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		var last int64 = 0
		for {
			select {
			case <-innerContext.Done():
				return
			default:
				break
			}
			if c.messageId > last {
				select {
				case ch <- c.current:
					last = c.messageId
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
