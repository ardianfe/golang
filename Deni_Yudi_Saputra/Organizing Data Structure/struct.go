package main

import "fmt"

//DEFINING STRUCT nama person
type Person struct {
	// DEKLARASI STRUCT
	Nama string
	Umur int
}

func main() {

	var person1 Person

	//DEKLARASI VALUE STRUCT
	person1.Nama = "Deni"
	person1.Umur = 32

	fmt.Printf("Nama : %s dan Umur : %d\n", person1.Nama, person1.Umur)

	//UPDATE STRUCT VALUE
	person1.Nama = "Alex"
	person1.Umur = 23

	fmt.Printf("Nama : %s dan Umur : %d\n", person1.Nama, person1.Umur)

	// ===============

	person2 := Person{
		Nama: "Yudi",
		Umur: 1,
	}

	fmt.Printf("Nama : %s dan Umur : %d", person2.Nama, person2.Umur)
}
