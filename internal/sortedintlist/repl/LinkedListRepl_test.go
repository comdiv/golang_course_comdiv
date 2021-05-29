package repl

import (
	"github.com/comdiv/golang_course_comdiv/internal/sortedintlist/linked"
	"testing"
)

func Exclude_TestLinkedListRepl_ExecuteCommand(t *testing.T) {
	GenericTestForReplCommand("repl_linked_command", linked.New(), t)
}

func Exclude_TestLinkedListRepl_Execute(t *testing.T) {
	GenericTestForReplExecute("repl_linked_execute", linked.New(), t)
}
