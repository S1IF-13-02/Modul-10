package main

import "fmt"

func main() {
	var b int
	fmt.Print("Bilangan: ")
	fmt.Scan(&b)

	fmt.Print("Faktor: ")

	hitungFaktor := 0

	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i, " ")
			hitungFaktor++
		}
	}

	prima := false
	if hitungFaktor == 2 {
		prima = true
	}

	fmt.Println()
	fmt.Println("Prima:", prima)
}
