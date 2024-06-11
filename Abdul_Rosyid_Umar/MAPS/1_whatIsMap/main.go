package main

import "fmt"

func main() {
	var m map[string]int                                // Deklarasi map tanpa inisialisasi
	m = make(map[string]int)                            // Inisialisasi map
	m["umurkaka"] = 50                                  //inisialisasi dengan value
	fmt.Printf("umur Kaka adalah: %d\n", m["umurkaka"]) // mengakses nilai map dengan key umurkaka

	m2 := make(map[string]int) // Deklarasi sekaligus inisialisasi
	m2["umurandi"] = 18        // menambahkan elemen
	m2["umurbavi"] = 24
	fmt.Printf("umur Bavi adalah: %d\n", m2["umurbavi"])
	m2["umurbavi"] = 70 //Memperbarui umur Bavi
	fmt.Printf("umur Bavi setelah diperbarui adalah: %d\n", m2["umurbavi"])

	m3 := map[string]int{ //deklarasi sekaligus inisialisasi dengan nilai
		"umursule":   18,
		"umurdadang": 29,
	}
	value := m3["umursule"]
	fmt.Printf("umur sule adalah: %d\n", value)
}
