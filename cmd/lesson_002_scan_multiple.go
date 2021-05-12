package main

import "fmt"

// не очень понял как именно он сканит в несколько переменных
func main() {
	var a int
	var b int
	fmt.Println("Input a:int and b:int - or a ENTER b ENTER , or a b ENTER")
	// чудеса да и только - можно и несколько ентером , а можно и через пробел и 1 ентер
	fmt.Scan(&a, &b)
	fmt.Println(a + b)
}
