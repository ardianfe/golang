package main

import "fmt"

type Person struct {
	Nama string
	Umur int
}

type Employee struct {
	ID     int
	Person *Person
}

func main() {
	p := Person{
		Nama: "Deni",
		Umur: 20,
	}

	e := Employee{
		ID:     12300,
		Person: &p, //salinan person, pointer = nanti akan merubah data asli ketika mendeklarasikan ulang, karena ketika tidak menggunakan pointer maka menyalin data tersebut dan nantinya tidak akan berubah
	}

	fmt.Printf("Employee ID : %d\n", e.ID)
	fmt.Printf("Employee Nama : %s\n", e.Person.Nama)
	fmt.Printf("Employee Umur : %d\n", e.Person.Umur)

	e.Person.Umur = 30
	fmt.Printf("Update Employee Umur  : %d\n", e.Person.Umur) //30
	fmt.Printf("Update Person Umur : %d\n", p.Umur)           //30 hasil pointer, ketika tidak pointer hasil akan tetap 20 karena itu merupakan salinan

	// jadi ketika membuat pointer perubahan terjadi ketika instance asli dari Person itu akan berubah.

	// Jadi ketika menggunakan pointer kita merubah langsung data asli dari (struct person) karena kita mengakses langsung bukan dari salinannya	 dan ketika tidak menggunakan pointer maka data asli tidak berubah karena ketika kita akan merubah di (Employee itu kan Person merupakan salinan dari struct person ).

	// Ketika Anda menggunakan pointer, Anda bekerja dengan alamat memori dari variabel asli. Ini berarti perubahan yang Anda buat melalui pointer akan langsung mempengaruhi variabel asli. Ini sangat efisien dalam hal penggunaan memori karena tidak membuat salinan dari data.

	// --	Dengan Pointer:
	// e.Person menyimpan alamat memori dari p. Jadi, setiap perubahan yang dilakukan melalui e.Person langsung mempengaruhi p.
	// Hasil menunjukkan bahwa perubahan umur melalui e.Person.Age juga terlihat di p.Age.
	// --	Tanpa Pointer:
	// e.Person adalah salinan dari p. Jadi, setiap perubahan yang dilakukan pada e.Person hanya mempengaruhi salinan tersebut dan tidak mempengaruhi p.
	// Hasil menunjukkan bahwa perubahan umur pada e.Person.Age tidak mempengaruhi p.Age.
}
