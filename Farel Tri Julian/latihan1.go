package main

import "fmt"

func main() {
	var kg, gram, hargaperkg, jasa int

	fmt.Print("Masukan berat total : ")
	fmt.Scan(&gram)

	hargaperkg = 10000

	kg = gram / 1000
	gram = gram % 1000

	hargakg := kg * hargaperkg

	if kg >= 10 {
		jasa = 0
	} else {
		if gram >= 500 {
			jasa = gram * 5

		} else {
			jasa = gram * 15

		}
	}
	totalharga := hargakg + jasa
	fmt.Printf("Detail berat : %d kg %d gram\n", kg, gram)
	fmt.Printf("biaya :Rp. %d + Rp. %d\n", hargakg, jasa)
	fmt.Printf("Total biaya : %d\n", totalharga)
}
