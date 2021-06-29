package dynmerger_test

import (
	"github.com/comdiv/golang_course_comdiv/internal/superchan/dynmerger"
	"github.com/stretchr/testify/assert"
	"strconv"
	"sync"
	"testing"
	"time"
)

// просто проверяем что вообще работает
func TestDynamicMerger_Start(t *testing.T) {
	in1 := make(chan string)
	in2 := make(chan string)
	out := make(chan string)
	dynmerger.New([]chan string{in1, in2}, out)
	var wg sync.WaitGroup
	counter := 0
	perInCount := 100
	wg.Add(1)
	go func(){
		defer wg.Done()
		for range out {
			counter++
		}
	}()
	for i:=0;i<perInCount;i++{
		in1 <- strconv.Itoa(i)
		in2 <- strconv.Itoa(i)
	}
	time.Sleep(10 * time.Millisecond) // ничего толком не сделал чтобы синхронить с dynmerger
	close(in1)
	close(in2)
	close(out)
	wg.Wait()
	assert.Equal(t, perInCount*2, counter)
}

// проверяем что можно динамически добавить канал на чтение
func TestDynamicMerger_Register(t *testing.T) {
	in1 := make(chan string)
	in2 := make(chan string)
	out := make(chan string)
	merger := dynmerger.New([]chan string{in1, in2}, out)

	var wg sync.WaitGroup
	counter := 0
	perInCount := 100
	wg.Add(1)
	go func(){
		defer wg.Done()
		for range out {
			counter++
		}
	}()

	in3 := make(chan string)

	for i:=0;i<perInCount;i++{

		in1 <- strconv.Itoa(i)
		in2 <- strconv.Itoa(i)
		if i == 20 {
			merger.Register(in3)
		}
		if i >= 20 {
			// не могу заставить стабильно работать канал после регистрации
			// почему-то регулярно в основной корутине несмотря на блокировки бывает так, что
			// Register добавил в slice (selects), а при этом почему то основной рабочий цикл этого не видит!!!
			// а иногда тест проходит
			in3 <- strconv.Itoa(i)
		}
	}
	close(in1)
	close(in2)
	close(in3)

	time.Sleep(150 * time.Millisecond) // ничего толком не сделал чтобы синхронить с dynmerger

	close(out)
	wg.Wait()
	assert.Equal(t, perInCount*2 + 80, counter)
}