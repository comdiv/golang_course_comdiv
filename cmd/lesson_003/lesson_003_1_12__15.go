package main

import "fmt"

// вывод четных элементов массива
func main() {
	var size int
	fmt.Scan(&size)
	values := make([]int, size, size)
	for i := range values {
		fmt.Scan(&values[i])
	}
	for i := 0; i < len(values); i = i + 2 {
		if i != 0 {
			fmt.Print(" ")
		}
		fmt.Print(values[i])
	}
}
