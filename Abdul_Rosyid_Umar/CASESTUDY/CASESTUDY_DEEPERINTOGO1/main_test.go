package main

import "testing"

func TestTambahBuku(t *testing.T) {
	inventaris = []Book{}

	tambahBuku("Bumi", "Tere Liye", 1, 12000)

	if len(inventaris) != 1 {
		t.Errorf("expected 1 book, got %d", len(inventaris))
	}
	if inventaris[0].Judul != "Bumi" {
		t.Errorf("expected Bumi, go %s", inventaris[0].Judul)
	}
}
