package main

import (
	"fmt"
	"sync"
)

type Akun struct {
	saldo int
	mu    sync.Mutex
}

// Metode untuk mendapatkan saldo
func (a *Akun) getSaldo() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.saldo
}

// Metode untuk setoran
func (a *Akun) setoran(jumlah int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.saldo += jumlah
}

// Metode untuk penarikan
func (a *Akun) penarikan(jumlah int) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.saldo < jumlah {
		return fmt.Errorf("saldo tidak mencukupi")
	}
	a.saldo -= jumlah
	return nil
}

type Transaksi struct {
	jenis  string
	jumlah int
}

func main() {
	akun := &Akun{saldo: 1000}
	transaksiCh := make(chan Transaksi)
	var wg sync.WaitGroup

	// Fungsi untuk memproses transaksi
	prosesTransaksi := func(t Transaksi) {
		defer wg.Done()
		if t.jenis == "setoran" {
			akun.setoran(t.jumlah)
			fmt.Printf("Setoran %d berhasil. Saldo sekarang: %d\n", t.jumlah, akun.getSaldo())
		} else if t.jenis == "penarikan" {
			err := akun.penarikan(t.jumlah)
			if err != nil {
				fmt.Printf("Penarikan %d gagal: %v. Saldo sekarang: %d\n", t.jumlah, err, akun.getSaldo())
			} else {
				fmt.Printf("Penarikan %d berhasil. Saldo sekarang: %d\n", t.jumlah, akun.getSaldo())
			}
		}
	}

	// Goroutine untuk menerima dan memproses transaksi
	go func() {
		for t := range transaksiCh {
			wg.Add(1)
			go prosesTransaksi(t)
		}
	}()

	// Simulasi beberapa klien yang melakukan transaksi
	transaksiCh <- Transaksi{jenis: "setoran", jumlah: 200}
	transaksiCh <- Transaksi{jenis: "penarikan", jumlah: 100}
	transaksiCh <- Transaksi{jenis: "setoran", jumlah: 500}
	transaksiCh <- Transaksi{jenis: "penarikan", jumlah: 700}

	// Menunggu semua transaksi selesai
	close(transaksiCh)
	wg.Wait()

	fmt.Printf("Saldo akhir: %d\n", akun.getSaldo())
}
