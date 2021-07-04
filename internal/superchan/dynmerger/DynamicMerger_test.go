package dynmerger_test

import (
	"context"
	"fmt"
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
	dynmerger.New(context.TODO(), []chan string{in1, in2}, out)
	var wg sync.WaitGroup
	counter := 0
	perInCount := 100
	wg.Add(1)
	go func() {
		defer wg.Done()
		for range out {
			counter++
		}
	}()
	for i := 0; i < perInCount; i++ {
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
	merger := dynmerger.New(context.TODO(), []chan string{in1, in2}, out)

	var wg sync.WaitGroup
	counter := 0
	perInCount := 100
	wg.Add(1)
	go func() {
		defer wg.Done()
		for range out {
			counter++
		}
	}()

	in3 := make(chan string)

	for i := 0; i < perInCount; i++ {

		in1 <- strconv.Itoa(i)
		in2 <- strconv.Itoa(i)
		if i == 20 {
			merger.Bind(context.TODO(), in3)
		}
		if i >= 20 {
			in3 <- strconv.Itoa(i)
		}
	}
	close(in1)
	close(in2)
	close(in3)

	time.Sleep(150 * time.Millisecond) // ничего толком не сделал чтобы синхронить с dynmerger

	close(out)
	wg.Wait()
	assert.Equal(t, perInCount*2+80, counter)
}

func TestDynamicMerger_Bind_And_Unbind(t *testing.T) {
	fmt.Println("Start TestDynamicMerger_Bind_And_Unbind")
	out := make(chan string)
	merger := dynmerger.New(context.TODO(), []chan string{}, out)
	in1 := make(chan string)
	hasprocessed := false
	// промотка входного канала
	go func(){
		for range in1 {
			time.Sleep(1 * time.Millisecond)
		}
	}()
	// промотка выходного канала
	go func(){
		for range out {
			if !hasprocessed {
				fmt.Println("Has some data in out!")
			}
			hasprocessed = true
		}
	}()
	job := merger.Bind(context.TODO(), in1)
	for i:=0;i<100;i++{
		in1 <- "any"
	}
	assert.True(t, hasprocessed)
	job.Finish()
	// не фает что промотка выходного канала все дообработала, увеличим гарантии!!
	time.Sleep(10 * time.Millisecond)
	hasprocessed = false
	for i:=0;i<100;i++{
		in1 <- "any"
	}
	assert.False(t, hasprocessed)
}
