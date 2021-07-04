package pipe_test

import (
	"context"
	"fmt"
	"github.com/comdiv/golang_course_comdiv/internal/superchan/pipe"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

var emptyTransformer = func(s string) string { return s }

func prefixer(prefix string) func(s string) string {
	return func(s string) string {
		return prefix + "-" + s
	}
}

// проверяем что вообще создается труба
func TestNewPipe(t *testing.T) {
	var pipe = pipe.New(make(chan string), make(chan string), emptyTransformer)
	assert.NotNil(t, pipe)
}

func TestPipeIsPassiveBeforeStart(t *testing.T) {
	in := make(chan string)
	out := make(chan string)
	pipe := pipe.New(in, out, emptyTransformer)
	assert.NotNil(t, pipe)
	// так как мы не стартовали пайп и он не читает in, то запущенная ниже корутина перехватит весь вход
	var wg sync.WaitGroup
	noPipeCounter := 0
	totalCount := 10000
	wg.Add(1)
	go func() {
		defer wg.Done()
		for range in {
			noPipeCounter++
		}
	}()
	for i := 0; i < totalCount; i++ {
		in <- strconv.Itoa(i)
	}
	close(in)
	wg.Wait()
	assert.Equal(t, totalCount, noPipeCounter)
}

func TestPipeStart(t *testing.T) {
	in := make(chan string)
	out := make(chan string)
	pipe := pipe.New(in, out, prefixer("test"))

	// трубу надо стартовать
	var pipegroup sync.WaitGroup
	pipegroup.Add(1)
	go func(){
		defer pipegroup.Done()
		pipe.Start(context.Background())
	}()

	// ну и она фоново перегоняет in в out с заданной трансформацией

	var wg sync.WaitGroup
	counter := 0
	totalCount := 10000
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range out {
			// проверяем что работала трансформация
			assert.True(t, strings.HasPrefix(v, "test-"))
			counter++
		}
	}()
	for i := 0; i < totalCount; i++ {
		in <- strconv.Itoa(i)
	}
	close(in)
	pipegroup.Wait() // так как in закрыт, то труба завершится, сама труба out не закрывает (нет полномочий),
	// но вдруг был буфер на in, дожищдаемя, чтобы он был обработан
	close(out)
	wg.Wait()
	// все перегнано
	assert.Equal(t, totalCount, counter)
}

func TestPipe_StartAsync(t *testing.T) {
	in := make(chan string)
	out := make(chan string)
	pipe := pipe.New(in, out, prefixer("test"))

	// трубу надо стартовать
	pipeJob := pipe.StartAsync(context.Background())
	var wg sync.WaitGroup
	counter := 0
	totalCount := 10000
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range out {
			// проверяем что работала трансформация
			assert.True(t, strings.HasPrefix(v, "test-"))
			counter++
		}
	}()
	for i := 0; i < totalCount; i++ {
		in <- strconv.Itoa(i)
	}
	close(in)
	pipeJob.Finish()
	// но вдруг был буфер на in, дожищдаемя, чтобы он был обработан
	close(out)
	wg.Wait()
	// все перегнано
	assert.Equal(t, totalCount, counter)
}


func TestUsesContext(t *testing.T) {
	in := make(chan string)
	out := make(chan string)
	ctx, cancel := context.WithCancel(context.TODO())
	pipe := pipe.New(in, out, prefixer("test"))
	// трубу надо стартовать
	pipeJob := pipe.StartAsync(ctx)
	// ну и она фоново перегоняет in в out с заданной трансформацией
	counter := 0
	totalCount := 10000
	go func() {
		for v := range out {
			// проверяем что работала трансформация
			assert.True(t, strings.HasPrefix(v, "test-"))
			counter++
			if counter > 1000 {
				go cancel()
			}
		}
	}()
	go func() {
		for i := 0; i < totalCount; i++ {
			in <- strconv.Itoa(i)
		}
		close(in)
	}()
	pipeJob.Wait() //  по идее закроется примерно на 1000-1001 элементе
	assert.GreaterOrEqual(t, counter, 1000)
	assert.LessOrEqual(t, counter, 1005)
}

// смотрим работу "трубы" с автозакрытием
func TestNewTransformChannel(t *testing.T) {
	in := make(chan string)
	out:= pipe.PipeChannel(context.TODO(),in, prefixer("test"))
	counter := 0
	totalCount := 10000
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := 0; i < totalCount; i++ {
			in <- strconv.Itoa(i)
		}
		close(in)
	}()
	for range out {
		counter++
	}
	assert.Equal(t, totalCount, counter)
}



// в этом тесте pipe подключен к "входу" и читает его конкурентно с другой корутиной
func TestPipeLineIsWorkingAfterStartConcurrently(t *testing.T) {
	in := make(chan string)
	out := make(chan string)
	pipe := pipe.New( in, out, emptyTransformer)
	assert.NotNil(t, pipe)
	pipeJob := pipe.StartAsync(context.TODO())
	var wg sync.WaitGroup
	pipeCounter := 0
	noPipeCounter := 0
	totalCount := 10000
	wg.Add(2)
	go func() {
		defer wg.Done()
		for range in {
			noPipeCounter++
			if noPipeCounter == 1000 {
				time.Sleep(10 * time.Millisecond) // ci ubuntu ranner always make preference for this counter, should pause
			}
		}
	}()
	go func() {
		defer wg.Done()
		for range out {
			pipeCounter++
		}
	}()
	for i := 0; i < totalCount; i++ {
		in <- strconv.Itoa(i)
	}
	close(in)
	pipeJob.Finish()
	close(out)
	wg.Wait()
	// все должно было обработаться
	assert.Equal(t, totalCount, noPipeCounter+pipeCounter)
	fmt.Println(noPipeCounter, pipeCounter)
	assert.Greater(t, noPipeCounter, 100)
	assert.Greater(t, pipeCounter, 100)
}
