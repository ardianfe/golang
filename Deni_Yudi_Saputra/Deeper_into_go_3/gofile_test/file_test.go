package main

import (
	"os"
	"testing"
)

func TestWriteToFile(t *testing.T) {
	filename := "testfile.docx"
	data := "Hello, Golang!"

	err := WriteToFile(filename, data)
	if err != nil {
		t.Fatalf("Gagal menulis ke file: %v", err)
	}

	defer os.Remove(filename)

	readData, err := ReadFromFile(filename)
	if err != nil {
		t.Fatalf("Gagal membaca dari file: %v", err)
	}

	if readData != data {
		t.Errorf("Data tidak sesuai, dapat: %s, harap: %s", readData, data)
	}
}

func TestReadFromFileFileNotExist(t *testing.T) {
	filename := "nonexistentfile.txt"

	_, err := ReadFromFile(filename)
	if err == nil {
		t.Fatal("Seharusnya gagal karena file tidak ada")
	}

	expectedError := "file tidak ada"
	if err.Error() != expectedError {
		t.Errorf("Error tidak sesuai, dapat: %s, harap: %s", err.Error(), expectedError)
	}
}
