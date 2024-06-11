package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Kartu struct {
	Jenis string
	Nilai string
}

type Dek []Kartu

func (d Dek) cetak() {
	for _, kartu := range d {
		fmt.Println(kartu)
	}
}

func (d Dek) keString() string {
	var str []string
	for _, kartu := range d {
		str = append(str, kartu.Nilai+" dari "+kartu.Jenis)
	}
	return strings.Join(str, ", ")
}

func buatDek() Dek {
	jenisKartu := []string{"Sekop", "Hati", "Wajik", "Keriting"}
	nilaiKartu := []string{"As", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}

	var dek Dek
	for _, jenis := range jenisKartu {
		for _, nilai := range nilaiKartu {
			dek = append(dek, Kartu{Jenis: jenis, Nilai: nilai})
		}
	}
	return dek
}

func (d Dek) bagiDek() (Dek, Dek) {
	return d[:len(d)/2], d[len(d)/2:]
}

func (d Dek) gabungkan() string {
	var kartuStrings []string
	for _, kartu := range d {
		kartuStrings = append(kartuStrings, kartu.Nilai+" dari "+kartu.Jenis)
	}
	return strings.Join(kartuStrings, ", ")
}

func (d Dek) simpanKeFile(namaFile string) error {
	return os.WriteFile(namaFile, []byte(d.keString()), 0666)
}

func bacaDariFile(namaFile string) (Dek, error) {
	bs, err := os.ReadFile(namaFile)
	if err != nil {
		return nil, err
	}

	s := strings.Split(string(bs), ", ")
	var dek Dek
	for _, kartuString := range s {
		kartuInfo := strings.Split(kartuString, " dari ")
		dek = append(dek, Kartu{Jenis: kartuInfo[1], Nilai: kartuInfo[0]})
	}
	return dek, nil
}

func (d Dek) acak() {
	sumber := rand.NewSource(time.Now().UnixNano())
	r := rand.New(sumber)

	for i := range d {
		posisiBaru := r.Intn(len(d))
		d[i], d[posisiBaru] = d[posisiBaru], d[i]
	}
}

func (d Dek) filterBerdasarkanJenis(jenis string) Dek {
	var filtered Dek
	for _, kartu := range d {
		if kartu.Jenis == jenis {
			filtered = append(filtered, kartu)
		}
	}
	return filtered
}

func (d Dek) filterBerdasarkanNilai(nilai string) Dek {
	var filtered Dek
	for _, kartu := range d {
		if kartu.Nilai == nilai {
			filtered = append(filtered, kartu)
		}
	}
	return filtered
}

func (d *Dek) tambahKartu(kartu Kartu) {
	*d = append(*d, kartu)
}

func (d *Dek) hapusKartu(kartu Kartu) {
	for i, k := range *d {
		if k == kartu {
			*d = append((*d)[:i], (*d)[i+1:]...)
			break
		}
	}
}

func (d *Dek) sisipkanKartu(kartu Kartu, posisi int) {
	if posisi < 0 || posisi > len(*d) {
					fmt.Println("Posisi tidak valid")
					return
	}
	*d = append((*d)[:posisi], append([]Kartu{kartu}, (*d)[posisi:]...)...)
}

func (d *Dek) updateKartu(posisi int, kartuBaru Kartu) {
	if posisi < 0 || posisi >= len(*d) {
					fmt.Println("Posisi tidak valid")
					return
	}
	(*d)[posisi] = kartuBaru
}


func main() {
	// Membuat dek
	dek := buatDek()
	fmt.Println("Dek baru dibuat:")
	dek.cetak()

	// Menambahkan kartu ke dek
	kartuBaru := Kartu{Jenis: "Sekop", Nilai: "As"}
	dek.tambahKartu(kartuBaru)
	fmt.Println("\nDek setelah menambahkan kartu baru:")
	dek.cetak()

	// Menghapus kartu dari dek
	dek.hapusKartu(kartuBaru)
	fmt.Println("\nDek setelah menghapus kartu:")
	dek.cetak()

	// Mengupdate kartu dari dek
	kartuUpdate := Kartu{Jenis: "Baru", Nilai: "xx"}
	dek.updateKartu(1, kartuUpdate)
	fmt.Println("\nDek setelah update kartu:")
	dek.cetak()

	// Menggunakan filter
	filteredByJenis := dek.filterBerdasarkanJenis("Sekop")
	fmt.Println("\nDek setelah difilter berdasarkan jenis 'Sekop':")
	filteredByJenis.cetak()

	filteredByNilai := dek.filterBerdasarkanNilai("As")
	fmt.Println("\nDek setelah difilter berdasarkan nilai 'As':")
	filteredByNilai.cetak()

	// Menyisipkan kartu di posisi tengah
	posisiTengah := len(dek) / 2
	kartuBaruLain := Kartu{Jenis: "Hati", Nilai: "Joker"}
	dek.sisipkanKartu(kartuBaruLain, posisiTengah)
	fmt.Println("\nDek setelah menyisipkan kartu di posisi tengah:")
	dek.cetak()

	// Mengacak dek
	dek.acak()
	fmt.Println("\nDek setelah diacak:")
	dek.cetak()

	// Menyimpan dek ke file
	namaFile := "_dekuji"
	err := dek.simpanKeFile(namaFile)
	if err != nil {
					fmt.Println("Error menyimpan dek ke file:", err)
					os.Exit(1)
	}
	fmt.Println("\nDek disimpan ke file:", namaFile)

	// Membaca dek dari file
	dekDariFile, err := bacaDariFile(namaFile)
	if err != nil {
					fmt.Println("Error membaca dek dari file:", err)
					os.Exit(1)
	}
	fmt.Println("\nDek dibaca dari file:")
	dekDariFile.cetak()

	// Menghapus file setelah pengujian
	err = os.Remove(namaFile)
	if err != nil {
					fmt.Println("Error menghapus file:", err)
	}
	fmt.Println("\nFile", namaFile, "dihapus")
}

