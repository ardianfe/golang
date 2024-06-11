package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jung-kurt/gofpdf"
	"github.com/tealeg/xlsx"
	"github.com/unidoc/unioffice/document"
)

func TxtOS() {
	filename := "file/data.txt"
	// Data yang akan disimpan

	data := "Coba data yang dicetak ya guys!"

	// Membuka atau membuat file
	file, err := os.Create(filename)

	if err != nil {
		fmt.Println("Error creating file : ", err)
		return
	}
	defer file.Close()

	// Menulis data ke file
	_, err = file.WriteString(data)

	if err != nil {
		fmt.Println("Error writing to file : ", err)
		return
	}

	fmt.Println("Data berhasil disimpan")
}

func TxtIOUtils() {
	dataio := "ini adalah contoh data yang disimpan ke file menggunkan ioutils"

	err := ioutil.WriteFile("data_io.txt", []byte(dataio), 0644) //0644 adalah permission pada ioUtil.writeFile dengan notasi oktal
	if err != nil {
		fmt.Println("Gagal menyimpan file :", err)
		return
	}

	fmt.Println("Data telah disimpan ke file txt")
}

func Word() {
	doc := document.New()

	// buat paragraf baru dengan teks yang diberikan

	paragraf := doc.AddParagraph()
	run := paragraf.AddRun()
	run.AddText("Halo, ini adalah dokumen word ")

	// menyimpan file word
	err := doc.SaveToFile("data.docx")
	if err != nil {
		fmt.Printf("Gagal menyimpan file : %s\n", err)
		return
	}

	fmt.Println("File word telah berhasil dibuat")
}

func PDF() {
	// buat objek pdf baru dengan ukuran A4
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Menambahkan halaman baru
	pdf.AddPage()

	// tambahkan teks ke pdf
	pdf.SetFont("Arial", "B", 16)

	pdf.Cell(100, 200, "Haloo semuanyaaa ini saya yaa") // nampilin dengan posisi lebar tinggi dan isinya

	pdf.AddPage()

	pdf.Cell(40, 10, "ini adalah halaman ke dua")

	err := pdf.OutputFileAndClose("file/Hello.pdf")
	if err != nil {
		panic(err)
	}
}

func Excel() {
	// buat excel baru mewaiki file excel yang kosong

	file := xlsx.NewFile()

	// membuat lembar kerja baru atau sheetnya
	sheet, err := file.AddSheet("Pertama")

	if err != nil {
		fmt.Println("Gagal membuat lembar kerja baru", err)
		return
	}

	// buat data menggunakan slice 2
	// [][] adalah kumpulan slice ,ketika menggunakan []string maka akan []string{"Kiw","Kew","Kuw"}
	data := [][]string{
		{"Nama", "Usia", "Kota"},
		{"Doe", "20", "Blitar"},
		{"Malik", "20", "Malang"},
		{"Lehman", "30", "Jombang"},
		{"Leh", "30", "Jombang"},
	}

	fmt.Println(data)

	// buat perulangan untuk menampilkan slice nya ke dalam excel

	for _, row := range data {
		newRow := sheet.AddRow()

		for _, cell := range row {
			newCell := newRow.AddCell()
			newCell.SetString(cell)
		}
	}

	// buat folder untuk menyimpan file nya
	if _, err := os.Stat("file"); os.IsNotExist(err) {
		err = os.Mkdir("file", os.ModePerm)
		if err != nil {
			fmt.Println("Gagal membuat folder ", err)
			return
		}
	}

	// menyimpan ke file excel
	err = file.Save("file/data.xlsx")
	if err != nil {
		fmt.Println("Gagal menyimpan data file excel :", err)
		return
	}

	fmt.Println("Data telah disimpan ke file excel")
}

func Html() {
	htmlContent := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Contoh HTML</title>
		<link rel="stylesheet" type="text/css" href="style.css">
	</head>
	<body>
		<h1>Halo, ini adalah contoh HTML</h1>
		<p>Ini adalah paragraf dalam file HTML.</p>
	</body>
	</html>
	`

	// Menulis konten HTML ke file
	err := os.WriteFile("file/coba.html", []byte(htmlContent), 0644)
	if err != nil {
		fmt.Printf("Gagal menulis file HTML: %s\n", err)
		return
	}

	fmt.Println("File HTML telah berhasil dibuat.")
}

func CSS() {
	cssContent := `
	body {
		font-family: Arial, sans-serif;
		background-color: #336273;
		margin: 0;
		padding: 0;
	}
	h1 {
		color: #fff;
	}

	.text-1{
		color : #fff
	}
	.text-2{
		color : #dddd23
	}
	.text-3{
		color : #qqqq23
	}
	`

	// Menulis konten CSS ke file
	err := os.WriteFile("file/style.css", []byte(cssContent), 0644)
	if err != nil {
		fmt.Printf("Gagal menulis file CSS: %s\n", err)
		return
	}

	fmt.Println("File CSS telah berhasil dibuat.")
}

func main() {
	PDF()
}
