package main

import "fmt"

func main() {
	var a int
	fmt.Print("Bilangan: ")
	fmt.Scan(&a)

	fmt.Printf("factor: ")
	jumlahFaktor := 0
	for i := 1; i <= a; i++ {
		if a%i == 0 {
			fmt.Print(i, " ")
			jumlahFaktor++
		}
	}
	fmt.Println()

	if jumlahFaktor == 2 {
		fmt.Printf("Prima: True")
	} else {
		fmt.Printf("Prima: False")
	}
}
