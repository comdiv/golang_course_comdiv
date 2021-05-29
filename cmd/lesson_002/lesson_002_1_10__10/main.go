package main

import "fmt"

// цифры, входяшие в первое и второе число, вывести через пробел
func main() {
	var fst, sec int
	fmt.Scan(&fst, &sec)
	fsts, secs := fmt.Sprint(fst), fmt.Sprint(sec)
	requireWS := false
	for _, fchar := range fsts {
		for _, schar := range secs {
			if fchar == schar {
				if requireWS {
					fmt.Print(" ")
				}
				fmt.Print(string(fchar))
				requireWS = true
				break
			}
		}
	}
}
