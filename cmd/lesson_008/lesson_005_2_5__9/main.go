// найти индекс первого вхождения одной строки в другую
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	smain, _ := reader.ReadString('\n')
	sfind, _ := reader.ReadString('\n')
	sfind = strings.Trim(sfind, "\r\n")
	fmt.Println(strings.Index(smain, sfind))
}
