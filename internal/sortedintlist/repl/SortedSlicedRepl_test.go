package repl

import (
	"github.com/comdiv/golang_course_comdiv/internal/sortedintlist/slices"
	"testing"
)

func Exclude_TestSlicedRepl_ExecuteCommand(t *testing.T) {
	GenericTestForReplCommand("repl_linked_command", slices.New(), t)
}

func Exclude_TestSlicedRepl_Execute(t *testing.T) {
	GenericTestForReplExecute("repl_linked_execute", slices.New(), t)
}
