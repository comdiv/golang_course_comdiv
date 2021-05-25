//  На вход подаются a и b - катеты прямоугольного треугольника. Нужно найти длину гипотенузы
package main

import (
	"fmt"
	"math"
)

func main() {
	var c1, c2 float64
	fmt.Scan(&c1, &c2)
	h := math.Sqrt(math.Pow(c1, 2) + math.Pow(c2, 2))
	hi := int(h)
	fmt.Println(hi)
}
