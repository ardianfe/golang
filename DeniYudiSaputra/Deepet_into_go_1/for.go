package main

import "fmt"

func main() {
	counter := 1

	for counter <=10{
		fmt.Println("Perulangan ke ",counter)
		counter++
	}

	for i := 0; i <= 20; i++ {
		fmt.Println("Perulangan ", i)
	}

	// ==== for range ===== 
	// untuk data collection seperti array,map 

	names := []string{"deni","yudi","saputra"}
	
	// ===SLICE===== 
	for i := 0; i < len(names); i++ {
		fmt.Println("coba :",names[i])
	}

	// ===== FOR RANGE ======
	// UNTUK NGAMBIL INDEX 
	for index, name:=range names{
		fmt.Println("index", index, "=", name)
	}

	// GAK NGAMBIL INDEX GANTI DENGAN UNDESCORE"_" 

	for _, name:=range names{
		fmt.Println(name)
	}
}

