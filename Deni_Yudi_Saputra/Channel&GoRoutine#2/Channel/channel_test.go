package channel

import (
	"fmt"
	"testing"
	"time"
)

// channel <- data = mengirim data ke channel
// data <- channel = mengambil/menerima data dari channel
// close() wajib untuk menutup channel ketika tidak digunakan

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	// anonymouse function
	// bikin channel harus ada yang mengirim dan menerima, apabila hanya mengirim tidak ada yang menerima maka kan ngehang dan di block, apabila ada yang menerima tidak ada yang mengirim maka akan deadlock. jika tidak ada salah satunya maka akan error

	// jadi apabila pada channel tidak ada yang menerima/mengambil maka goroutine akan diblock,ngehang,diem sampai data yang dikirim ada yang mengambil. error mengirim channel pada yang close

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Deni Yudi Saputra" //mengirim
		fmt.Println("Selesai mengirim data ke channel")
	}()

	// mengambil data tapi tidak dikirim
	data := <- channel //menerima
	fmt.Println(data)


	time.Sleep(5 * time.Second)
}
