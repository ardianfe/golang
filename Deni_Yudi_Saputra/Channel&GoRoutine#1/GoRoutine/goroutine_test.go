package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func HelloWord() {
	fmt.Println("Hellooooo dunia")
}


// Jadi goroutine ini berjalan secara asyncronous jadi eksekusinya tidak menunggu proses sebelumnya selesai apa belom,
// Misal proses goroutinenya belom selesai maka akan men skip dan akan berlanjut ke proses selanjutnya, 
// Untuk menggunakan goroutine hanya perlu menambahkan go didepan fungsi nya 
// Dalam contoh ini, kita menambahkan time.Sleep untuk memastikan bahwa goroutine memiliki cukup waktu untuk mengeksekusi sebelum program utama selesai.
func TestGoroutine(t *testing.T) {
	go HelloWord()
	fmt.Println("Upss")

	time.Sleep(1*time.Minute) // untuk memastikan bahwa ketika goroutine belum selesai mengeksekus, aplikasi tidak berhenti duluan 
}

func DisplayNumber(number int)  {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T){
	for i := 0; i < 1000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
