package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/jung-kurt/gofpdf/v2"
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

// Fungsi untuk menyimpan file dengan data tertentu
func saveFile(filePath string, data []byte) error {
	err := ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Fungsi untuk membuat file PDF
func createPDF(folderPath, fileName string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, World!")
	filePath := filepath.Join(folderPath, fileName)
	err := pdf.OutputFileAndClose(filePath)
	if err != nil {
		return err
	}
	return nil
}

func createHtml(folderpath, filename string) error {
	htmlData := []byte("!html")
	filePath := filepath.Join(folderpath, filename)
	return saveFile(filePath, htmlData)
}

// Fungsi untuk membuat file Word
func createWord(folderPath, fileName string) error {
	wordData := []byte("This is a Word file")
	filePath := filepath.Join(folderPath, fileName)
	return saveFile(filePath, wordData)
}

// Fungsi untuk membuat file Excel
func createExcel(folderPath, fileName string) error {
	excelData := []byte("This is an Excel file")
	filePath := filepath.Join(folderPath, fileName)
	return saveFile(filePath, excelData)
}

// Fungsi untuk membaca file
func readFile(filePath string) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	fmt.Println("File content:")
	fmt.Println(string(data))
	return nil
}

func main() {
	// Path folder tujuan
	folderPath := "documents"

	// Membuat folder jika belum ada
	err := createFolder(folderPath)
	if err != nil {
		log.Fatal("Error creating folder:", err)
	}

	// Membuat file PDF
	pdfFileName := "example.pdf"
	err = createPDF(folderPath, pdfFileName)
	if err != nil {
		log.Fatal("Error creating PDF file:", err)
	}

	// Membuat file Word
	wordFileName := "example.docx"
	err = createWord(folderPath, wordFileName)
	if err != nil {
		log.Fatal("Error creating Word file:", err)
	}

	// Membuat file Excel
	excelFileName := "example.xlsx"
	err = createExcel(folderPath, excelFileName)
	if err != nil {
		log.Fatal("Error creating Excel file:", err)
	}

	htmlFileName := "index.html"
	err = createHtml(folderPath, htmlFileName)
	if err != nil {
		log.Fatal("Error creating HTML file:", err)
	}

	cssFileName := "style.css"
	cssdata := []byte("/*/ini css")
	pathfile := filepath.Join(folderPath, cssFileName)
	err = saveFile(pathfile, cssdata)
	if err != nil {
		log.Fatal("Error creating CSS file:", err)
	}

	log.Println("Files have been written to", folderPath)

	// Membaca file PDF
	err = readFile(filepath.Join(folderPath, pdfFileName))
	if err != nil {
		log.Fatal("Error reading PDF file:", err)
	}

	// Membaca file Word
	err = readFile(filepath.Join(folderPath, wordFileName))
	if err != nil {
		log.Fatal("Error reading Word file:", err)
	}

	// Membaca file Excel
	err = readFile(filepath.Join(folderPath, excelFileName))
	if err != nil {
		log.Fatal("Error reading Excel file:", err)
	}

}
