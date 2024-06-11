// tipe data slice adalah potongan dari data array
// slice mirip dengan array, yang membedakan adalah ukuran slice bisa berubah
// slice dan array selalu terkoneksi, dimana slice adalah data yang mengakses sebagian atau seluruh data di array

// detail tipe data slice
// tipe data slice memiliki 3 data, yaitu pointer, length, dan capacity
// pointer adalah penunjuk data pertama di array para slice
// length adalah panjang dari slice
// capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity

package main

import "fmt"

func main() {
	name := [...]string{"oke", "mari", "kita", "koding", "sampe", "pusing"}
	slice1 := name[4:6]
	slice2 := name[:3]
	slice3 := name[3:]
	slice4 := name[:]

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(cap(slice4))
	fmt.Println(len(slice4))

	// function slice
	// len(slice) = untuk mendapatkan panjang dari slicenya
	// cap(slice) = untuk mendapatkan kapasitas slicenya
	// append(slice, data) = membuat slice baru dengan menambah data ke posisi terakhir slice, jika kapasitas sudah penuh, maka akan membuat array baru
	// make([]TypeData, length, capacity) = membuat slice baru
	// copy(destination, source) = menyalin slice dari source ke destination

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	daySlice1 := days[5:] //sabtu, minggu
	fmt.Println(daySlice1)

	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "libur baru ")
	// array baru akan terbentuk
	//daysbaru := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu", "libur baru"}
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	var newslice []string = make([]string, 2, 5)
	newslice[0] = "anjay"
	newslice[1] = "anjay"
	// newslice[2] = "anjay" //akan menghasilkan error karena panjangnya sudah ditentukan menjadi 2
	// harusnya menggunakan append

	fmt.Println(newslice)
	fmt.Println(len(newslice))
	fmt.Println(cap(newslice))

	newslice2 := append(newslice, "anjay")
	fmt.Println(newslice2)
	fmt.Println(len(newslice2))
	fmt.Println(cap(newslice2))

	newslice2[0] = "oke"
	fmt.Println(newslice2)
	fmt.Println(newslice)

	fromslice := days[:]
	toslice := make([]string, len(fromslice), cap(fromslice))

	copy(toslice, fromslice)

	fmt.Println(fromslice)
	fmt.Println(toslice)

	// perbedaan array dan slice
	iniarray := [...]int{1, 2, 3, 4, 5}
	inislice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniarray)
	fmt.Println(inislice)

}
