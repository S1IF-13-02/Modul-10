package main

import (
	"fmt"
	"math"
)
func main() {
	var beratParselGrams int
	
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&beratParselGrams)

	gramsDalamKg := 1000
	biayaPerKg := 10000
	totalKg := int(math.Floor(float64(beratParselGrams) / float64(gramsDalamKg)))
	sisaGram := beratParselGrams % gramsDalamKg
	biayaDasar := totalKg * biayaPerKg
	biayaTambahan := 0

	if beratParselGrams > 10 * gramsDalamKg { 
		fmt.Print("Sisa gram digratiskan jika total > 10kg")
	} else {

		if sisaGram >= 500 {
			biayaPerGram := 5
			biayaTambahan = sisaGram * biayaPerGram
		} else if sisaGram > 0 && sisaGram < 500 {
			biayaPerGram := 15
			biayaTambahan = sisaGram * biayaPerGram
		}
	}

	totalBiaya := biayaDasar + biayaTambahan

	fmt.Printf("Detail berat: %d kg + %d gr\n", totalKg, sisaGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaDasar, biayaTambahan)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}