package main

import "fmt"

func main() {
	var gram int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&gram)

	kg := gram / 1000
	sisa := gram % 1000

	biayaKg := kg * 10000
	biayaSisa := 0

	if kg <= 10 {
		if sisa >= 500 {
			biayaSisa = sisa * 5
		} else {
			biayaSisa = sisa * 15
		}
	}

	total := biayaKg + biayaSisa

	fmt.Println("Berat:", kg, "kg +", sisa, "gr")
	fmt.Println("Biaya per kg:", biayaKg)
	fmt.Println("Biaya sisa:", biayaSisa)
	fmt.Println("Total:", total)
}
