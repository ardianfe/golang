package main

import "fmt"

func cariKey(m map[string]int, keyi string) {
	for key, value := range m {
		if key == keyi {
			fmt.Printf("key %s exists with value %d\n", key, value)
			return
		}
	}
	fmt.Printf("key %s does not exist", keyi)
}

func main() {
	m := map[string]int{ // Deklarasi sekaligus inisialisasi dengan nilai
		"one": 1, //menambah nilai
		"two": 2,
	}
	value := m["one"]  // Mengakses elemen dengan kunci "one"
	fmt.Println(value) // Output: 10

	cariKey(m, "one")

	// alternative sintaks untuk mencari key
	// value, exists := m["onee"]
	// if exists {
	// 	fmt.Printf("key 'one' exists with value: %d", value)
	// } else {
	// 	fmt.Printf("key 'one' does not exist")
	// }

	delete(m, "one") // menghapus elemen yang ada di dalam map
	for key, value := range m {
		fmt.Printf("key: %s, value: %d\n", key, value)
	}
}
