package dynmerger

import (
	"context"
	"github.com/comdiv/golang_course_comdiv/internal/superchan"
	"github.com/comdiv/golang_course_comdiv/internal/superchan/pipe"
)

// позволяет направлять несколько каналов
type DynamicMerger struct {
	jobs           map[int32]superchan.Job // карта функций отмен отдельных пайпов
	out            chan string
	defaultContext context.Context
}

func New(ctx context.Context, inputs []chan string, out chan string) *DynamicMerger {
	result := &DynamicMerger{jobs: make(map[int32]superchan.Job), out: out, defaultContext: ctx}
	if result.defaultContext == nil || result.defaultContext == context.TODO() {
		result.defaultContext = context.Background()
	}
	for _, ch := range inputs {
		result.Bind(ctx, ch)
	}
	return result
}

// регистрирует новый канал и возвращает его токн
func (m *DynamicMerger) Bind(ctx context.Context, ch chan string) superchan.Job {
	parentCtx := ctx
	if parentCtx == nil || parentCtx == context.TODO() {
		parentCtx = m.defaultContext
	}
	pipe := pipe.New(ch, m.out, func(s string) string { return s })
	job := pipe.StartAsync(parentCtx)
	job.OnFinish = func() {
		delete(m.jobs, job.Id)
	}
	m.jobs[job.Id] = job
	return job
}
