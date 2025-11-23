package main

import "fmt"

func main() {
	var a int

	fmt.Print("Bilangan: ")
	fmt.Scan(&a)

	fmt.Print("Faktor: ")

	jmlhFaktor := 0

	for i := 1; i <= a; i++ {
		if a%i == 0 {
			fmt.Print(i, " ")
			jmlhFaktor++
		}
	}

	prima := false
	if jmlhFaktor == 2 {
		prima = true
	}

	fmt.Println()
	fmt.Println("Prima:", prima)
}
