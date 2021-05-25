//  Дается строка. Нужно удалить все символы, которые встречаются более одного раза и вывести получившуюся строку
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
	trunes := make([]rune, 0, len(s))
	for _, r := range s {
		if strings.Count(s, string(r)) == 1 {
			trunes = append(trunes, r)
		}
	}
	fmt.Println(string(trunes))
}
