package main

import "fmt"

type Buku struct {
	Judul, Penulis string
}

func main() {
	buku := Buku{
		Judul:   "Langit",
		Penulis: "Tere Liye",
	}
	fmt.Printf("judul buku sebelum diperbaru adalah : %s\n", buku.Judul)
	buku.Judul = "Bumi"
	fmt.Printf("judul buku sebelum diperbaru adalah : %s\n", buku.Judul)
}
