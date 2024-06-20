package main

import (
	"fmt"
	"os"
)

// CustomWriter adalah tipe baru yang mengimplementasikan io.Writer
type CustomWriter struct {
	// Properti tambahan bisa ditambahkan di sini jika diperlukan
	file *os.File
}

// // Write mengimplementasikan metode io.Writer untuk menulis data
// func (c *CustomWriter) Write(p []byte) (int, error) {
// 	// Contoh: Menulis data ke os.Stdout (konsol)
// 	n, err := os.Stdout.Write(p)
// 	if err != nil {
// 		return n, fmt.Errorf("gagal menulis: %v", err)
// 	}
// 	return n, nil
// }

func NewCustomWriter(filename string) (*CustomWriter, error) {
    file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
    if err != nil {
        return nil, fmt.Errorf("gagal membuka file: %v", err)
    }
    return &CustomWriter{file: file}, nil
}

// Write mengimplementasikan metode io.Writer untuk menulis data ke file
func (c *CustomWriter) Write(p []byte) (int, error) {
    n, err := c.file.Write(p)
    if err != nil {
        return n, fmt.Errorf("gagal menulis ke file: %v", err)
    }
    return n, nil
}

// Close menutup file setelah selesai digunakan
func (c *CustomWriter) Close() error {
    if c.file != nil {
        err := c.file.Close()
        if err != nil {
            return fmt.Errorf("gagal menutup file: %v", err)
        }
    }
    return nil
}


func main() {
	// Membuat instance dari CustomWriter
	// cw := &CustomWriter{}

	cw, err := NewCustomWriter("output.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer cw.Close()

	// Contoh penggunaan CustomWriter sebagai io.Writer
	data := []byte("Halo, ini data untuk ditulis yaaaaaaaaaaaaaaaa\n")
	bytesWritten, err := cw.Write(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Jumlah byte yang berhasil ditulis: %d\n", bytesWritten)
}
