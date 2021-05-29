package main

import "fmt"

// проверка знака числа
func main() {
	var i int
	fmt.Scan(&i)
	var message string
	switch {
	case i > 0:
		message = "Число положительное"
	case i < 0:
		message = "Число отрицательное"
	default /* i == 0 */ :
		message = "Ноль"
	}
	fmt.Print(message)
}
