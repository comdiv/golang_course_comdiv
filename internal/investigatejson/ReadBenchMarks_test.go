package investigatejson

import (
	"encoding/json"
	"github.com/buger/jsonparser"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"testing"
)

const fileName = "./testdata/sample.json"

func Benchmark_EncodingJson_Read_Bytes_ForSingle_Value(b *testing.B) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var items []StructToRead
		json.Unmarshal(data, &items)
		item1000 := items[999]
		if item1000.Number != 1000 {
			b.Fatal("Не тот номер", item1000.Number)
		}
	}
}

func Benchmark_JsonIter_Read_Bytes_ForSingle_Value(b *testing.B) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		b.Fatal(err)
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var items []StructToRead
		json.Unmarshal(data, &items)
		item1000 := items[999]
		if item1000.Number != 1000 {
			b.Fatal("Не тот номер", item1000.Number)
		}
	}

}

func Benchmark_JsonParser_Read_Bytes_ForSingle_Value(b *testing.B) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		number, err := jsonparser.GetInt(data, "[999]", "number")
		if err != nil {
			b.Fatal("Что-то не так с парсером", err)
		}
		if number != 1000 {
			b.Fatal("Не тот номер", number)
		}
	}

}
