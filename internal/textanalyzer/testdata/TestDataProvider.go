package testdata_test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

var testDataFileName = "../testdata/main_test_text.txt"
var testDataJsonFileName = "../testdata/main_test_json.json"
var testDataLargeJsonFileName = "../testdata/large_test_json.json"

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

func loadTestJsonData() []byte {
	data, err := ioutil.ReadFile(testDataJsonFileName)
	if err != nil {
		panic(fmt.Errorf("file load error %s: %v", testDataFileName, err))
	}
	if len(data) < 12000 {
		panic(fmt.Errorf("too small file %s : %d", testDataFileName, len(data)))
	}
	return data
}

func loadTestLargeJsonData() []byte {
	data, err := ioutil.ReadFile(testDataLargeJsonFileName)
	if err != nil {
		panic(fmt.Errorf("file load error %s: %v", testDataFileName, err))
	}
	if len(data) < 12000*4 {
		panic(fmt.Errorf("too small file %s : %d", testDataFileName, len(data)))
	}
	return data
}


var testData = loadTestData()
var testJsonData = loadTestJsonData()
var testLargeJsonData = loadTestLargeJsonData()

func TestDataReader() io.Reader {
	return bytes.NewReader(testData)
}

func TestDataJsonReader() io.Reader {
	return bytes.NewReader(testJsonData)
}

func TestDataLargeJsonReader() io.Reader {
	return bytes.NewReader(testLargeJsonData)
}
