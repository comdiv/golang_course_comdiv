//  Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования.
// Длина пароля должна быть не менее 5 символов, он должен содержать только цифры и буквы латинского алфавита.
// На вход подается строка-пароль. Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s = strings.Trim(s, "\r\n")
	sr := []rune(s)
	valid := true
	if len(sr) < 5 {
		valid = false
	} else {
		for _, r := range sr {
			if !(unicode.IsDigit(r) || unicode.Is(unicode.Latin, r)) {
				valid = false
				break
			}
		}
	}
	if valid {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}
