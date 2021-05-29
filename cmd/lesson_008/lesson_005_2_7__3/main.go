// Дана строка, содержащая только английские буквы (большие и маленькие).
// Добавить символ ‘*’ (звездочка) между буквами (перед первой буквой и после последней символ ‘*’ добавлять не нужно).
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.Trim(text, "\r\n")
	result := make([]byte, 0, len(text)*2-1)
	for i := 0; i < len(text); i++ {
		result = append(result, text[i])
		if i < len(text)-1 {
			result = append(result, byte('*'))
		}
	}
	fmt.Println(string(result))
}
