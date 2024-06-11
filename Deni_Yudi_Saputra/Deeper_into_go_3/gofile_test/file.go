package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// WriteToFile menulis data ke file, membuat file jika belum ada
func WriteToFile(filename, data string) error {
	return ioutil.WriteFile(filename, []byte(data), 0644)
}

// ReadFromFile membaca data dari file, mengembalikan error jika file tidak ada
func ReadFromFile(filename string) (string, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return "", errors.New("file tidak ada")
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func main() {
	filename := "example.txt"
	data := "Hello, Golang!"

	// Menulis data ke file
	err := WriteToFile(filename, data)
	if err != nil {
		fmt.Println("Gagal menulis ke file:", err)
		return
	}
	fmt.Println("Berhasil menulis ke file.")

	// Membaca data dari file
	readData, err := ReadFromFile(filename)
	if err != nil {
		fmt.Println("Gagal membaca dari file:", err)
		return
	}
	fmt.Printf("Data dari file: %s\n", readData)

	// Menghapus file setelah selesai
	// err = os.Remove(filename)
	// if err != nil {
	// 	fmt.Println("Gagal menghapus file:", err)
	// 	return
	// }
	// fmt.Println("File berhasil dihapus.")
}
