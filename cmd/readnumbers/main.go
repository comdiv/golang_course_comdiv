package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
Поэтапный поиск данных
Данная задача в основном ориентирована на изучение типа bufio.Reader, поскольку этот тип позволяет считывать данные постепенно.

В тестовом файле, который вы можете скачать из нашего репозитория на github.com, содержится длинный ряд чисел, разделенных символом ";". Требуется найти, на какой позиции находится число 0 и указать её в качестве ответа. Требуется вывести именно позицию числа, а не индекс (то-есть порядковый номер, нумерация с 1).

Например:  12;234;6;0;78 :
Правильный ответ будет порядковый номер числа: 4
*/

func main() {
	file, err := os.Open("./cmd/readnumbers/task.data")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	idx := 1
	found := false
	for line, err := reader.ReadString(';'); err == nil; line, err = reader.ReadString(';') {
		num, numerror := strconv.Atoi(line[:len(line)-1]) // обходим кейс ;00000000;
		if numerror != nil {
			panic(numerror)
		}
		if num == 0 {
			found = true
			break
		}
		idx++
	}
	if !found {
		panic("Not found!")
	}
	fmt.Println(idx)
}
