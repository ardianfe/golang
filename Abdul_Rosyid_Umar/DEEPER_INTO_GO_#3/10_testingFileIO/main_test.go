package main

import (
	"os"
	"testing"
)

func TestFileIO(t *testing.T) {
	fileName := "testfile.txt"
	data := []byte("Hello, Test!")

	// Menulis ke file
	err := os.WriteFile(fileName, data, 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Membaca dari file
	readData, err := os.ReadFile(fileName)
	if err != nil {
		t.Fatal(err)
	}

	if string(readData) != string(data) {
		t.Errorf("Expected %s, got %s", string(data), string(readData))
	}

	// Menghapus file setelah tes
	err = os.Remove(fileName)
	if err != nil {
		t.Fatal(err)
	}
}
