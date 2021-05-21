package SortedLinkedList

import (
	"fmt"
	"os"
	"strconv"
)

type LinkedListRepl struct {
	in   *os.File
	out  *os.File
	list *SortedLinkedList
}

func NewLinkedListRepl(in *os.File, out *os.File) *LinkedListRepl {
	if in == nil {
		in = os.Stdin
	}
	if out == nil {
		out = os.Stdout
	}
	return &LinkedListRepl{list: NewSortedLinkedList(), in: in, out: out}
}

func (this *LinkedListRepl) PrintHelp() {
	fmt.Println("Sorted list holder application")
	fmt.Println("Commands:")
	fmt.Println("any positive int ( 10 )  - add it to list")
	fmt.Println("any negative int ( -10)  - remove it counterpart from list")
	fmt.Println("size  - prints list size (unique value count)")
	fmt.Println("count - prints list size (all value count)")
	fmt.Println("all - prints all values (with duplicates)")
	fmt.Println("unique - prints only unique values")
}

func (this *LinkedListRepl) Execute() {
	var cmd string
	for {
		fmt.Fscan(this.in, &cmd)
		exit := false
		switch cmd {
		case "exit":
			exit = true
		case "help":
			this.PrintHelp()
		default:
			this.ExecuteCommand(cmd)
		}
		if exit {
			break
		}
	}
}

func (this *LinkedListRepl) ExecuteCommand(cmd string) {
	switch cmd {
	case "all":
		fmt.Fprintf(this.out, "%v\n", this.list.GetAll())
	case "unique":
		fmt.Fprintf(this.out, "%v\n", this.list.GetDistinct())
	case "count":
		fmt.Fprintf(this.out, "%v\n", this.list.ItemCount())
	case "size":
		fmt.Fprintf(this.out, "%v\n", this.list.IndexSize())
	default:
		ival, err := strconv.Atoi(cmd)
		if err != nil {
			fmt.Fprintf(this.out, "Error command: %v (%v)\n", cmd, err)
		} else {
			if ival >= 0 {
				this.list.Insert(ival)
			} else {
				this.list.Delete(-ival)
			}
		}
	}
}
