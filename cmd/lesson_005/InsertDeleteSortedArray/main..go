package main

import (
	"github.com/comdiv/golang_course_comdiv/cmd/lesson_005/InsertDeleteSortedArray/SortedLinkedList"
	"os"
)

func main() {
	repl := SortedLinkedList.NewLinkedListRepl(os.Stdin, os.Stdout)
	repl.PrintHelp()
	repl.Execute()
}
