package go_test_test

// penamaan testing supaya efisien tapi untuk misal fungsi dari beda package

import (
	"testing"
	//cukup ambil dari gomod init misal masuk subfolder tinggal tambahin root nya
	// jadi untuk misal kita ambil package lain menggunakan . seperti dibawah ini misal masuk subfolder maka tinggal tambahin rootnya seperti file nya
	. "github.com/deniyudi/deeper_into_go_3/go_test"
)

// ===== TESTING BIASA ======
func TestAdd(t *testing.T) {
	result := Add(1, 4)

	if result != 5 {
		t.Logf("Hasilnya %d yang seharusnya %d", result, 5)
		t.Fail()
	}
}

// ===== TEST TABLE =====

// Bikin test table misal ada space input beberapa atau bisa apa aja misal ada (-) sampai (+)  dan misal mau bikin beberapa input space ga mungkin kita buat berulang ulang jadi kita menggunakan test table biar rapi dan efsien, testTable adalah setOfInputSpace dan expectedOutcome.
// setInputSpcae adalah Add(1,4),Add(-1,3), expectedOutcome adalah hasilnya, biasanya menggunakan slice struct =  []struct, kemudian bikin iterasinya

func TestTableAdd(t *testing.T) {
	testTable := []struct {
		a, b            int
		expectedOutcome int
	}{
		{a: 1, b: 4, expectedOutcome: 5},
		{a: -1, b: 4, expectedOutcome: 3},
		{a: 1, b: -4, expectedOutcome: -3},
		{a: -1, b: -4, expectedOutcome: -5},
		{a: 0, b: 0, expectedOutcome: 0},
	}

	// bikin perulangan

	for _, test := range testTable {
		result := Add(test.a, test.b)
		if result != test.expectedOutcome {
			t.Logf("Hasilnya %d yang seharusnya adalah %d", result, test.expectedOutcome)
			t.Fail()
		}

	}
}

// ====== TEST HIERARCHICAL =======

//subtest, go seperti nested test didalam test ada test, dimana ada parent tes nya nanji dijalankan paralel

func TestHierarchicalAdd(t *testing.T) {

	// menggunakan t.Run()
	t.Run("a=positif", func(t *testing.T) {
		a := 10
		t.Run("b=positif", func(t *testing.T) {
			result := Add(a, 5)
			if result != 15 {
				t.Logf("Hasilnya %d yang seharusnya adalah %d", result, 15)
				t.Fail()
			}
		})
		t.Run("b=negatif", func(t *testing.T) {
			result := Add(a, -5)
			if result != 5 {
				t.Logf("Hasilnya %d yang seharusnya adalah %d", result, 5)
				t.Fail()
			}
		})
	})
}
