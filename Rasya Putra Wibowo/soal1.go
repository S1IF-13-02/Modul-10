package main

import (
	"fmt"
)

func main() {
	var beratParsel int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&beratParsel)

	kg := beratParsel / 1000
	sisa := beratParsel % 1000

	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, sisa)

	biayaDasar := kg * 10000

	biayaTambahan := 0

	if kg > 10 {
		biayaTambahan = 0
	} else {
		if sisa > 0 {
			if sisa < 500 {
				biayaTambahan = sisa * 5
			} else {
				biayaTambahan = sisa * 15
			}
		}
	}

	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaDasar, biayaTambahan)

	total := biayaDasar + biayaTambahan
	fmt.Printf("Total biaya: Rp. %d\n", total)
}
