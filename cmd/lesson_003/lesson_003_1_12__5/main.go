package main

import "fmt"

// работа с массивом, чтение из Scan, вывод через пробел без скобок
func main() {
	workArray := [10]uint8{}

	for i := range workArray {
		fmt.Scan(&workArray[i])
	}

	for i := 0; i < 3; i++ {
		var fromIndex, toIndex uint8
		fmt.Scan(&fromIndex, &toIndex)
		workArray[fromIndex], workArray[toIndex] = workArray[toIndex], workArray[fromIndex]
	}

	for _, value := range workArray {
		fmt.Printf("%d ", value)
	}
}
