package main

import "fmt"

func main() {
	var b int
	fmt.Scan(&b)

	jumlah := 0
	fmt.Print("Faktor: ")

	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i, " ")
			jumlah++
		}
	}

	prima := jumlah == 2
	fmt.Println("\nPrima:", prima)
}
