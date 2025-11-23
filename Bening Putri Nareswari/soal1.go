package main

import "fmt"

func main() {
	var beratGram int
	fmt.Scan(&beratGram)

	beratKg := beratGram / 1000
	sisaGram := beratGram % 1000

	biayaDasar := beratKg * 10000
	biayaSisa := 0

	if beratKg <= 10 {
		if sisaGram >= 500 {
			biayaSisa = sisaGram * 5
		} else {
			biayaSisa = sisaGram * 15
		}
	}

	totalBiaya := biayaDasar + biayaSisa

	fmt.Printf("Detail berat: %d kg + %d gr\n", beratKg, sisaGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaDasar, biayaSisa)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
