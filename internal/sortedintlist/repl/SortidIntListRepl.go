package repl

import (
	"fmt"
	"github.com/comdiv/golang_course_comdiv/internal/sortedintlist"
	"github.com/comdiv/golang_course_comdiv/internal/sortedintlist/linked"
	"os"
	"strconv"
	"strings"
)

type SortedIntListRepl struct {
	in   *os.File
	out  *os.File
	list sortedintlist.ISortedIntList
}

func NewLinkedListRepl() *SortedIntListRepl {
	return NewLinkedListReplF(nil, nil) //default files
}

func NewLinkedListReplF(in *os.File, out *os.File) *SortedIntListRepl {
	return NewSortedListReplF(in, out, linked.NewSortedLinkedList())
}

func NewSortedListRepl(list sortedintlist.ISortedIntList) *SortedIntListRepl {
	return NewSortedListReplF(nil, nil, list)
}

func NewSortedListReplF(in *os.File, out *os.File, list sortedintlist.ISortedIntList) *SortedIntListRepl {
	if in == nil {
		in = os.Stdin
	}
	if out == nil {
		out = os.Stdout
	}
	return &SortedIntListRepl{list: list, in: in, out: out}
}

func (r *SortedIntListRepl) PrintHelp() {
	fmt.Println("Sorted list holder application")
	fmt.Println("Commands:")
	fmt.Println("any positive int ( 10 )  - add it to list")
	fmt.Println("any negative int ( -10)  - remove it counterpart from list (single)")
	fmt.Println("double minus negative int ( --10)  - remove it counterpart from list (all)")
	fmt.Println("size  - prints list size (unique value count)")
	fmt.Println("count - prints list size (all value count)")
	fmt.Println("all - prints all values (with duplicates)")
	fmt.Println("unique - prints only unique values")
}

func (r *SortedIntListRepl) Execute() {
	var cmd string
	for {
		fmt.Fscan(r.in, &cmd)
		exit := false
		switch cmd {
		case "exit":
			exit = true
		case "help":
			r.PrintHelp()
		default:
			r.ExecuteCommand(cmd)
		}
		if exit {
			break
		}
	}
}

func (r *SortedIntListRepl) ExecuteCommand(cmd string) error {
	switch cmd {
	case "all":
		fmt.Fprintf(r.out, "%v\n", r.list.GetAll())
	case "unique":
		fmt.Fprintf(r.out, "%v\n", r.list.GetUnique())
	case "count":
		fmt.Fprintf(r.out, "%v\n", r.list.Size())
	case "size":
		fmt.Fprintf(r.out, "%v\n", r.list.UniqueSize())
	default:
		var removeAll bool = false
		var numberPart string = cmd
		if strings.HasPrefix(cmd, "--") {
			removeAll = true
			numberPart = cmd[1:]
		}
		ival, err := strconv.Atoi(numberPart)
		if err != nil {
			fmt.Fprintf(r.out, "Error command: %v (%v)\n", cmd, err)
			return err
		} else {
			if ival >= 0 {
				r.list.Insert(ival)
			} else {
				r.list.Delete(-ival, removeAll)
			}
		}
	}
	return nil
}
