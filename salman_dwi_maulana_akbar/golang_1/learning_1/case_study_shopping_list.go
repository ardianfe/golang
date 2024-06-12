package main

import (
	"fmt"
	"sort"
)

var shoppingList []string

func addItem(item string) {
	for _, listItem := range shoppingList {
		if listItem == item {
			fmt.Println("Item", item, "telah tersedia di daftar belanjaan, silahkan masukkan daftar belanjaan yang lain!")
			return
		}
	}
	shoppingList = append(shoppingList, item)
	fmt.Println(item, "Telah ditambahkan ke daftar belanjaan")
}

func deleteItem(item string) {
	for index, listItem := range shoppingList {
		if listItem == item {
			shoppingList = append(shoppingList[:index], shoppingList[index+1:]...)
			fmt.Println("Item", item, "telah dihapus dari daftar belanjaan")
			return
		}
	}
	fmt.Println("Item", item, "tidak ditemukan dalam daftar belanjaan")
}

func updateItem(oldItem, newItem string) {
	for index, listItem := range shoppingList {
		if listItem == oldItem {
			shoppingList[index] = newItem
			fmt.Println("Item", oldItem, "telah diubah menjadi", newItem)
			return
		}
	}
	fmt.Println("Item", oldItem, "tidak ditemukan dalam daftar belanjaan")
}

func viewList() {
	fmt.Println("Daftar belanjaan:")
	if len(shoppingList) == 0 {
		fmt.Println("Belum ada item di daftar belanjaan, silahkan tambahkan daftar belanjaan anda")
	} else {
		for index, item := range shoppingList {
			fmt.Println(index+1, ". ", item)
		}
	}
}

func searchItem(item string) {
	found := false
	for index, listItem := range shoppingList {
		if listItem == item {
			fmt.Printf("%s ditemukan pada indeks %d\n", item, index+1)
			found = true
		}
	}
	if !found {
		fmt.Printf("%s tidak ditemukan dalam daftar belanjaan\n", item)
	}
}

func sortList() {
	sort.Strings(shoppingList)
	fmt.Println("Daftar belanjaan telah diurutkan:")
	viewList()
}

func main() {
	var choice int
	var item, newItem string
	var oldItem string

	for {
		fmt.Println("\nMenu Pilihan:")
		fmt.Println("1. Tambahkan item")
		fmt.Println("2. Hapus item")
		fmt.Println("3. Ubah item")
		fmt.Println("4. Lihat daftar belanjaan")
		fmt.Println("5. Cari item")
		fmt.Println("6. Urutkan daftar belanjaan")
		fmt.Println("7. Keluar")
		fmt.Print("Pilihan Anda: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Masukkan item baru: ")
			fmt.Scanln(&item)
			addItem(item)
		case 2:
			fmt.Print("Masukkan nama item yang ingin dihapus: ")
			fmt.Scanln(&item)
			deleteItem(item)
		case 3:
			fmt.Print("Masukkan nama item yang ingin diubah: ")
			fmt.Scanln(&oldItem)
			fmt.Print("Masukkan item baru: ")
			fmt.Scanln(&newItem)
			updateItem(oldItem, newItem)
		case 4:
			viewList()
		case 5:
			fmt.Print("Masukkan item yang ingin dicari: ")
			fmt.Scanln(&item)
			searchItem(item)
		case 6:
			sortList()
		case 7:
			fmt.Println("Terima kasih telah menggunakan aplikasi daftar belanjaan!")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih menu yang sesuai.")
		}
	}
}
