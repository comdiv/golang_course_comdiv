package main

import "fmt"

// сумма всех числе в диапазоне [from..to]
func main() {
	var from, to, sum uint
	fmt.Scan(&from, &to)
	for c := from; c <= to; sum, c = sum+c, c+1 {
	}
	fmt.Print(sum)
}
