package main

import "fmt"

type Book struct {
	Judul   string
	Penulis string
	Jumlah  int
	Harga   float64
}

var inventaris []Book

//fungsi tambah buku
//fungsi tambah stock buku
//fungsi cetak inventaris buku
//fungsi cari buku
//fungsi hapus buku
//fungsi tambah atau update stok buku
//fungsi hitung nilai harga buku
//fungsi edit buku

func tambahBuku(judul, penulis string, jumlah int, harga float64) {
	bukubaru := Book{
		Judul:   judul,
		Penulis: penulis,
		Jumlah:  jumlah,
		Harga:   harga,
	}
	inventaris = append(inventaris, bukubaru)
}

func perbaruiJumlah(judul string, jumlah int) {
	for i := range inventaris {
		if inventaris[i].Judul == judul {
			inventaris[i].Jumlah += jumlah
			return
		}
		fmt.Printf("buku dengan judul: %s tidak ditemukan", judul)
	}
}

func cetakInventaris() {
	fmt.Printf("Daftar Buku yang tersedia di dalam inventaris\n")
	for _, buku := range inventaris {
		fmt.Printf("Judul Buku: %s, Penulis: %s, Jumlah: %d, Harga: %.2f\n",
			buku.Judul, buku.Penulis, buku.Jumlah, buku.Harga)
	}
}

func cariBuku(judul string) {
	for _, buku := range inventaris {
		if buku.Judul == judul {
			fmt.Printf("Buku ditemukan dengan data berikut:\n Judul: %s\n Penulis: %s\n Harga: %.2f\n Jumlah: %d\n",
				buku.Judul, buku.Penulis, buku.Harga, buku.Jumlah)
			return
		}
	}
	fmt.Printf("Buku dengan judul: %s tidak ditemukan\n", judul)
}

func hapusBuku(judul string) {
	for i := range inventaris {
		if inventaris[i].Judul == judul {
			inventaris = append(inventaris[:i], inventaris[i+1:]...)
			fmt.Printf("Buku: %s, berhasil dihapus\n", judul)
			return
		}
	}
	fmt.Printf("Buku: %s tidak ditemukan\n", judul)
}

func tambahAtauPerbaruiBuku(judul, penulis string, jumlah int, harga float64) {
	for i := range inventaris {
		if inventaris[i].Judul == judul {
			inventaris[i].Jumlah += jumlah
			fmt.Printf("buku sudah tersedia, maka persediaan ditambahkan menjadi total = %d buku\n", inventaris[i].Jumlah)
			return
		} else {
			bukubaru := Book{
				Judul:   judul,
				Penulis: penulis,
				Harga:   harga,
				Jumlah:  jumlah,
			}
			inventaris = append(inventaris, bukubaru)
			fmt.Printf("buku %s berhasil ditambahkan\n", judul)
			return
		}
	}
}

func hitungNilaiTotal() {
	var totalHarga float64 = 0.0
	totalBuku := 0
	for _, buku := range inventaris {
		totalHarga += (float64(buku.Jumlah) * buku.Harga)
		totalBuku += buku.Jumlah
	}
	fmt.Printf("Total Nilai: Rp%.2f\nTotal Buku: %d\n", totalHarga, totalBuku)
}

func editbuku(judul, penulis string, jumlah int, harga float64) {
	for i := range inventaris {
		if inventaris[i].Judul == judul {
			inventaris[i].Penulis = penulis
			inventaris[i].Jumlah = jumlah
			inventaris[i].Harga = harga
			fmt.Printf("Data buku %s berhasil di perbarui\n", judul)
			return
		}
	}
	fmt.Printf("buku dengan judul %s tidak ditemukan", judul)
}

func main() {
	tambahBuku("Bumi", "Tere Liye", 3, 15000)
	tambahBuku("Langit", "Abdul", 1, 17000)
	tambahBuku("Bintang", "Salman", 1, 20000)
	cetakInventaris()
	hitungNilaiTotal()

	// hapusBuku("Bintang")
	// cetakInventaris()
	// hitungNilaiTotal()

	// tambahAtauPerbaruiBuku("Gunung", "Tere Liye", 1, 17000)
	// cetakInventaris()
	// hitungNilaiTotal()

	// editbuku("Bumi", "issac", 8, 50000)
	// cetakInventaris()
	// hitungNilaiTotal()

	// perbaruiJumlah("Bumi", 2)
	// cetakInventaris()
	// hitungNilaiTotal()

	// cariBuku("Langit")
}
