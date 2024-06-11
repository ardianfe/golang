package main

import (
	"os"
	"testing"
)

func TestBuatDek(t *testing.T) {
	dek := buatDek()
	if len(dek) != 52 {
		t.Errorf("Jumlah kartu dalam dek harusnya 52, tapi dapat %v", len(dek))
	}
}

func TestBagiDek(t *testing.T) {
	dek := buatDek()
	tangan1, tangan2 := dek.bagiDek()
	if len(tangan1) != 26 {
		t.Errorf("Jumlah kartu dalam tangan1 harusnya 26, tapi dapat %v", len(tangan1))
	}
	if len(tangan2) != 26 {
		t.Errorf("Jumlah kartu dalam tangan2 harusnya 26, tapi dapat %v", len(tangan2))
	}
}

func TestSimpanKeFileDanBacaDariFile(t *testing.T) {
	os.Remove("_dekuji")

	dek := buatDek()
	err := dek.simpanKeFile("_dekuji")
	if err != nil {
		t.Errorf("Gagal menyimpan dek ke file: %v", err)
	}

	dekDariFile, err := bacaDariFile("_dekuji")
	if err != nil {
		t.Errorf("Gagal membaca dek dari file: %v", err)
	}

	if len(dekDariFile) != 52 {
		t.Errorf("Jumlah kartu dalam dek yang dibaca dari file harusnya 52, tapi dapat %v", len(dekDariFile))
	}

	os.Remove("_dekuji")
}

func TestAcak(t *testing.T) {
	dek := buatDek()
	dek.acak()

	// Tidak ada cara pasti untuk memeriksa apakah dek benar-benar teracak
	// Tapi kita bisa periksa bahwa elemen dalam dek acak tidak sama dengan dek yang diurutkan
	acakSama := true
	for i, kartu := range buatDek() {
		if dek[i] != kartu {
			acakSama = false
			break
		}
	}
	if acakSama {
		t.Errorf("Dek tidak diacak dengan benar")
	}
}

func TestFilterBerdasarkanJenis(t *testing.T) {
	dek := buatDek()
	filtered := dek.filterBerdasarkanJenis("Sekop")
	for _, kartu := range filtered {
		if kartu.Jenis != "Sekop" {
			t.Errorf("Kartu dengan jenis berbeda ditemukan: %v", kartu)
		}
	}
}

func TestFilterBerdasarkanNilai(t *testing.T) {
	dek := buatDek()
	filtered := dek.filterBerdasarkanNilai("As")
	for _, kartu := range filtered {
		if kartu.Nilai != "As" {
			t.Errorf("Kartu dengan nilai berbeda ditemukan: %v", kartu)
		}
	}
}

func TestTambahKartu(t *testing.T) {
	dek := buatDek()
	kartuBaru := Kartu{Jenis: "Sekop", Nilai: "Joker"}
	dek.tambahKartu(kartuBaru)

	if len(dek) != 53 {
		t.Errorf("Jumlah kartu dalam dek setelah menambahkan harusnya 53, tapi dapat %v", len(dek))
	}

	if dek[len(dek)-1] != kartuBaru {
		t.Errorf("Kartu terakhir harusnya %v, tapi dapat %v", kartuBaru, dek[len(dek)-1])
	}
}

func TestHapusKartu(t *testing.T) {
	dek := buatDek()
	kartuUnik := Kartu{Jenis: "Sekop", Nilai: "Joker"}
	dek.tambahKartu(kartuUnik)
	dek.hapusKartu(kartuUnik)

	if len(dek) != 52 {
		t.Errorf("Jumlah kartu dalam dek setelah menghapus harusnya 52, tapi dapat %v", len(dek))
	}

	for _, kartu := range dek {
		if kartu == kartuUnik {
			t.Errorf("Kartu %v harusnya sudah dihapus dari dek", kartuUnik)
		}
	}
}

func TestSisipkanKartu(t *testing.T) {
	dek := buatDek()
	kartuBaru := Kartu{Jenis: "Hati", Nilai: "Joker"}
	posisiTengah := len(dek) / 2
	dek.sisipkanKartu(kartuBaru, posisiTengah)

	if len(dek) != 53 {
		t.Errorf("Jumlah kartu dalam dek setelah menyisipkan harusnya 53, tapi dapat %v", len(dek))
	}

	if dek[posisiTengah] != kartuBaru {
		t.Errorf("Kartu di posisi %v harusnya %v, tapi dapat %v", posisiTengah, kartuBaru, dek[posisiTengah])
	}
}

func TestUpdateKartu(t *testing.T) {
	dek := buatDek()
	posisiUpdate := 10
	kartuUpdate := Kartu{Jenis: "Wajik", Nilai: "10"}
	dek.updateKartu(posisiUpdate, kartuUpdate)

	if dek[posisiUpdate] != kartuUpdate {
		t.Errorf("Kartu di posisi %v harusnya %v, tapi dapat %v", posisiUpdate, kartuUpdate, dek[posisiUpdate])
	}
}
