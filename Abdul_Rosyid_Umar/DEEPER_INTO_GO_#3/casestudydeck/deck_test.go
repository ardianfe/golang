package main

import (
	"os"
	"testing"
)

// test fungsi buat deck
func TestBuatDeck(t *testing.T) {
	baru := BuatDeck()
	if len(baru) != 52 {
		t.Errorf("diharapkan 52, yang di dapat %v", len(baru))
	}
}

// test fungsi simpan dan baca file
func TestSimpanBacaFile(t *testing.T) {
	os.Remove("example.txt")
	baru := BuatDeck()
	baru.SimpanKeFile("example.txt")
	hasilbaca := bacaDariFile("example.txt")
	if len(hasilbaca) != 52 {
		t.Errorf("diharapkan 52, yang di dapat %v", len(hasilbaca))
	}
	os.Remove("example.txt")
}

// test fungsi acak
func TestAcak(t *testing.T) {
	dek := BuatDeck()
	dek.acak()
	if dek.keString() == BuatDeck().keString() {
		t.Errorf("Dek harusnya teracak, tapi tidak berubah")
	}
}
