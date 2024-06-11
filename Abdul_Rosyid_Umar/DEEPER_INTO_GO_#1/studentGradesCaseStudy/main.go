package main

import (
	"fmt"
	"studentgrades/utils"
)

func main() {
	//menghitung nilai rata rata siswa
	scores := []float64{84.6, 92.7, 89.3, 94.4}
	ratarata := utils.Hitungratarata(scores)
	fmt.Printf("Rata-rata nilai siswa %.2f\n", ratarata)

	//mengconvert celcius ke fahrenheit
	convcel := utils.Celstofahrenheit(45)
	fmt.Printf("45 derajat celcius menjadi %.2f derajat fahrenheit\n", convcel)
}
