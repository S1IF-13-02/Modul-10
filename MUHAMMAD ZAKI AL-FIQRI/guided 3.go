package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukan nilai bilangan : ")
	fmt.Scan(&n)

	if n >= 1000 && n <= 9999 {
		digit1 := n / 1000
		digit2 := (n / 100) % 10
		digit3 := (n / 10) % 10
		digit4 := n % 10

		if digit1 < digit2 && digit2 < digit3 && digit3 < digit4 {
			fmt.Print("Bilangan terurut membesar")
		} else if digit1 > digit2 && digit2 > digit3 && digit3 > digit4 {
			fmt.Print("Bilangan terurut mengecil")
		} else {
			fmt.Print("Bilangan tidak terurut")
		}

	} else {
		fmt.Print("Masukan harus 1000-9999")
	}
}
