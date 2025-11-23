package main

import "fmt"

func main() {
	var b int
	var isPrima bool = true

	fmt.Print("Bilangan: ")
	fmt.Scan(&b)

	counter := 0
	fmt.Print("Faktor: ")
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i, " ")
			counter++
		}
	}
	fmt.Println()

	if counter != 2 {
		isPrima = false
	}

	fmt.Printf("Prima: %t\n", isPrima)
}
