// function
// sebelumnya kita sudah mengenal sebuah function yang wajib dibuat agar program kita bisa berjalan, yaitu function main
// function adalah sebuah blok kode yang sengaja dibuat dalam program agar bisa digunakan berulang-ulang
// cara membuat function sangat sederhana, hanya menggunakan kata kunci func lalu diikuti dengan nama functionnya dan blok kode isi functionnya
// setelah membuat function, kita bisa mengeksekusi function tersebut dengan memanggilnya menggunakan kata kunci nama functionnya diikuti tanda kurung buka, kurung tutup

package main

import (
	"fmt"
)

func main() {
	hello()
	helloto("bara", "eko", 18)
	helloto("panglima", "perang", 20)

	// return value
	result := gethello("bara")
	fmt.Println(result)
	fmt.Println(gethello("anwar"))
	fmt.Println(gethello("beni"))

	// returning multiple value
	firstname, lastname, age := getfullname()
	fmt.Println(firstname, lastname, age)

	// menghiraukan return value
	// multiple return value wajib ditangkap semua value nya
	// jika kita ingin menghiraukan return value tersebut, kita bisa menggunakan tanda "_ (garis bawah)"
	//alert! jika ingin mengeksekusi kode di bawah harus komen terlebih dahulu pada "returning multiple value!!!"
	// firstname, _, _ := getfullname()
	// fmt.Println(firstname)

	// named return values
	firstname, middlename, laslastname, age := getcomplatename()
	fmt.Println(firstname, middlename, laslastname, age)

	// variadic function
	// total := sumALL(10, 5, 10, 5, 5)
	// fmt.Println(total),

	// kita juga bisa memanggil langsung tanpa membuat variable atau slice
	fmt.Println(sumALL(10, 10, 10, 10))
	fmt.Println(sumALL(10, 10, 10, 10, 10))
	fmt.Println(sumALL(10, 10, 10, 10, 10, 10))

	// slice parameter
	// kadang ada kasus dimana kita menggunakan variadic function, namun memiliki variable berupa slice.
	// kita bisa menjadikan slice sebagai vararg parameter.
	numbers := []int{10, 10, 10}
	fmt.Println(sumALL(numbers...))

	// function value
	// kita bisa merubah function menjadi sebuah variable atau string
	// function juga bisa digunakan sebagai value
	goodbye := getGoodbye
	fmt.Println(goodbye("dim"))
	contoh := getGoodbye
	fmt.Println(contoh("bro"))

	// function as parameter
	hellofilter("broo", spamfilter)

	filter := spamfilter
	hellofilter("anjing", filter)

	// anonymous func
	blacklist := func(name string) bool {
		return name == "anjing"

	}
	registeruser("anjay", blacklist)
	// atau juga bisa menggunakan ini
	registeruser("anjin", func(name string) bool {
		return name == "anjing"
	})

}

func hello() {
	fmt.Println("halo function")
}

// function parameter
// saat membuat function, kadang kadang kita membutuhkan data dari luar, atau kita sebut parameter.
// kita bisa menambahkan parameter di function, bisa lebih dari satu.
// parameter tidaklah wajib, jadi kita bisa membuat function tanpa parameter seperti sebelumnya yang sudah kita buat
// namun jika kita menambahkan parameter di function, maka ketika memanggil function tersebut, kita wajib memasukan data ke parameternya.

// alert aja rel, firstname, lastname, dan age itu termasuk kedalam parameter jangan sampe lupa!!
func helloto(firstname string, lastname string, age int) {
	fmt.Println("hallo", firstname, lastname, age)
}

// function return value
// function bisa mengembalikan data
// untuk memberitahu bahawa function bisa mengembalikan data, kita harus menuliskan tipe data kembalian dari function tersebut
// jika function tersebut di deklarasikan dengan tipe data pengembalian, maka wajib di dalam function nya kita harus mengembalikan data
// untuk mengembalikan data dari function, kita bisa menggunakan kata kunci return, diikuti dengan datanya

func gethello(name string) string {
	hello := "hello " + name
	return hello
}

// returning multiple values
// function tidak hanya dapat mengembalikan satu value, tapi juga bisa multiple value.
// untuk memberitahu jika function mengembalikan multiple value, kita harus menulis semua tipe data return nya di function.

func getfullname() (string, string, int) {
	return "bara", "mukbang", 22
}

// named return values
// biasanya saat kita memberi tahu bahwa sebuah function mengembalikan value, maka kita hanya mendeklarasikan tipe data return di function
// namun kita juga bisa membuat variable secara langsung di tipe data return function nya

func getcomplatename() (firstname, middlename, lastname string, age int) {
	firstname = "joko"
	middlename = "anwar"
	lastname = "pro"
	age = 29

	return firstname, middlename, lastname, age
}

// variadic function
// parameter yang berada di posisi terakhir, memiliki kemampuan dijadikan sebuah varags.
// varags artinya datanya bisa menerima lebih dari satu input, atau anggap saja semacam array.
// apa perbedaannya dengan parameter biasa dengan tipe data array?
// jika parameter tipe array, kita wajib membuat array terlebih dahulu sebelum mengirimkan function.
// jika parameter menggunakan varags, kita bisa langsung mengirim data nya, jika lebih dari satu, cukup gunakan tanda koma ",".

func sumALL(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// function value
// function adalah first class citizen
// function juga merupakan tipe data, dan bisa disimpan di dalam variable.

func getGoodbye(name string) string {
	return "Good Bye friend " + name
}

// function as parameter
// function tidak hanya bisa kita simpan di dalam variable sebagai value.
// namun juga bisa kita gunakan sebagai parameter untuk function lain

// function Type Declarations
// kadang jika function terlalu panjang, agak ribet untuk menuliskannya di dalam parameter
// Type Declarations juga bisa digunakan untuk membuat alias function, sehingga akan mempermudah kita menggunakan function sebagai parameter

type Filter func(string) string

func hellofilter(name string, filter Filter) {
	filteredname := filter(name)
	fmt.Println("hello", filteredname)
}

func spamfilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

// anonymous function
// sebelumnya setiap membuat function, kita akan selalu memberikan sebuah nama pada function tersebut
// namun kadang ada kalanya lebih mudah membuat function secara langsung di variable atau parameter tanpa harus membuat function terlebih dahulu
// hal tersebut dinamakan anonymous function, atau function tanpa nama.

type Blacklist func(string) bool

func registeruser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("you are blocked", name)
	} else {
		fmt.Println("selamat datang", name)
	}
}
