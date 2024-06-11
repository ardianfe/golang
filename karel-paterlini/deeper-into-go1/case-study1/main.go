package main

import "fmt"

// Slice paralel untuk menyimpan detail buku
var judulBuku []string
var penulisBuku []string
var hargaBuku []float64
var jumlahBuku []int

// tambahBuku menambahkan buku baru ke dalam inventaris.
// Judul, Penulis, Harga, dan Jumlah adalah parameter yang diperlukan untuk menyimpan detail buku baru.
func tambahBuku(judul, penulis string, harga float64, jumlah int) {
	judulBuku = append(judulBuku, judul)
	penulisBuku = append(penulisBuku, penulis)
	hargaBuku = append(hargaBuku, harga)
	jumlahBuku = append(jumlahBuku, jumlah)
}

// perbaruiJumlah memperbarui jumlah buku yang ada dalam inventaris.
// Judul adalah judul buku yang akan diperbarui jumlahnya, sementara JumlahBaru adalah jumlah baru buku tersebut.
func perbaruiJumlah(judul string, jumlahBaru int) {
	for i := range judulBuku {
		if judulBuku[i] == judul {
			jumlahBuku[i] = jumlahBaru
			return
		}
	}
	fmt.Printf("Buku dengan judul '%s' tidak ditemukan.\n", judul)
}

// hitungNilaiTotal menghitung nilai total semua buku dalam inventaris.
// Ini mengalikan harga buku dengan jumlahnya dan menambahkan hasilnya untuk semua buku.
func hitungNilaiTotal() float64 {
	var total float64
	for i := range judulBuku {
		total += hargaBuku[i] * float64(jumlahBuku[i])
	}
	return total
}

// cetakInventaris mencetak semua buku yang ada dalam inventaris beserta detailnya.
func cetakInventaris() {
	fmt.Println("Daftar buku dalam inventaris:")
	for i := range judulBuku {
		fmt.Printf("Judul: %s, Penulis: %s, Harga: Rp%.2f, Jumlah: %d\n",
			judulBuku[i], penulisBuku[i], hargaBuku[i], jumlahBuku[i])
	}
}

func main() {
	// Menambahkan beberapa buku ke dalam inventaris
	tambahBuku("dasar dasar belajar", "rel", 25000.00, 100)
	tambahBuku("dongeng", "peter", 55000.00, 50)

	// Memperbarui jumlah buku "dasar dasar belajar" menjadi 20
	perbaruiJumlah("dasar dasar belajar", 20)

	// Menghitung dan mencetak nilai total inventaris
	total := hitungNilaiTotal()
	fmt.Printf("Nilai total inventaris: Rp%.2f\n", total)

	// Mencetak daftar semua buku dalam inventaris
	cetakInventaris()
}
