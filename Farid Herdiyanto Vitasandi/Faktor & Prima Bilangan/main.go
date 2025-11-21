package main

import "fmt"

func main() {
	
	var bilangan int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bilangan)

	jumlahFaktor := 0
	fmt.Print("Faktor-faktor dari bilangan ", bilangan, " adalah: ")

	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Printf("%d ", i)
			jumlahFaktor++
		}
	}

	isPrime := (jumlahFaktor == 2)
	fmt.Print("\nPrima? ", isPrime)
}
