package main

import (
	"fmt"
)

// func Pembagian(angka1, angka2 int) (int, error) {
// 	if angka2 == 0 {
// 		return 0, errors.New("gabisa dibagi dengan 0 cuy")
// 	}
// 	return angka1 / angka2, nil
// }

// custom error

// buat struct buat nampung error
type ShortNamaError struct {
	Name    string
	Message string
}

func (e *ShortNamaError) Error() string {
	return fmt.Sprintf("ShortNameError: %s (name:%s)", e.Message, e.Name)
}

// fungsi validasi untuk menampilkan error jika nama terlalu pendek kurang dari 3 lenght nya

func ValidateName(name string) (string, error) {
	var jeneng = &ShortNamaError{
		Name:    name,
		Message: "nama terlalu pendek",
	}
	if len(name) <= 3 {
		return jeneng.Name, fmt.Errorf(jeneng.Message)
	}

	return jeneng.Name, nil
}

func main() {
	// error adalah interface harus ada object interface error
	// hasil, err := Pembagian(100, 0)
	// if err == nil {
	// 	fmt.Println("Hasil", hasil)
	// } else {

	// 	fmt.Println("Error", err.Error())
	// }

	// panggil untuk cek validate error
	nama, err := ValidateName("Tit")
	if err != nil {
		fmt.Println("Error woi", err, nama)
	} else {
		fmt.Println("Nama valid", nama)
	}

}
