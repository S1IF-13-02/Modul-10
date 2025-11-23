package main

import "fmt"

func main() {
	var b, jumlahFaktor, i, f int
	var faktor [1000]int
	var prima bool

	fmt.Print("Bilangan: ")
	fmt.Scan(&b)

	jumlahFaktor = 0
	for f = 1; f <= b; f++ {
		if b%f == 0 {
			faktor[jumlahFaktor] = f
			jumlahFaktor++
		}
	}
	fmt.Print("Faktor: ")
	for i = 0; i < jumlahFaktor; i++ {
		fmt.Print(faktor[i])
		if i < jumlahFaktor-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
	if jumlahFaktor == 2 {
		prima = true
	} else {
		prima = false
	}
	fmt.Printf("Prima: %t\n", prima)
}