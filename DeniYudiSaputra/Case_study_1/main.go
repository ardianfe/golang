package main

import "fmt"

type Buku struct {
	judul   string
	penulis string
	harga   float64
	jumlah  int
}

var inventaris []Buku

func tambahBuku(judul string, penulis string, harga float64, jumlah int) {
	buku := Buku{
		judul:   judul,
		penulis: penulis,
		harga:   harga,
		jumlah:  jumlah,
	}
	inventaris = append(inventaris, buku)
}

func perbaruiJumlah(judul string, jumlahbaru int) {
	// ngambil data dulu dari inventaris yang merupakan slice dari buku untuk mengecek judul para parameter apakah ada atau sama pada inventaris
	for i := range inventaris {
		if inventaris[i].judul == judul {
			inventaris[i].jumlah = jumlahbaru
		}
	}

	// fmt.Println("judul buku tidak ditemukan pada inventaris")
}

func hitungNilaiTotal() float64 {
	total := 0.0

	for _, i := range inventaris {
		total += i.harga * float64(i.jumlah)
	}

	return total
}

func main() {
	// menambah buku
	tambahBuku("Simpson", "Deni", 3000.0, 10)
	tambahBuku("Starlest", "Jono", 40000.0, 10)

	// memperbarui jumlah buku
	perbaruiJumlah("Simpson", 30)
	
	for _, buku := range inventaris {
		fmt.Printf("Judul: %s, Penulis: %s, Harga: %.2f, Jumlah: %d\n", buku.judul, buku.penulis, buku.harga, buku.jumlah)
	}
	// menghitung nilai total inventaris
	totalValue := hitungNilaiTotal()

	fmt.Printf("Nilai total inventaris: Rp.%.2f\n", totalValue)

	// menampilkan data buku dengan perulangan range

}
