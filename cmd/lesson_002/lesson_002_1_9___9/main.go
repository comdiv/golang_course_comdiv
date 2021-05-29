package main

import "fmt"

// определение високосного года
func main() {
	var y uint
	fmt.Scan(&y)
	if (y%400 == 0) || (y%4 == 0 && y%100 != 0) {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
