package main

import "fmt"

func main() {
	var n int
	fmt.Print("masukan angka (bilangan bulat) : ")
	fmt.Scan(&n)
	fmt.Print("Faktor : ")
	prima := true

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Print(i, " ")
			if i == 1 || i == n {
			} else if i > 1 && i < n {
				prima = false
			}

		}
	}
	if n == 1 {
		prima = false
	}
	fmt.Println()
	fmt.Println("Prima : ", prima)
}
