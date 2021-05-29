package main

import "fmt"

// сколько лет нужно на рост кредита до y при заданном проценте и округлением
func main() {
	var x, p, y, years int
	fmt.Scan(&x, &p, &y)
	for c := x; c < y; years, c = years+1, int(float32(c)*(1.0+float32(p)/100.0)) {
	}
	fmt.Print(years)
}
