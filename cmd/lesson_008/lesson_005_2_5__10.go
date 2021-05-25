//  На вход дается строка, из нее нужно сделать другую строку, оставив только нечетные символы (считая с нуля)
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.Trim(s, "\r\n")
	srunes := []rune(s)
	tl := len(srunes) / 2
	trunes := make([]rune, 0, tl)
	for i := 1; i < len(srunes); i += 2 {
		trunes = append(trunes, srunes[i])
	}
	fmt.Println(string(trunes))
}
