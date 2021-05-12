package experiments

import "testing"

func TestThatAFloatMultBIntIsNotBIntMult(t *testing.T) {
	// умножэение не зависит от порядка операндов
	var floatMult float32 = 5.1 * 2
	var intByDocMult float32 = 2 * 5.1
	if floatMult != intByDocMult && floatMult != 10.2 {
		t.Errorf("a * b != b * a, 5.1 * 2 != 10.2, but `%f` and `%f`", floatMult, intByDocMult)
	}

	var floatDiv float32 = 10.2 / 2
	var intByDocDiv float32 = 10 / 2.5

	if floatDiv != 5.1 || intByDocDiv != 4.0 {
		t.Errorf("`%f`!=5.1 and `%f`!=4.0", floatDiv, intByDocDiv)
	}
}
