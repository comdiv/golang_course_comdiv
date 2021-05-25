package repl

import (
	"fmt"
	"github.com/comdiv/golang_course_comdiv/internal/sortedintlist"
	"github.com/comdiv/golang_course_comdiv/internal/sortedintlist/linked"
	"github.com/comdiv/golang_course_comdiv/internal/sortedintlist/slices"
	"io"
	"os"
	"strconv"
	"strings"
)

type SortedIntListRepl struct {
	in     io.Reader
	out    io.Writer
	list   sortedintlist.IIntListMutable
	set    sortedintlist.IIntSet
	minmax sortedintlist.IIntMinMax
}

func NewLinkedListRepl() *SortedIntListRepl {
	return NewLinkedListReplF(nil, nil) //default files
}

func NewLinkedListReplF(in io.Reader, out io.Writer) *SortedIntListRepl {
	return NewSortedListReplF(in, out, linked.NewSortedLinkedList())
}

func NewSlicedListReplF(in io.Reader, out io.Writer) *SortedIntListRepl {
	return NewSortedListReplF(in, out, slices.NewSortedIntListSliced())
}

func NewSortedListRepl(list sortedintlist.IIntListMutable) *SortedIntListRepl {
	return NewSortedListReplF(nil, nil, list)
}

func NewSortedListReplF(in io.Reader, out io.Writer, list sortedintlist.IIntListMutable) *SortedIntListRepl {
	if in == nil {
		in = os.Stdin
	}
	if out == nil {
		out = os.Stdout
	}
	return &SortedIntListRepl{
		list:   list,
		in:     in,
		out:    out,
		set:    list.(sortedintlist.IIntSet),
		minmax: list.(sortedintlist.IIntMinMax),
	}
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
	fmt.Println("min - prints min value")
	fmt.Println("max - prints max value")
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
		if nil == r.set {
			fmt.Println("Not supported")
		} else {
			fmt.Fprintf(r.out, "%v\n", r.set.GetUnique())
		}
	case "count":
		fmt.Fprintf(r.out, "%v\n", r.list.Size())
	case "size":
		if nil == r.set {
			fmt.Println("Not supported")
		} else {
			fmt.Fprintf(r.out, "%v\n", r.set.UniqueSize())
		}
	case "min":
		if nil == r.minmax {
			fmt.Println("Not supported")
		} else {
			min, err := r.minmax.GetMin()
			if nil != err {
				fmt.Printf("error: %v", err)
			} else {
				fmt.Fprintf(r.out, "%v\n", min)
			}
		}
	case "max":
		if nil == r.minmax {
			fmt.Println("Not supported")
		} else {
			max, err := r.minmax.GetMax()
			if nil != err {
				fmt.Printf("error: %v", err)
			} else {
				fmt.Fprintf(r.out, "%v\n", max)
			}
		}
	default:
		var removeAll = false
		var numberPart = cmd
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
