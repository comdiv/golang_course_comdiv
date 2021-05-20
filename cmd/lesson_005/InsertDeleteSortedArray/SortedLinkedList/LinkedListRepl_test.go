package SortedLinkedList

import (
	"io/fs"
	"os"
	"testing"
)

func TestLinkedListRepl_ExecuteReplCommand(t *testing.T) {
	os.Mkdir("./tmp", fs.ModeDir)
	out, err := os.Create("./tmp/TestLinkedListRepl_ExecuteReplCommand.out.txt")
	if err != nil {
		t.Fail()
	}
	repl := NewLinkedListRepl(nil, out)
	repl.ExecuteCommand("1")
	repl.ExecuteCommand("2")
	repl.ExecuteCommand("2")
	repl.ExecuteCommand("2")
	repl.ExecuteCommand("3")
	repl.ExecuteCommand("4")
	repl.ExecuteCommand("-3")
	repl.ExecuteCommand("count")
	repl.ExecuteCommand("size")
	repl.ExecuteCommand("all")
	repl.ExecuteCommand("unique")

	result, err := os.ReadFile("./tmp/TestLinkedListRepl_ExecuteReplCommand.out.txt")
	if err != nil {
		t.Fail()
	}
	text := string(result)
	expected := "5\n3\n[1 2 2 2 4]\n[1 2 4]\n"
	if text != expected {
		t.Errorf("Expected `%s` but was `%s`", expected, text)
	}
}

func TestLinkedListRepl_ExecuteRepl(t *testing.T) {
	os.Mkdir("./tmp", fs.ModeDir)
	out, err := os.Create("./tmp/TestLinkedListRepl_ExecuteRepl.out.txt")
	if err != nil {
		t.Fail()
	}
	in, err := os.Create("./tmp/TestLinkedListRepl_ExecuteRepl.in.txt")
	if err != nil {
		t.Fail()
	}
	in.WriteString("1\n2\n2\n2\n3\n4\n-3\ncount\nsize\nall\nunique\nexit\n")
	in.Close()

	in, err = os.Open(in.Name())
	if err != nil {
		t.Fail()
	}

	repl := NewLinkedListRepl(in, out)
	repl.Execute()

	result, err := os.ReadFile("./tmp/TestLinkedListRepl_ExecuteRepl.out.txt")
	if err != nil {
		t.Fail()
	}
	text := string(result)
	expected := "5\n3\n[1 2 2 2 4]\n[1 2 4]\n"
	if text != expected {
		t.Errorf("Expected `%s` but was `%s`", expected, text)
	}

}
