package main

import "fmt"

func f(text string) {
	fmt.Print(text)
}

// какая то минимальная функция
func main() {
	var inputText string
	fmt.Scan(&inputText)
	f(inputText)
}
