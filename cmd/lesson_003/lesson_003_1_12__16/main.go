package main

import "fmt"

// подсчет положительных чисел в массиве
func main() {
	var size int
	fmt.Scan(&size)
	values := make([]int, size, size)
	for i := range values {
		fmt.Scan(&values[i])
	}
	var positivesCount int
	for i := range values {
		if values[i] > 0 {
			positivesCount++
		}
	}
	fmt.Print(positivesCount)
}
