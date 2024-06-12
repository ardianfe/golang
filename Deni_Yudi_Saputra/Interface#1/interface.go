// 1. Purpose Interface
// Interface adalah kumpulan method method yang tidak memiliki isi dan dibungkus dalam nama tertentu dan juga bisa dipakai type data. digunakan untuk mendifinisikan method2 yang harus diimplementasikan oleh sebuah tipe

// 2. Problem without interface
// Tanpa interface, kita harus menulis kode yang sangat spesifik untuk setiap tipe, yang membuat kode menjadi kaku dan sulit untuk dikembangkan atau diubah. Misalnya, jika kita ingin membuat fungsi yang bisa bekerja dengan berbagai tipe, kita harus menulis banyak versi dari fungsi tersebut untuk setiap tipe.

// 3. Interface in practice
package main

import "fmt"

// Deklarasi
type Nilai struct {
	a, b int
}

type NamaLengkap struct {
	namaDepan    string
	namaBelakang string
}

// Method method
func (n Nilai) Perkalian() int {
	return n.a * n.b
}
func (n Nilai) Pertambahan() int {
	return n.a + n.b
}

// ==============
func (nl NamaLengkap) CetakNama() string {
	nama := nl.namaDepan + " " + nl.namaBelakang
	return nama
}

func (nl *NamaLengkap) UbahNama(namaBelakangBaru string) {

	nl.namaBelakang = namaBelakangBaru
}

// Interface nya
type Hitung interface {
	Perkalian() int
	Pertambahan() int
}

// ===========
type HasilNama interface {
	CetakNama() string
	UbahNama(string)
}

func main() {
	var hitung Hitung
	hitung = Nilai{2, 4}

	fmt.Println("Hasil perkalian adalah :", hitung.Perkalian())
	fmt.Println("Hasil perkalian adalah :", hitung.Pertambahan())

	fmt.Println("===========")

	var nama HasilNama

	nama = &NamaLengkap{"Deni", "Yudi"}

	fmt.Printf("Nama lengkap : %s\n", nama.CetakNama())

	nama.UbahNama("Cagur") //ubah namaE

	fmt.Printf("Nama lengkap setelah di ubah : %s\n", nama.CetakNama())

	fmt.Println("===========")
	moreInterface()

}

// 4. Roles interface
// -	interface mendifinisikan serangkaian metode
// - tipe yang mengimplementasikan semua metode dalam interface dianggap mengimplementasikan interface tersebut
// - implementasi metode harus sesuai dengan tanda tangan yang ada di interface

// 5. Extra interface nodes
//  - interface didalam interface

type NestedInterface interface {
	Hitung
	HasilNama
}

// 6. HTTP package
// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }

// 7. Reading docs interface

// 8. More interface syntax

// Definisikan sebuah interface dengan satu metode
type Sapaan interface {
	Sapa() string
}

// Definisikan sebuah struct yang akan mengimplementasikan interface
type Manusia struct {
	Nama string
}

// Implementasikan metode Sapa() dari interface Sapaan pada struct Manusia
func (m Manusia) Sapa() string {
	return "Halo, nama saya " + m.Nama
}

// Fungsi untuk menerima argumen yang bertipe Sapaan
func SapaSemua(s Sapaan) {
	fmt.Println(s.Sapa())
}

func moreInterface ()  {
	// Buat objek Manusia dengan nama Deni
    deni := Manusia{"Deni"}

    // Panggil fungsi SapaSemua dengan argumen objek Manusia
    SapaSemua(deni)

    // Deklarasi variabel bertipe interface Sapaan
    var sapaan Sapaan

    // Inisialisasi variabel sapaan dengan objek Manusia
    sapaan = Manusia{"Budi"}

    // Panggil metode Sapa() dari variabel sapaan
    fmt.Println(sapaan.Sapa())
}