package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	randomInt := rand.Intn(10) // random buat integer

	fmt.Printf("Hasil acak : %d \n", randomInt)

	randomFloat := rand.Float64()

	fmt.Printf("Hasil acak float : %f \n", randomFloat)

	min := 10
	max := 15

	randomRange := make([]int, 0)
    for i := min; i <= max; i++ {
        randomRange = append(randomRange, i)
    }

	
	randomRangeInt := rand.Intn(max-min+1)+min

	fmt.Printf("Hasil range dari %d sampai %d : %d",min,max,randomRangeInt)

	// rand.Seed(time.Now().UnixNano())

	// randomInt := rand.Intn(100)
	// fmt.Printf("Angkat bulat acak : %d\n", randomInt)

	// randomFloat := rand.Float64()
	// fmt.Printf("Angkat pecahan acak : %f\n", randomFloat)

	// min := 10
	// max := 20

	// randomInRange := rand.Intn(max-min+1) + min
	// fmt.Printf("Angka bulat antara %d dan %d : %d\n", min, max, randomInRange)
}
