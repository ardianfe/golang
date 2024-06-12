package maps

import (
	"fmt"
)

func ManipulatingMaps() map[string]int {
	makan := make(map[string]int)

	makan["soto"] = 2000
	makan["ayam"] = 2000
	makan["bakso"] = 3000

	harga := makan["soto"]
	fmt.Printf("Harga makanan soto : %d\n", harga)

	// - Delete map
	fmt.Println("Sebelum di delete :", makan)

	delete(makan, "bakso")

	fmt.Println("Setelah di delete : ", makan)

	return makan

}

func main() {

	// 1.Apa itu Map

	// MAPS adalah type data yang menyimpan key dan value, key yang bersifat uniqe

	// var buah map[string]int //deklarasi, belom diinisialisasi jadi masih nil

	buah := map[string]int{} //inisialisasi
	// buah1 := make(map[string]int)

	buah["anggur"] = 2000 //[key] = value
	buah["apel"] = 3000
	buah["melon"] = 4000

	fmt.Println(buah)
	fmt.Printf("Harga buah melon : %d\n", buah["melon"])

	// 2. Manipulating map

	makan := make(map[string]int)

	// - Tambah / Mengubah nilai
	makan["soto"] = 2000
	makan["bakso"] = 3000

	// - Akses nilai

	harga := makan["soto"]
	fmt.Printf("Harga makanan soto : %d\n", harga)

	// - Delete map
	fmt.Println("Sebelum di delete :", makan)

	delete(makan, "bakso")

	fmt.Println("Setelah di delete : ", makan)

	// 3. Iterating over map atau perulangan map menggunakan for

	for key, value := range buah {
		fmt.Println("Kunci:", key, "Nilai:", value)
	}

	// 4.Perbedaan Maps dan Struct

	// Maps contoh
	mapContoh := make(map[string]string)

	mapContoh["hi"] = "Selamat datang"
	mapContoh["bye"] = "Selamat jalan"

	mapContoh2 := map[string]string{
		"halo": "Yooiiii",
	}
	fmt.Println("===========================")
	fmt.Println(mapContoh)
	fmt.Println(mapContoh2)

	// Struct contoh

	type mapContohStruct struct {
		Name string
		Age  int
	}

	contoh := mapContohStruct{
		Name: "Deni",
		Age:  32,
	}
	fmt.Println(contoh)
}
