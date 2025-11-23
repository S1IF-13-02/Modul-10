package main

import "fmt"

func main() {
	var n int
	fmt.Print("Bilangan: ")
	fmt.Scan(&n)

	fmt.Print("Faktor: ")

	hitung := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Print(i, " ")
			hitung++
		}
	}

	fmt.Println()

	if hitung == 2 {
		fmt.Println("Prima: true")
	} else {
		fmt.Println("Prima: false")
	}
}
