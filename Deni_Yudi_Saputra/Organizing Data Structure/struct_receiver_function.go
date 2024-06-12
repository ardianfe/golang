package main

import "fmt"

type FullName struct {
	FirstName string
	LastName  string
}

func (f FullName) PrintFullName() {
	fmt.Printf("Hi, Nama lengkap saya %s %s \n", f.FirstName, f.LastName)
}

func (fb *FullName) UpdateName(belakangBaru string) {
	fb.LastName = belakangBaru
}

func main() {
	namaLengkap := FullName{"Deni", "Yudi"}

	namaLengkap.PrintFullName()

	namaLengkap.UpdateName("Slebew")
	namaLengkap.PrintFullName()

}
