package main

import (
	"github.com/comdiv/golang_course_comdiv/internal/sortedintlist/repl"
	"os"
)

func main() {
	repl := repl.NewSlicedListReplF(os.Stdin, os.Stdout)
	repl.PrintHelp()
	repl.Execute()
}
