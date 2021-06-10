package investigatejson

import "fmt"
import "github.com/buger/jsonparser"
import jsoniter "github.com/json-iterator/go"
func PrintIHaveDependencies() {
	fmt.Printf("from json parser jsonparser.String: %v\n", jsonparser.String)
	fmt.Printf("from jsoniter jsoniter.StringValue: %v\n", jsoniter.StringValue)
}