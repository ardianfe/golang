package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func main() {
	file, err := os.Open("file/data.xlsx")

	if err != nil {
		fmt.Println("Error membuka file", err)
		return
	}

	// tutup file ketika sudah selesai
	defer file.Close()

	data, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println("Eror membaca file ", err)
	}

	// cetak isi file

	fmt.Println([]byte(data))

	//==== BUKA FILE NYA LANGSUNG =======

	// Path ke file yang akan dibuka
	filePath := "file/data.xlsx"

	// Membuka file menggunakan perintah 'start' di Windows
	cmd := exec.Command("cmd", "/c", "start", filePath)
	err = cmd.Start()
	if err != nil {
		fmt.Println("Error membuka file:", err)
		return
	}

	fmt.Println("File berhasil dibuka")

}
