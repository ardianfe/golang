package main

import "fmt"

// variable global
var ageGlobal int = 302

func VariableDeclarations() {

	fmt.Println("Variable global :", ageGlobal)

	var name string

	name = "Deni Yudi"

	fmt.Println(name)

	name = "anjay"
	fmt.Println(name)

	middleName := "Slebew"

	fmt.Println(middleName)

	var age, height string = "sepuluh", "seratus"
	fmt.Println("Age:", age, "Height:", height)

	// variable tidak perlu didefinisikan karena golang otomatis tau varuabke tersebut
	var umur = 30
	fmt.Println(umur)

	age1 := 19 // integer ini
	fmt.Println("Age:", age1)

	// variable dengan blok var
	var (
		age2    int    = 30
		height2 int    = 175
		name2   string = "John"
	)

	fmt.Println("Name:", name2, "Age:", age2, "Height:", height2)

	// variable constanta atau yang tidak bisa diubah, wajib initialiasasi langsung datanya
	// tidak berubah lagi tapi kalau tidak digunakan tidak masalah 
	const firstname string = "myke"
	const lastname = "cici"

	// error tidak bisa dirubah
	// firstname= "kikuk"
	// lastname= "kiw"

	fmt.Println(firstname,lastname)

	// variable custom 

	type Age int

    var myAge Age = 30
    fmt.Println("My Age:", myAge)

}
