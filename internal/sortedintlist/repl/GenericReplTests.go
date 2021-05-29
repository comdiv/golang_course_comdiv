package repl

import (
	"fmt"
	"github.com/comdiv/golang_course_comdiv/internal/sortedintlist"
	"io/fs"
	"os"
	"testing"
)

func GenericTestForReplExecute(name string, impl sortedintlist.IIntListMutable, t *testing.T) {
	err := os.Mkdir("./tmp", fs.ModeDir)
	if err != nil {
		panic(fmt.Errorf("error create tmp dir: %v", err))
	}
	out, err := os.Create("./tmp/" + name + ".out.txt")
	if err != nil {
		t.Fail()
	}
	in, err := os.Create("./tmp/" + name + ".in.txt")
	if err != nil {
		t.Fail()
	}
	_, err = in.WriteString("1\n2\n2\n2\n3\n4\n-3\ncount\nsize\nall\nunique\nmin\nmax\nexit\n")
	if err != nil {
		panic(fmt.Errorf("cannot write command text file: %v", err))
	}
	in.Close()

	in, err = os.Open(in.Name())
	if err != nil {
		t.Fail()
	}

	repl := NewCustom(in, out, impl)
	repl.Execute()

	result, err := os.ReadFile(out.Name())
	if err != nil {
		t.Fail()
	}
	text := string(result)
	expected := "5\n3\n[1 2 2 2 4]\n[1 2 4]\n1\n4\n"
	if text != expected {
		t.Errorf("Expected `%s` but was `%s`", expected, text)
	}
}

func GenericTestForReplCommand(name string, impl sortedintlist.IIntListMutable, t *testing.T) {
	err := os.Mkdir("./tmp", fs.ModeDir)
	if err != nil {
		panic(fmt.Errorf("error create tmp dir: %v", err))
	}
	out, err := os.Create("./tmp/" + name + ".out.txt")
	if err != nil {
		t.Fail()
	}
	repl := NewCustom(nil, out, impl)
	_ = repl.ExecuteCommand("1")
	_ = repl.ExecuteCommand("2")
	_ = repl.ExecuteCommand("2")
	_ = repl.ExecuteCommand("2")
	_ = repl.ExecuteCommand("3")
	_ = repl.ExecuteCommand("4")
	_ = repl.ExecuteCommand("-3")
	_ = repl.ExecuteCommand("count")
	_ = repl.ExecuteCommand("size")
	_ = repl.ExecuteCommand("all")
	_ = repl.ExecuteCommand("unique")
	_ = repl.ExecuteCommand("min")
	_ = repl.ExecuteCommand("max")

	result, err := os.ReadFile(out.Name())
	if err != nil {
		t.Fail()
	}
	text := string(result)
	expected := "5\n3\n[1 2 2 2 4]\n[1 2 4]\n1\n4\n"
	if text != expected {
		t.Errorf("Expected `%s` but was `%s`", expected, text)
	}
}
