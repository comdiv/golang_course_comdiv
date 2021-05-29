package main

/*
// типа более сложный пример для ToUpper с
// использованием UTF-8 - сделал через указатели,
// чтобы проще было подменять байты с рунами
func (p *path) ToUpperUTF() {
	utfs := []rune(string(*p))
	for i, b := range utfs {
		utfs[i] = unicode.ToUpper(b)
	}
	*p = []byte(string(utfs))
}

func main() {
	pathName := path("/usr/bin/привет")
	pathName.ToUpperUTF()
	fmt.Printf("%s\n", pathName)
}
*/
