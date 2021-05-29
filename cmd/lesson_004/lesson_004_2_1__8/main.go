package main

import "fmt"

// типа функция голосования
func vote(x int, y int, z int) int {
	if x+y+z > 1 {
		return 1
	} else {
		return 0
	}
}

func main() {
	fmt.Print(vote(0, 0, 1))
	fmt.Print(vote(0, 1, 1))
}
