package main

import (
	"fmt"
)

func main() {
	var b int
	fmt.Print("Bilangan: ")
	fmt.Scan(&b)

	fmt.Print("Faktor: ")
	
	faktorCount := 0
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i, " ")
			faktorCount++
		}
	}
	fmt.Println()

	isPrima := (faktorCount == 2)

	fmt.Println("Prima:", isPrima)
}