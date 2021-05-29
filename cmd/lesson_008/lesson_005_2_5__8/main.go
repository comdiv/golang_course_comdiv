// проверка строки, что она палиндром, не было разъяснения как учитывать пробелы, решил их удалять
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
	filteredSymbols := make([]rune, 0, len(symbols))
	for _, r := range symbols {
		if unicode.IsLetter(r) {
			filteredSymbols = append(filteredSymbols, unicode.ToLower(r))
		}
	}
	isPalindrome := true
	for i := 0; i < len(filteredSymbols)/2; i++ {
		if filteredSymbols[i] != filteredSymbols[len(filteredSymbols)-1-i] {
			isPalindrome = false
			break
		}
	}
	if isPalindrome {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}
