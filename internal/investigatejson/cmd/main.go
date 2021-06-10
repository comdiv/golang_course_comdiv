package main

import (
	"flag"
	"github.com/comdiv/golang_course_comdiv/internal/investigatejson/tools"
)

func main() {
	generate := flag.Bool("generate", false, "Regenerate test file")
	flag.Parse()
	if *generate {
		tools.GenerateTestJson()
	}
}
