// На вход подается целое число. Необходимо возвести в квадрат каждую цифру числа и вывести получившееся число.
//Например, у нас есть число 9119. Первая цифра - 9. 9 в квадрате - 81. Дальше 1. Единица в квадрате - 1. В итоге получаем 811181
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.Trim(text, "\r\n")
	var result []byte
	for i := 0; i < len(text); i++ {
		vi, _ := strconv.Atoi(string(text[i]))
		vi = vi * vi
		vis := strconv.Itoa(vi)
		result = append(result, []byte(vis)...)
	}
	fmt.Println(string(result))
}
