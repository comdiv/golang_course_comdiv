package tools

import (
	"fmt"
	"os"
)

func GenerateTestJson() {
	f, err := os.Create("./internal/investigatejson/testdata/sample.json")
	if err != nil {
		panic(fmt.Errorf("error creating sample file: %v", err))
	}
	defer f.Close()
	f.WriteString("[\n")
	delimiter := ",\n"
	for i := 1; i <= 10000; i++ {
		if i == 10000 {
			delimiter = ""
		}
		f.WriteString(fmt.Sprintf(`{"number":%d, "text": "it is a text #%d", "iseven" : %v, "substruct" : {"name" :"name %d", "value" : "my value %d"} }%s`, i, i, i%2 == 0, i, i, delimiter))
	}
	f.WriteString("\n]")
}
