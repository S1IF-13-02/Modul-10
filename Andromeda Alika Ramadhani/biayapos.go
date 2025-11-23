package main

import "fmt"

func main() {
	var berat, total_berat, kg, sisa, jasa_pengiriman, biaya_tambahan, total_biaya int
	fmt.Print("Berat parsel: ")
	fmt.Scan(&berat)

	kg = berat / 1000
	sisa = berat % 1000
	total_berat = kg + sisa/1000

	jasa_pengiriman = kg * 10000

	biaya_tambahan = 0
	if total_berat > 10 {
		biaya_tambahan = 0
	} else if sisa >= 500 {
		biaya_tambahan = sisa * 5
	} else {
		biaya_tambahan = sisa * 15
	}

	total_biaya = jasa_pengiriman + biaya_tambahan

	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, sisa)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", jasa_pengiriman, biaya_tambahan)
	fmt.Printf("Total Biaya: Rp. %d\n", total_biaya)
}
