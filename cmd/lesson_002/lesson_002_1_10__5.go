package main

import "fmt"

// количество максимальных чисел в ряду, завершающегося на 0
func main() {
	var current, max, maxcount uint
	for {
		fmt.Scan(&current)
		if current == 0 {
			break
		}
		if current >= max {
			if current > max {
				maxcount = 0
				max = current
			}
			maxcount++
		}
	}
	fmt.Print(maxcount)
}
