package main

import "fmt"

// просто какое-то чтение в слайс с выводом фиксированного элемента
func main() {
	var size int
	fmt.Scan(&size)
	var result = make([]int, size, size)
	var value int
	for i := 0; i < size; i++ {
		fmt.Scan(&value)
		result[i] = value
	}
	fmt.Print(result[3])
}
