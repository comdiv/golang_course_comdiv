package main

import "fmt"

// вывод первой цифры в числе
func main() {
	var i, r uint
	fmt.Scan(&i)
	for r = i; r >= 10; r = r / 10 {
	}
	fmt.Println(r)
	// была еще альтернатива со строками, но показалась очень корявой
	// fmt.Println(string(fmt.Sprint(i)[0]))
}
