package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	fmt.Print(x / 10 % 10)
}
