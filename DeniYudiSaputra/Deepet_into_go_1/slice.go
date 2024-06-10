package main

import "fmt"

// SLICE adalah potongan dari array, tapi yang membedakan ukuran slice bisa berubah, Slice dan Array selalu terkoneksi, dimana Slice adalah data yang mengakses sebagian atau seluruh data di Array

// atribut slice
// Pointer: Merujuk ke elemen pertama dari array yang dapat diakses melalui slice.
// Length: Jumlah elemen dalam slice atau panjang.
// Capacity: Jumlah elemen dari elemen pertama slice sampai elemen terakhir array yang mendasarinya/ jumlah kapasitas dari slicenya .

// ===== SLICE PADA ARRAY  ======
// ketika mengambil
// array [low:high] = index low ke high
// array [low:] = index low sampai index akhir di array
// array [:high] = index 0 sampai sebelum highnya
// array [:] = index 0 sampai index akhir array (semua)

//  contoh array[
// 	Index 0: Januari
// 	Index 1: Februari
// 	Index 2: Maret
// 	Index 3: April
// 	Index 4: Mei
// 	Index 5: Juni
// 	Index 6: Juli
// 	Index 7: Agustus
// 	Index 8: September
// 	Index 9: Oktober
// 	Index 10: November
// 	Index 11: Desember
// ]

//  contoh slice
// array[4:7] = berarti pointer 4 (mei),length = 3 (4-7 jadi 4,5,6 jadi 3= mei, juni, juli), Capacity jadi 8 karena penunjuk pointer sampai akhir yaitu mei sampai akhir yaiu desember

func main() {

	name := [...]string{"Deni", "Yudi", "Saputra", "Joko", "Samudro", "Nugroho"}

	slice1 := name[4:6]
	fmt.Println(slice1)

	slice2 := name[:3]
	fmt.Println(slice2)

	slice3 := name[3:]
	fmt.Println(slice3)

	// bedanya dengan array yaitu kita tidak menentukan ukurannya seperti array [3] jadi kalau slice tidak ada, mau bikin slice yang sudah ada
	var slice4 []string= name[:]
	fmt.Println(slice4)

	// APPEND SLICE 

	days := [...] string {"Senin","Selasa","Rabu","Kamis","Jum'at","Sabtu","Minggu"}

	daySlice1 := days[5:]//Sabtu Minggu
	fmt.Println(daySlice1)
	// timpa value slice [5:] menggunakan daySlice1[0] maka index tersebut ditimpa
	daySlice1[0] ="Sabtu Baru "
	daySlice1[1] ="Minggu Baru "

	fmt.Println(daySlice1)
	fmt.Println(days)

	// append jika tidak bisa menampung maka akan membuat slice baru nanti dibelakang akan ditambah hari baru
	daySlice2 := append(daySlice1, "Hari Baru")
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	// make new slice 

	newSlice := make([]string,2,5)
	newSlice[0] = "Eko"
	newSlice[1] = "Eko"
	// newSlice[2] = "Eko"// error harusnya menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))//panjang
	fmt.Println(cap(newSlice))//capacity

	newSlice2 := append(newSlice, "Deni")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	// copy slice 

	fromSlice :=days[:]
	toSlice := make([]string,len(fromSlice),cap(fromSlice))

	copy(toSlice,fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// beda array dan slice 

	// GOLANG BIASANYA MEMAKAI SLICE 
	iniArray := [...]int{0,1,2,3,4,5}
	iniSlice := []int{0,1,2,3,4,5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
