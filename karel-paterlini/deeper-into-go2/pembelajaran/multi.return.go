package main

import (
	"fmt"
	"math"
)

// func kalkulasi  memiliki parameter [hitung] diameter lingkaran 
//di dalam fungsi terdapat operasi perhitungan nilai luas dan keliling dari nilai hitung
//kedua hasilnya dijadikan sebagai return value


// func kalkulasi sudah kita sederhanakan dari 
// func calculate(d float64) (area float64, circumference float64) {

// kenapa bisa begitu? 
// karena dideklarasikan memiliki 2 variable yang tipe data nya sama,
// mengembalikan nilai cukup dengan memanggil return tanpa perlu diikuti variabel apapun.
// Nilai terakhir area dan around sebelum pemanggilan keyword return adalah hasil dari fungsi di atas.

// menjadi di bawah seperti ini 
func kalkulasi (hitung float64) (float64, float64) { //Cara pendefinisian banyak nilai balik bisa dilihat pada kode di samping,
// langsung tulis tipe data semua nilai balik dipisah tanda koma, lalu ditambahkan kurung di antaranya.
	
	
	// Menghitung luas lingkaran
	// Rumus luas: π * (jari-jari)^2, di mana jari-jari = diameter / 2
	var area = math.Pi * math.Pow(hitung / 2, 3)

	// Menghitung keliling lingkaran
	// Rumus keliling: π * diameter
	var arround = math.Pi * hitung

	// return 2 nilai 
	return area, arround
// dibagian return juga sama kita membalikan semua data dengan pemisah koma ,
}


func main() {
	// Mendefinisikan variabel diameter dengan tipe float64 dan nilai 20
	var diameter float64 = 20

	// Memanggil fungsi kalkulasi dengan parameter diameter dan menyimpan hasilnya ke dalam variabel area dan arround
	var area, arround = kalkulasi(diameter)

	// Menampilkan hasil luas lingkaran dengan format dua angka desimal
	fmt.Printf("hasil luas\t : %.2f \n", area)

	// Menampilkan hasil keliling lingkaran dengan format dua angka desimal
	fmt.Printf("hasil keliling \t : %.2f \n", arround)
}



// Penggunaan Fungsi math.Pow()
//Fungsi math.Pow() digunakan untuk operasi pangkat nilai. math.Pow(2, 3) berarti 2 pangkat 3, hasilnya 8. Fungsi ini berada dalam package math.

//◉ Penggunaan Konstanta math.Pi
//math.Pi adalah konstanta bawaan package math yang merepresentasikan Pi atau 22/7.