package main

import "fmt"

func main() {
	var i float64
	fmt.Scan(&i)
	switch {
	case i <= 0:
		fmt.Printf("число %2.2f не подходит", i)
	case i > 10000:
		fmt.Printf("%e", i)
	default:
		fmt.Printf("%.4f", i*i)
	}
}
