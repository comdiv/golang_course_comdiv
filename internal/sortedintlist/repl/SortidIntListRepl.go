package repl

import (
	"fmt"
	"github.com/comdiv/golang_course_comdiv/internal/sortedintlist"
	"github.com/comdiv/golang_course_comdiv/internal/sortedintlist/slices"
	"io"
	"strconv"
	"strings"
)

type SortedIntListRepl struct {
	in   io.Reader
	out  io.Writer
	list sortedintlist.IIntListMutable
}

func New(in io.Reader, out io.Writer) *SortedIntListRepl {
	return NewCustom(in, out, slices.New())
}

func NewCustom(in io.Reader, out io.Writer, list sortedintlist.IIntListMutable) *SortedIntListRepl {
	return &SortedIntListRepl{
		list: list,
		in:   in,
		out:  out,
	}
}

func (r *SortedIntListRepl) asSet() sortedintlist.IIntSet {
	return r.list.(sortedintlist.IIntSet)
}
func (r *SortedIntListRepl) asMinMax() sortedintlist.IIntMinMax {
	return r.list.(sortedintlist.IIntMinMax)
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
	if nil == r.in {
		return
	}
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
			err := r.ExecuteCommand(cmd)
			if err != nil {
				fmt.Printf("Error in repl: %v\n", err)
				break
			}
		}
		if exit {
			break
		}
	}
}

func (r *SortedIntListRepl) ExecuteCommand(cmd string) error {
	print := func(format string, a ...interface{}) {
		if nil != r.out {
			fmt.Fprintf(r.out, format, a...)
		}
	}
	switch cmd {
	case "all":
		print("%v\n", r.list.GetAll())
	case "unique":
		if nil == r.asSet() {
			print("Is not a set - `unique` - not supported\n")
			break
		}
		print("%v\n", r.asSet().GetUnique())
	case "count":
		print("%v\n", r.list.Size())
	case "size":
		if nil == r.asSet() {
			print("Is not a set - `size` - not supported\n")
			break
		}
		print("%v\n", r.asSet().UniqueSize())
	case "min":
		if nil == r.asMinMax() {
			print("Is not a IIntMinMax - `min` - not supported\n")
			break
		}
		min, err := r.asMinMax().GetMin()
		if nil != err {
			print("error: %v\n", err)
			break
		}
		print("%v\n", min)
	case "max":
		if nil == r.asMinMax() {
			print("Is not a IIntMinMax - `max` - not supported\n")
			break
		}
		max, err := r.asMinMax().GetMax()
		if nil != err {
			print("error: %v\n", err)
			break
		}
		print("%v\n", max)

	default:
		removeAll := false
		numberPart := cmd
		if strings.HasPrefix(cmd, "--") {
			removeAll = true
			numberPart = cmd[1:]
		}
		ival, err := strconv.Atoi(numberPart)
		if err != nil {
			print("Error command: %v (%v)\n", cmd, err)
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
