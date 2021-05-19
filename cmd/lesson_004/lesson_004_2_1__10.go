package main

import "fmt"

// сделал с хвостовой рекурсией, должен по идее оптимизировать
/* ЗАДАНИЕ:
 * Напишите функцию sumInt, получающую переменное число аргументов типа int,
 * и возвращающую количество переданных аргументов и их сумму.
 */
func sumInt(a ...int) (count int, sum int) {
	for _, current := range a {
		count++
		sum += current
	}
	return
}

func main() {
	var n int
	var numbers []int
	for {
		fmt.Scan(&n)
		if n > 0 {
			numbers = append(numbers, n)
		} else {
			break
		}
	}
	fmt.Print(sumInt(numbers...))
}
