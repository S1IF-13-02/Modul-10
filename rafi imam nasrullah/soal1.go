package main

import (
	"fmt"
)

func main() {
	var totalGram int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&totalGram)

	tarifPerKg := 10000

	kg := totalGram / 1000
	sisaGram := totalGram % 1000

	biayaKg := kg * tarifPerKg
	biayaSisa := 0

	if totalGram > 10000 {
		biayaSisa = 0
	} else {
		if sisaGram >= 500 {
			biayaSisa = sisaGram * 5
		} else {
			biayaSisa = sisaGram * 15
		}
	}

	totalBiaya := biayaKg + biayaSisa

	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, sisaGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaSisa)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
