package main

import (
	"fmt"
)

func main() {
	var berat int
	fmt.Print("Masukkan Berat: ")
	fmt.Scan(&berat)

	kg := berat / 1000
	sisa := berat % 1000
	BiayaPerKG := kg * 10000
	var BiayaSisa int
	if kg > 10 {
		BiayaSisa = 0
	} else {
		if sisa >= 500 {
			BiayaSisa = sisa * 5
		} else {
			BiayaSisa = sisa * 15
		}
	}
	Total := BiayaPerKG + BiayaSisa
	fmt.Printf("Berat: %d kg %d gram\n", kg, sisa)
	fmt.Printf("Biaya: Rp. %d + Rp. %d\n", BiayaPerKG, BiayaSisa)
	fmt.Printf("Total Biaya: Rp. %d\n", Total)
}
