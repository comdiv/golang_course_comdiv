package main

import "fmt"

// минимальное число из четырех с консоли
func minimumFromFour() int {
	// инициализируем максимальным возможным значением для int32
	var min int = 1<<31 - 1
	for i := 0; i < 4; i++ {
		var current int
		fmt.Scan(&current)
		if current < min {
			min = current
		}
	}
	return min
}

func main() {
	fmt.Print(minimumFromFour())
}
