package main

import (
	"fmt"
)

func main() {
	var berat, biayaTambah int

	fmt.Print("Berat parcel (gram): ")
	fmt.Scan(&berat)

	kg := berat / 1000
	sisa := berat % 1000

	if kg >= 10 {
		sisa = 0
	}

	biayaDasar := kg * 10000

	if sisa >= 500 {
		biayaTambah = sisa * 5
	} else {
		biayaTambah = sisa * 15
	}

	totalBiaya := biayaDasar + biayaTambah

	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, sisa)
	fmt.Printf("Detail biaya: Rp %d + Rp %d\n", biayaDasar, biayaTambah)
	fmt.Printf("Total biaya: Rp %d\n", totalBiaya)
}
