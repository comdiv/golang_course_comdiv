package main

import "fmt"

// сделал с хвостовой рекурсией, должен по идее оптимизировать
func fibonacci(n int) int {
	return fibonacciTailed(1, 0, 1, n)
}

func fibonacciTailed(result int, last int, count int, terminal int) int {
	if count == terminal {
		return result
	} else {
		return fibonacciTailed(result+last, result, count+1, terminal)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Print(fibonacci(n))
}
