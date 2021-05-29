package main

import "fmt"

// печатать числа [10..100] и если придет > 100 - остановиться
// задание было на break, но я не вижу в нем тут особого смысла
func main() {
	var c int
	for fmt.Scan(&c); c <= 100; fmt.Scan(&c) {
		if c >= 10 {
			fmt.Println(c)
		}
	}
}
