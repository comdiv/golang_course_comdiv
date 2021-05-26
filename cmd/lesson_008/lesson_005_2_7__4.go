// Дана строка, содержащая только десятичные цифры. Найти и вывести наибольшую цифру.
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
	var max byte = 0
	for i := 0; i < len(text); i++ {
		if text[i] > max {
			max = text[i]
		}
		if max == byte('9') {
			break
		}
	}
	fmt.Println(string(max))
}
