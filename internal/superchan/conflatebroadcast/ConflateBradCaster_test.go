package dynbroadcaster_test

import (
	"context"
	dynbroadcaster "github.com/comdiv/golang_course_comdiv/internal/superchan/conflatebroadcast"
	"github.com/stretchr/testify/assert"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestConflateBroadCaster_AllowWriteEvenIfNoListeners(t *testing.T) {
	// у нас два источника и две цели
	src_1 := make(chan string)
	src_2 := make(chan string)

	// и у нас есть собственно конфлейтер
	conflater := dynbroadcaster.New(context.TODO())
	conflater.StartAsync(context.TODO())
	// заведем в него источники и слушатели
	conflater.Publish(context.TODO(), src_1)
	conflater.Publish(context.TODO(), src_2)

	// и даже добавим коналы слушатели, которые НИКТО не будет читать
	target_1 := make(chan func() string)
	target_2 := make(chan func() string)
	conflater.Listen(context.TODO(), target_1)
	conflater.Listen(context.TODO(), target_2)

	var wg sync.WaitGroup
	wg.Add(2)
	// первый продюсер раздает четные числа каждые 10ms
	go func() {
		defer wg.Done()
		for i:=0;i<100; i+=2 {
			src_1 <- strconv.Itoa(i)
			time.Sleep(10 * time.Millisecond)
		}
	}()
	// второй продюсер раздает нечетные числа каждые 20ms
	go func() {
		defer wg.Done()
		for i:=1;i<100; i+=2 {
			src_2 <- strconv.Itoa(i)
			time.Sleep(20 * time.Millisecond)
		}
	}()

	// по идее все спокойно завершится без единого слушателя
	wg.Wait()
}

func TestConflateBroadCaster_LazyListeners(t *testing.T) {
	// у нас два источника и две цели
	src_1 := make(chan string)
	src_2 := make(chan string)

	// и у нас есть собственно конфлейтер
	conflater := dynbroadcaster.New(context.TODO())
	conflater.StartAsync(context.TODO())
	// заведем в него источники и слушатели
	conflater.Publish(context.TODO(), src_1)
	conflater.Publish(context.TODO(), src_2)

	//  добавим коналы слушатели, и на этот раз обвесим их слушателями
	target_1 := make(chan func() string)
	target_2 := make(chan func() string)
	conflater.Listen(context.TODO(), target_1)
	conflater.Listen(context.TODO(), target_2)


	// вот 2 получателя из своих каналов, работают медленно, с задержками
	counter_1 := 0
	counter_2 := 0

	var wg sync.WaitGroup
	wg.Add(4)

	// Заодно убедимся что ВСЕ слушатели получают одни и те же значения
	// в смысле что могут получить одинаковые, они не эксклюзивны, например оба должны получить сообщение 99
	// так как нечетные шлются в 2 раза медленнее и это последнее число

	go func() {
		defer wg.Done()
		for v := range target_1 {
			val := v()
			counter_1++
			if val == "99" {
				return
			}
			time.Sleep(30 * time.Millisecond)
		}
	}()
	go func() {
		defer wg.Done()
		for v := range target_2 {
			val := v()

			// заодно проверяем, что функция - с мемоизацией, значение не меняется
			if (counter_2 > 3) {
				time.Sleep(30 * time.Millisecond) // точно поменялся current
				// но вызов v()  нет
				assert.Equal(t, val, v())
			}

			counter_2++
			if val == "99" {
				return
			}
			time.Sleep(40 * time.Millisecond)
		}
	}()



	// первый продюсер раздает четные числа каждые 10ms
	go func() {
		defer wg.Done()
		for i:=0;i<100; i+=2 {
			src_1 <- strconv.Itoa(i)
			time.Sleep(10 * time.Millisecond)
		}
	}()
	// второй продюсер раздает нечетные числа каждые 20ms
	go func() {
		defer wg.Done()
		for i:=1;i<100; i+=2 {
			src_2 <- strconv.Itoa(i)
			time.Sleep(20 * time.Millisecond)
		}
	}()

	wg.Wait()

	// убеждаемся что оба листенера получили значения и что у более медленного их меньше
	assert.Greater(t, counter_1, 5)
	assert.Greater(t, counter_2, 5)
	assert.Greater(t, counter_1, counter_2)
}
