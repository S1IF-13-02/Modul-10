package main

import "fmt"

func main() {
	var beratGram int

	fmt.Print("masukan berat (gram) : ")
	fmt.Scan(&beratGram)

	beratKg := beratGram / 1000
	sisaGram := beratGram % 1000

	biayaKg := beratKg * 10000

	var biayaSisa int

	if beratKg > 10 {
		biayaSisa = 0
	} else {
		if sisaGram >= 500 {
			biayaSisa = sisaGram * 5
		} else {
			biayaSisa = sisaGram * 15
		}
	}
	total := biayaKg + biayaSisa

	fmt.Printf("detail berat: %d kg + %d gr\n", beratKg, sisaGram)
	fmt.Printf("detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaSisa)
	fmt.Printf("total biaya: Rp. %d\n", total)
}
