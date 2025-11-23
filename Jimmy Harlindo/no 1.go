package main

import (
	"fmt"
)

func main() {
	var berat int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&berat)

	// Hitung kg dan sisa
	kg := berat / 1000
	sisa := berat % 1000

	// Hitung biaya
	biayaKg := kg * 10000

	var biayaSisa int
	if sisa >= 500 {
		biayaSisa = sisa * 5
	} else {
		biayaSisa = sisa * 15
	}

	total := biayaKg + biayaSisa

	// Output
	fmt.Println("Detail berat :", kg, "kg +", sisa, "gram")
	fmt.Println("Detail biaya : Rp.", biayaKg, "+ Rp.", biayaSisa)
	fmt.Println("Total biaya  : Rp.", total)
}
