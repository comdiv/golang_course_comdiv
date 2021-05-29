package main

import "fmt"

// первое число из [1..n], кратное c и не кратное d
func main() {
	var n, c, d int
	fmt.Scan(&n, &c, &d)
	for i := 1; i <= n; i++ {
		if i%c == 0 && i%d != 0 {
			fmt.Print(i)
			break
		}
	}
}
