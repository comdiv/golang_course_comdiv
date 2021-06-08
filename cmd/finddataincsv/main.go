package main

import (
	"archive/zip"
	"encoding/csv"
	"fmt"
	"strconv"
)

/*
Поиск файла в заданном формате и его обработка
Данная задача поможет вам разобраться в пакете encoding/csv и path/filepath, хотя для решения может быть использован также пакет archive/zip (поскольку файл с заданием предоставляется именно в этом формате).

В тестовом архиве, который вы можете скачать из нашего репозитория на github.com, содержится набор папок и файлов. Один из этих файлов является файлом с данными в формате CSV, прочие же файлы структурированных данных не содержат.

Требуется найти и прочитать этот единственный файл со структурированными данными (это таблица 10х10, разделителем является запятая), а в качестве ответа необходимо указать число, находящееся на 5 строке и 3 позиции (индексы 4 и 2 соответственно).
*/

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
