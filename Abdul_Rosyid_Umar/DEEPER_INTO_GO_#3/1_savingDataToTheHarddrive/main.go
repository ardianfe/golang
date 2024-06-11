package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Fungsi untuk membuat folder jika belum ada
func createFolder(folderPath string) error {
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		err = os.MkdirAll(folderPath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func saveFile(filePath string, data []byte) error {
	err := ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// Membuat slice byte dari string "Hello, World!"
	data := []byte("Hello, World!")

	// Menulis slice byte ke file "example.txt" dengan izin 0644
	err := ioutil.WriteFile("example.txt", data, 0644)

	// Memeriksa jika ada error saat menulis file
	// ini disebut error handling
	if err != nil {
		log.Fatal(err)
	}

	// Mencetak pesan log bahwa data telah ditulis ke "example.txt"
	log.Println("Data has been written to example.txt")

	//membaca file
	dataa, err := ioutil.ReadFile("example.pdf")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(dataa))
	pdfdll()
}

func tambahan() {
	// Membuka file example.txt
	file, err := os.Open("example.pdf")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close() // Memastikan file ditutup setelah selesai

	// Membaca seluruh konten file
	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Mencetak konten file ke konsol
	fmt.Println("File content:")
	fmt.Println(string(content))

}

func pdfdll() {
	// Data untuk disimpan
	pdfData := []byte("%PDF-1.4 ...")            // Contoh konten PDF
	wordData := []byte("This is a Word file")    // Contoh konten Word
	excelData := []byte("This is an Excel file") // Contoh konten Excel

	// Path folder tujuan
	folderPath := "documents"
	err := createFolder(folderPath)
	if err != nil {
		log.Fatal("Error creating folder:", err)
	}

	// Menyimpan file PDF
	pdfFilePath := filepath.Join(folderPath, "example.pdf")
	err = saveFile(pdfFilePath, pdfData)
	if err != nil {
		log.Fatal("Error saving PDF file:", err)
	}

	// Menyimpan file Word
	wordFilePath := filepath.Join(folderPath, "example.docx")
	err = saveFile(wordFilePath, wordData)
	if err != nil {
		log.Fatal("Error saving Word file:", err)
	}

	// Menyimpan file Excel
	excelFilePath := filepath.Join(folderPath, "example.xlsx")
	err = saveFile(excelFilePath, excelData)
	if err != nil {
		log.Fatal("Error saving Excel file:", err)
	}

	log.Println("Files have been written to", folderPath)
}
