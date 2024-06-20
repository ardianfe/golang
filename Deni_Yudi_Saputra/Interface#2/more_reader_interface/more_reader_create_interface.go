package main

import (
	"fmt"
	"io/ioutil"
)

// Interface

type FileHandler interface {
	Create(data string) error
	Read() ([]byte, error)
}

// Deklarasi

type FileReaderWriter struct {
	filename string
}

// Method receiver

func (f *FileReaderWriter) Create(data string) error {
	return ioutil.WriteFile(f.filename, []byte(data), 0644) //buat file dengan namanya dari struct
}

func (f *FileReaderWriter) Read() ([]byte, error) {
	data, err := ioutil.ReadFile(f.filename) // cek file, returnan ReadFile byte

	if err != nil {
		return nil, err
	}
	return data, nil
}

// Fungsi perantara interface File Handler dan method nya, daripada kita nanti definisiin lagi

func readData(fh FileHandler) ([]byte, error) {
	return fh.Read()
}

func createData(fh FileHandler, data string) error {
	return fh.Create(data)
}

func main() {

	// ========== PROCESS CREATE FILE DAN ISI DATA DI FILE =============

	//  nama file nya
	nameFile := "contoh.txt"
	dataToWrite := "ANjayyy slebeww kiw kiw"

	// membuat instance atau mengisikan isi struct ke filehandler
	fileHandlerCreated := &FileReaderWriter{filename: nameFile}

	// coba create data ke file menggunakan fungsite  dari createData
	err := createData(fileHandlerCreated, dataToWrite)
	if err != nil {
		fmt.Println("Error create to file:", err)
		return
	}

	fmt.Println("Create file and create data successs")

	// =========== PROCESS READING DATA DARI FILE ============

	fileHandlerReader := &FileReaderWriter{filename: nameFile}

	data, err := readData(fileHandlerReader)

	if err != nil {
		fmt.Println("Error reading data", err)
		return
	}

	fmt.Printf("Baca data dari file %s yang berisikan %s", nameFile, string(data))
}
