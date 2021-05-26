// проверка что первая буква строки большая, а последний символ - точка
package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	symbols := []rune(text)
	if len(symbols) >= 2 && unicode.IsUpper(symbols[0]) && '.' == symbols[len(symbols)-1] {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}
