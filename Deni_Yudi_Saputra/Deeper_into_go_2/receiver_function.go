package main

import "fmt"


// pertama kita deklarasikan struct 

type FullName struct{
	Depan string
	Belakang string
}

// buat function recivernya 

func (nl FullName) CetakNama(){
	fmt.Printf("Hi, nama lengkap saya %s %s \n", nl.Depan, nl.Belakang)
}

// fungsi buat merubah namanya menggunakan params 

func (nl *FullName) UbahNama(belakangBaru string){
	nl.Belakang = belakangBaru
}

func main()  {
	// buat var nilai struct 
	namaLengkap := FullName{"Deni", "Myke"}

	namaLengkap.CetakNama()

	namaLengkap.UbahNama("Slebew")

	namaLengkap.CetakNama()
}



// // pertama deklarasikan struct

// type NamaLengkap struct {
// 	namaDepan    string
// 	namaBelakang string
// }

// // buat fungsi reciver dan method, receiver tuh hampir sama kayak param, func trs ada param kemudian methodnya

// func (n NamaLengkap) CetakNama() {
// 	fmt.Printf("Halo Perkenalkan nama saya %s %s \n",n.namaDepan,n.namaBelakang)
// }


// //harus pakai pointer supaya data berubah, bukan merubah dari salinan 
// func (n *NamaLengkap) UbahNama(namaBelakangBaru string){
// 	n.namaBelakang = namaBelakangBaru
// }

// func main()  {
// 	// kita buat untuk mengisikan structnya 
// 	fullName := NamaLengkap{"Deni","Yudi"}

// 	// ambil dari methodnya yg ada recivernya
// 	fullName.CetakNama()

// 	fullName.UbahNama("Myke")

// 	fullName.CetakNama()
// }

// ====== COBA RECEIVER FUNCTION 2 ======
// type Saya struct {
// 	Name string
// 	Umur int
// }

// // fungsi receiver trs pada method
// func (p Saya) Ucapan() {
// 	fmt.Printf("Halo, nama saya adalah %s dan umur saya %d \n", p.Name, p.Umur)
// }

// func (p *Saya) Ulangtahun(){
// 	p.Umur++ //kalau gak pakai pointer hanya mengubah salinan dari p
// 	fmt.Println("tahun depan umur saya: ",p.Umur)//mencetak nilai baru dari salinan
// }

// func main() {
// 	// buat data baru dengan mengisi dari struct Saya
// 	person1 := Saya{"Deni", 70}

// 	fmt.Println(person1)

// 	person1.Ucapan()
// 	person1.Ulangtahun()// hasil 71
// 	fmt.Println(person1.Umur) // hasil 70 tetap, ketika dicetak tanpa pointer maka akan tetep tidak bertambah person1.Umur kalau gapakai pointer
// 	person1.Ucapan()
// }

// ====== COBA RECEIVER FUNCTION 1 ======

// type rak []string

// // panggil receiver sebelum method print(), jadi bisa mengambil value dari rak itu
// func (r rak) print() {

// 	fmt.Println(r)
// }

// // ======== SET RECEIVER ==========

// type siswa struct{
// 	name string
// 	age int
// }

// // pakai pointer pada *siswa supaya data bisa berubah karena kalau gak pakai pointer maka yang dirubah cuman fungsi ini gak merubah dari struct siswa yang atas
// func (s *siswa) setUmur(umur int){//(s siswa adalah receiver)
// 	s.age = umur
// }

// func (s siswa) setUmur2(umur int){//(s siswa adalah receiver)
// 	s.age = umur
// }

// func (s siswa) printUmur() {
// 	fmt.Println(s.age)
// }

// func main() {
// 	rak1 := rak{"buku", "komputer", "laptop"}
// 	rak1.print() // manggil maka rak dari fungsi receiver, manggil dari methodnya dia yaitu rak1

// 	// SET RECEIVER
// 	siswa2 := siswa{name:"peter",age: 5}
// 	siswa1 := siswa{name:"deni",age: 10}

// 	siswa1.setUmur(20)
// 	siswa2.setUmur2(203)

// 	siswa1.printUmur()
// 	siswa2.printUmur()

// }
