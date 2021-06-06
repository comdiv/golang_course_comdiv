package main

import (
	"archive/zip"
	"encoding/csv"
	"fmt"
	"strconv"
)

func main() {
	zf, err := zip.OpenReader("./cmd/finddataincsv/testdata/task.zip")
	if err != nil {
		panic(err)
	}
	defer zf.Close()

	var result int
	var found bool

	for _, file := range zf.File {
		file.Open()
		if file.FileInfo().IsDir() {
			continue
		}
		result, err = processFile(file)
		if err == nil {
			found = true
			break
		}
	}

	if !found {
		panic("Cannot find CSV with data")
	}

	fmt.Println(result)
}

func processFile(file *zip.File) (int, error) {
	reader, err := file.Open()
	if err != nil {
		panic(err)
	}
	defer reader.Close()
	csvReader := csv.NewReader(reader)
	var result int
	for i := 0; i < 5; i++ {
		line, err := csvReader.Read()
		if err != nil {
			return 0, err
		}
		if i == 4 {
			result, err = strconv.Atoi(line[2])
			if err != nil {
				return 0, err
			}
		}
	}
	return result, nil
}
