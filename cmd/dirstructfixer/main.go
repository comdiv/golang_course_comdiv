package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	for _, dir := range getLessonDirs() {
		for _, f := range getMainFiles(dir) {
			fmt.Println(f)
			changeFileToDir(dir, f)
		}
	}
}

func changeFileToDir(dir, file string) {
	dirname := strings.Replace(file, ".go", "", 1)
	dirpath := filepath.Join(dir, dirname)
	err := os.Mkdir(dirpath, os.ModeDir)
	if err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			log.Fatalf("Error create dir: %v", err)
		}
	}
	err = os.Rename(filepath.Join(dir, file), filepath.Join(dirpath, "main.go"))
	if err != nil {
		log.Fatalf("Error move  file to: %v", err)
	}

}

func getMainFiles(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatalf("Error read dir: %v", err)
	}
	res := make([]string, 0)
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".go") && !f.IsDir() {
			name := filepath.Join(dir, f.Name())
			fread, err := os.Open(name)
			defer fread.Close()
			if err != nil {
				log.Fatalf("Error read file: %v", err)
			}
			reader := bufio.NewReader(fread)

			for {
				s, err := reader.ReadString('\n')
				if strings.Contains(s, "func main(") {
					res = append(res, f.Name())
				}
				if err != nil {
					break
				}
			}
		}
	}
	return res
}

func getLessonDirs() []string {
	files, err := ioutil.ReadDir("./cmd")
	if err != nil {
		log.Fatalf("Error read dir: %v", err)
	}
	res := make([]string, 0)
	for _, f := range files {
		if strings.HasPrefix(f.Name(), "lesson_") && f.IsDir() {
			res = append(res, filepath.Join("./cmd", f.Name()))
		}
	}
	return res
}
