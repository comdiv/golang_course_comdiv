package repl

import (
	"github.com/comdiv/golang_course_comdiv/internal/sortedintlist/linked"
	"testing"
)

func TestLinkedListRepl_ExecuteCommand(t *testing.T) {
	GenericTestForReplCommand("repl_linked_command", linked.NewSortedLinkedList(), t)
}

func TestLinkedListRepl_Execute(t *testing.T) {
	GenericTestForReplExecute("repl_linked_execute", linked.NewSortedLinkedList(), t)
}
