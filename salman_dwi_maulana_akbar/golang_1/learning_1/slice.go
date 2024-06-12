package main

import (
	"fmt"
	"sort"
)

// Fungsi untuk mengurutkan slice menggunakan algoritma Merge Sort
func mergeSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }
    mid := len(arr) / 2
    left := mergeSort(arr[:mid])
    right := mergeSort(arr[mid:])
    return merge(left, right)
}

// Fungsi helper untuk menggabungkan dua slice yang sudah terurut
func merge(left, right []int) []int {
    result := make([]int, 0, len(left)+len(right))
    for len(left) > 0 || len(right) > 0 {
        if len(left) == 0 {
            return append(result, right...)
        }
        if len(right) == 0 {
            return append(result, left...)
        }
        if left[0] <= right[0] {
            result = append(result, left[0])
            left = left[1:]
        } else {
            result = append(result, right[0])
            right = right[1:]
        }
    }
    return result
}

// Fungsi untuk memfilter angka ganjil dari slice
func filterOddNumbers(arr []int) []int {
    result := make([]int, 0)
    for _, num := range arr {
        if num%2 != 0 {
            result = append(result, num)
        }
    }
    return result
}

// Fungsi untuk menampilkan slice
func printSlice(arr []int) {
    fmt.Println("Slice:", arr)
}

func main() {
    sort_numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    subSlice := sort_numbers[2:5] // Mengambil elemen dari indeks 2 hingga 4 (5 tidak termasuk)
    fmt.Println(subSlice)

     // Membuat slice dengan angka acak
     numbers := []int{12, 4, 7, 3, 9, 2, 18, 6}

     // Mengurutkan slice menggunakan Merge Sort
     sortedNumbers := mergeSort(numbers)
     fmt.Println("Setelah diurutkan:")
     printSlice(sortedNumbers)
 
     // Memfilter angka ganjil dari slice
     oddNumbers := filterOddNumbers(numbers)
     fmt.Println("Angka Ganjil:")
     printSlice(oddNumbers)
 
     // Menambahkan beberapa angka ke slice
     numbers = append(numbers, 11, 15, 8)
     fmt.Println("Setelah ditambahkan angka baru:")
     printSlice(numbers)
 
     // Mengurutkan kembali slice setelah penambahan angka baru
     sort.Ints(numbers)
     fmt.Println("Setelah diurutkan kembali:")
     printSlice(numbers)
}
