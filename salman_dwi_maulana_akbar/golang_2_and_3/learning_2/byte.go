package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Membuat byte slice dari string
	str1 := "Hello, World!"
	byteSlice := []byte(str1)
	fmt.Println("Byte slice dari string:", byteSlice)

	// Membuat byte slice langsung
	directByteSlice := []byte{'H', 'e', 'l', 'l', 'o'}
	fmt.Println("Byte slice langsung:", directByteSlice)

	// Mengubah byte slice menjadi string
	str2 := string(byteSlice)
	fmt.Println("String dari byte slice:", str2)

	// Menambahkan elemen ke byte slice
	byteSlice = append(byteSlice, '!')
	fmt.Println("Byte slice setelah ditambahkan:", byteSlice)

	// Mengubah elemen byte slice
	byteSlice[0] = 'h'
	fmt.Println("Byte slice setelah diubah:", byteSlice)

	// Menghapus elemen terakhir byte slice
	byteSlice = byteSlice[:len(byteSlice)-1]
	fmt.Println("Byte slice setelah dihapus:", byteSlice)

	// Menghapus elemen ke-2
	byteSlice = append(byteSlice[:1], byteSlice[2:]...)
	fmt.Println("Byte slice setelah penghapusan:", byteSlice)

	// Menghapus elemen ke-5 (atau yang terakhir, karena sebelumnya sudah dihapus beberapa elemen)
	if len(byteSlice) > 4 {
		byteSlice = append(byteSlice[:4], byteSlice[5:]...)
		fmt.Println("Byte slice setelah penghapusan:", byteSlice)
	} else {
		fmt.Println("Byte slice tidak memiliki elemen ke-5 untuk dihapus")
	}

	// Memeriksa keberadaan file
	filename := "example.txt"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("File tidak ada, membuat file baru:", filename)
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
	} else {
		fmt.Println("File ditemukan:", filename)
	}

	// Membaca file
	fileData, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Isi file dalam byte slice sebelum ditulis:", fileData)

	// Mengubah byte slice menjadi string
	str3 := string(fileData)
	fmt.Println("Isi file dalam string sebelum ditulis:", str3)

	// Data baru untuk ditulis ke file
	newData := []byte("Hello, GoLang!")

	// Menulis byte slice ke file
	err = os.WriteFile(filename, newData, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data berhasil ditulis ke file")

	// Membaca kembali file setelah menulis
	fileData, err = os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Isi file dalam byte slice setelah ditulis:", fileData)

	// Mengubah byte slice menjadi string setelah menulis
	str4 := string(fileData)
	fmt.Println("Isi file dalam string setelah ditulis:", str4)
}
