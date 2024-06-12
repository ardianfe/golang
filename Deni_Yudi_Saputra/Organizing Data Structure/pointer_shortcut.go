package main

import "fmt"

func main() {
	x := 10

	ptr := &x //untuk mendapatkan memori dari x untuk disimpan di ptr, mendapatkan lokasi tersebut nanti bisa dirubah disitu.  Anda hanya mendapatkan referensi ke lokasi di mana nilai variabel itu sendiri disimpan di memori

	fmt.Println("Alamat memori dari x :", ptr)

	fmt.Println("nilai dari x : ", *ptr)

	// mengubah variable menggunakan pointer

	*ptr = 20 // mengakses pointer untuk merubah nilai dari x yang sebelumnya 10 menggunakan pointer

	fmt.Println("nilai baru dari x setelah diubah : ", x)

	// Pointer shorchur merujuk pada penggunaan operator `&` dan `*` di go

	// Operator & digunakan untuk mendapatkan alamat memori dari suatu variable contohnya  ptr := &x mengambil alamat memori dari variabel x dan menyimpannya dalam variabel ptr.

	// Operator * digunakan untuk deferensi pointer, mengakses nilai yang disimpan di alamat memori yang ditunjuk oleh pointer tersebut. contoh  *ptr mengakses nilai asli dari variabel x.

	// ====GOTCHAS POINTER =====

	var ptrGotchas *int // Inisialisasi pointer, tetapi tidak diberi nilai awal, sehingga secara otomatis nilainya adalah nil.

	if ptr != nil { // Pengecekan apakah pointer tidak nil.
		fmt.Println(*ptrGotchas) // Operasi dereference yang dapat menyebabkan panic jika ptrGotchas nil.
	} else {
		fmt.Println("Pointer is nil") // Pesan yang dicetak jika pointer nil.
	}

	// 	Gotcha dalam contoh tersebut adalah ketika pointer ptrGotchas tidak diinisialisasi secara eksplisit, sehingga secara otomatis nilainya menjadi nil. Namun, kita tidak memeriksa apakah pointer tersebut memiliki nilai nil sebelum melakukan operasi dereference.
	// Dalam kasus ini, ptrGotchas tidak diinisialisasi, sehingga nilainya adalah nil. Kemudian, ketika kita mencoba untuk mencetak nilai yang ditunjuk oleh pointer (*ptrGotchas), hal itu akan menyebabkan panic karena pointer nil tidak menunjuk ke alamat memori yang valid.

	// Jadi, gotcha terjadi karena tidak memeriksa apakah pointer ptrGotchas memiliki nilai nil sebelum melakukan operasi dereference.
}
