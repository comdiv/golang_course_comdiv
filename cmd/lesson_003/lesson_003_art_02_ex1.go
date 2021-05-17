package main

import "fmt"

// просто вариант, переделанный на передачу параметра по ссылке
func (p *path) ToUpperPtr() {
	for i, b := range *p {
		if 'a' <= b && b <= 'z' {
			(*p)[i] = b + 'A' - 'a'
		}
	}
}

func main() {
	pathName := path("/usr/bin/tso")
	pathName.ToUpperPtr()
	fmt.Printf("%s\n", pathName)
}
