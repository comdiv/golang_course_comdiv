package testdata_test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

var testDataFileName = "../testdata/main_test_text.txt"

func loadTestData() []byte {

	data, err := ioutil.ReadFile(testDataFileName)
	if err != nil {
		panic(fmt.Errorf("file load error %s: %v", testDataFileName, err))
	}
	if len(data) < 12000 {
		panic(fmt.Errorf("too small file %s : %d", testDataFileName, len(data)))
	}
	return data
}

var testData = loadTestData()

func TestDataReader() io.Reader {
	return bytes.NewReader(testData)
}
