package main

import "fmt"

// сумма двузначных чисел из списка указанной длины, кратных 8
func main() {
	var length, current, sum int
	fmt.Scan(&length)
	for i := 0; i < length; i++ {
		fmt.Scan(&current)
		if current >= 10 && current <= 99 && current%8 == 0 {
			sum += current
		}
	}
	fmt.Print(sum)
}
