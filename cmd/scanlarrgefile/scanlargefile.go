package main

import (
	"bufio"
	"fmt"
	"os"
)

const DEFAULT_LINE_SIZE_LIMIT = 5*1024*1024 + 1 // 5mb + \n

func main() {
	// установим предельный рамер буфера для строки
	file, err := os.Open("./tmp/large.txt")
	if nil != err {
		fmt.Printf("Error %v", err)
		return
	}
	defer file.Close()
	buffered := bufio.NewReaderSize(file, DEFAULT_LINE_SIZE_LIMIT)
	lineNumber := 0
	lastWasPrefix := false
	longLength := 0
	for {
		line, prefix, err := buffered.ReadLine()
		if nil != err {
			break
		}
		if prefix {
			lastWasPrefix = true
			longLength += DEFAULT_LINE_SIZE_LIMIT
			continue
		}

		lineNumber++

		if lastWasPrefix {
			lastWasPrefix = false
			fmt.Printf("Line #%d - SKIPPED - too long - %d bytes!\n", lineNumber, longLength+len(line))
			longLength = 0
			continue
		}

		fmt.Printf("Line #%d - PROCESSED - size - %d bytes\n", lineNumber, len(line))

	}
}
