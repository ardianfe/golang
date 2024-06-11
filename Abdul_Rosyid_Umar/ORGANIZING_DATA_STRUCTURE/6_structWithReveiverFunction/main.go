package main

import "fmt"

type Orang struct {
	Nama, Umur string
}

func (b Orang) Greet() {
	fmt.Printf("Halo, %s", b.Nama)
}

func main() {
	ahmad := Orang{
		Nama: "Ahmad",
		Umur: "20",
	}
	ahmad.Greet()
}
