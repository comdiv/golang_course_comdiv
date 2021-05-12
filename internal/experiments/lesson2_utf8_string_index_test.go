package experiments

import (
	"reflect"
	"testing"
)

func TestStringsAreJustBytesByIndex(t *testing.T) {
	ruHello := "Привет"
	// получаем как мы думаем первый символ
	firstChar := ruHello[0]
	// смотрим тип и хотелось бы int32 для UTF-8
	charType := reflect.TypeOf(firstChar).Kind()
	// но нет, это действительно будет байт
	if charType != reflect.Uint8 {
		t.Errorf("Expected type is Uint8 but it was %s", charType.String())
	}
	// на всякий случай страхуемся - вдруг у нас однобайтовая кодировка затесалась
	if string(firstChar) != "Ð" {
		t.Errorf("Expected encoding UTF-8 and char `Ð` but was %s", string(firstChar))
	}
}

func TestHowGetUTF8Rune(t *testing.T) {
	ruHello := "Привет"
	asUtfChars := []rune(ruHello)
	first3rus := string(asUtfChars[0:3])
	if first3rus != "При" {
		t.Errorf("Expected `При` but was `%s`", first3rus)
	}
	firstChar := asUtfChars[0]
	charType := reflect.TypeOf(firstChar).Kind()
	// как ни странно ну руны не Uint32, Int32
	if charType != reflect.Int32 {
		t.Errorf("Expected type is Int32 but it was %s", charType.String())
	}
	if string(firstChar) != "П" {
		t.Errorf("Expected encoding UTF-8 and char `П` but was %s", string(firstChar))
	}
}
