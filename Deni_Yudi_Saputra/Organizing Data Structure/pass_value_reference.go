package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func modifyPerson(p Person) {
	p.Name = "Bob"
}

func modifyPersonRef(p *Person) {
	p.Name = "Bob"
}

func main() {
	// Value type
	person := Person{Name: "Alice", Age: 30}
	modifyPerson(person) // Pass by value
	fmt.Println(person)  // Output: {Alice 30}, tidak berubah karena salinan independen dibuat

	// Reference type
	personRef := Person{Name: "Alice", Age: 30}
	modifyPersonRef(&personRef) // Pass by reference menggunakan pointer
	fmt.Println(personRef)      // Output: {Bob 30}, berubah karena nilai asli dimodifikasi melalui referensi
}

// Value Type:

// Menyimpan nilai sebenarnya langsung di dalam variabel.
// Ketika variabel disalin atau diteruskan ke fungsi lain, salinan independen dari nilai tersebut dibuat.
// Contoh: int, float, bool, struct.
// Modifikasi terhadap salinan tidak mempengaruhi nilai asli.

// Reference Type:

// Menyimpan alamat memori (referensi) di mana nilai sebenarnya disimpan.
// Ketika variabel disalin atau diteruskan ke fungsi lain, hanya salinan referensi yang dibuat, bukan salinan nilai.
// Contoh: pointer, slice, map, channel.
// Modifikasi terhadap nilai melalui referensi mempengaruhi nilai asli.
