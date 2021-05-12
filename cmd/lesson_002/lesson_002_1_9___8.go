package main

import "fmt"

// определение счастливого билетика
func main() {
	var n int
	fmt.Scan(&n)
	var l, r = n / 1000, n % 1000
	if dsum(l) == dsum(r) {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}

// сумма цифр в числе
func dsum(i int) int {
	var s int
	for c := i; c > 0; s, c = s+c%10, c/10 {
	}
	return s
}
