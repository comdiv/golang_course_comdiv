package main

import (
	"bufio"
	"io/fs"
	"os"
	"testing"
)

func TestGenerateLargeFile(t *testing.T) {
	os.Mkdir("./tmp", fs.ModeDir)
	f, err := os.Create("./tmp/large.txt")
	if err != nil {
		t.Fail()
	}
	defer f.Close()
	writer := bufio.NewWriter(f)
	defer writer.Flush()
	buf := make([]byte, 10*1024*1024) // large buff
	for i := 0; i < cap(buf); i++ {
		buf[i] = byte('0') + byte(i%10) // 01234567890123456789
	}
	line := func(size int) {
		writer.Write(buf[0:size])
		writer.WriteByte(byte('\n'))
	}
	line(10)
	line(20)
	line(1024)
	line(5 * 1024 * 1024)
	line(5*1024*1024 - 1) // кстати видно, что \n должен поместиться в буфер!
	line(6 * 1024 * 1024)
	line(1000)
}
