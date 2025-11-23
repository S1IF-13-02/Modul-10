package main

import (
	"fmt"
)

func main() {
	var bilangan int
	fmt.Print("Bilangan: ")
	fmt.Scan(&bilangan)

	faktorCount := 0
	prima := "false"

	fmt.Print("Faktor: ")

	for f := 1; f <= bilangan; f++ {
		if bilangan%f == 0 {
			fmt.Printf("%d ", f)
			faktorCount++
		}
	}
	fmt.Println()
	if faktorCount == 2 {
		prima = "true"
	}
	fmt.Println("Prima:", prima)
}
