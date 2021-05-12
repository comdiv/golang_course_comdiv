package main

import "fmt"
import "os"

// проверка на то что все цифры в числе разные
func main() {
	var i int
	fmt.Scan(&i)
	if i < 100 || i >= 999 {
		fmt.Fprintf(os.Stderr, "Invalid 3-sign %d should be (100..999)", i)
		os.Exit(1)
	} else {
		var h, d, s = i / 100, (i % 100) / 10, i % 10
		var allDifferent = (h != d) && (h != s) && (s != d)
		if allDifferent {
			fmt.Print("YES")
		} else {
			fmt.Print("NO")
		}
	}
}
