package main

import "fmt"

var titles []string
var authors []string
var prices []float64
var counts []int

func main() {
	addBook("Go Programming", "John Doe", 39.99, 10)
	addBook("Learning Python", "Jane Doe", 29.99, 5)
	addBook("Angular Programming", "Bob Smith", 49.99, 20)

	// Menampilkan inventaris buku
	displayInventory()

	// Mengurutkan buku berdasarkan judul
	sortBooksByTitle()

	// Menampilkan inventaris buku setelah diurutkan
	displayInventory()

	// Mencari buku berdasarkan judul
	fmt.Println(findBookByTitle("Go Programming"))

	// Memperbarui jumlah buku
	updateCountBook("Go Programming", 8)
	fmt.Println(findBookByTitle("Go Programming"))

	// Menghitung total nilai inventaris
	total := handleTotalInventory()
	fmt.Printf("Total nilai inventaris: $%.2f\n", total)

	// Membeli buku
	message := buyBook("Go Programming", 3, 120.00)
	fmt.Println(message)

	message = buyBook("Go Programming", 15, 600.00)
	fmt.Println(message)

	message = buyBook("Learning Python", 2, 50.00)
	fmt.Println(message)

	message = buyBook("Java Programming", 1, 40.00)
	fmt.Println(message)
}

func addBook(title, author string, price float64, count int) {
	if price <= 0 || count <= 0 {
					fmt.Println("Harga dan jumlah buku harus positif")
					return
	}
	titles = append(titles, title)
	authors = append(authors, author)
	prices = append(prices, price)
	counts = append(counts, count)
}

func displayInventory() {
	fmt.Println("Inventaris Buku:")
	for i := range titles {
					fmt.Printf("Judul: %s, Penulis: %s, Harga: $%.2f, Jumlah: %d\n", titles[i], authors[i], prices[i], counts[i])
	}
}

func sortBooksByTitle() {
	for i := range titles {
					for j := i + 1; j < len(titles); j++ {
									if titles[i] > titles[j] {
													titles[i], titles[j] = titles[j], titles[i]
													authors[i], authors[j] = authors[j], authors[i]
													prices[i], prices[j] = prices[j], prices[i]
													counts[i], counts[j] = counts[j], counts[i]
									}
					}
	}
}

func findBookByTitle(title string) (string, string, float64, int) {
	for i, t := range titles {
					if t == title {
									return titles[i], authors[i], prices[i], counts[i]
					}
	}
	return "", "", 0, 0
}


func updateCountBook(title string, count int) {
	for i, t := range titles {
		if t == title {
			counts[i] = count
			return
		}
	}
}

func handleTotalInventory() float64 {
	total := 0.0
	for i, c := range counts {
		total += float64(c) * prices[i]
	}
	return total
}

func buyBook(bookTitle string, quantity int, cash float64) string {
	for i, t := range titles {
		if t == bookTitle {
			if counts[i] < quantity {
				return "Jumlah buku yang tersedia tidak mencukupi."
			} else {
				totalPrice := prices[i] * float64(quantity)
				if cash < totalPrice {
					return "Uang yang anda miliki tidak mencukupi untuk membeli buku ini."
				} else {
					counts[i] -= quantity
					return "Pembelian berhasil!"
				}
			}
		}
	}
	return "Buku yang anda cari tidak tersedia, silahkan cari buku yang lain."
}

