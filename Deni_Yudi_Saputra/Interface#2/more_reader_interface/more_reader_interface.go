package main

import (
	"fmt"
	"io/ioutil"
)

type FileReader struct {
	filename string
}

// method receiver
func (f *FileReader) Read() ([]byte, error) {
	data, err := ioutil.ReadFile(f.filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Function

// Fungsi perantara untuk menggunakan interface reader
func readData(r Reader) ([]byte, error) {
	return r.Read()
}

type Reader interface {
	Read() ([]byte, error)
}

func main() {
	reader := &FileReader{filename: "contoh.txt"}

	data, err := readData(reader) // buat perantara dengan membuat parameter yang isinya nyanti dari FileReader nya misal gak pakai dari readData gitu kita bisa menggunakan langsung sperti dibawah ini

	// var r Reader = reader
	// data,err := r.Read()

	if err != nil {
		fmt.Println("Error reading", err)
		return
	}

	fmt.Println("Data read from file : ")
	fmt.Println(string(data))

}
