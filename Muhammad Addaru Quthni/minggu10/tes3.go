package main

import "fmt"

func main() {
	var angka int
	fmt.Print("Masukkan angka 4 digit: ")
	fmt.Scan(&angka)

	if angka < 0 {
		angka = -angka
	}

	if angka >= 1000 && angka <= 9999 {
		a := angka / 1000
		b := (angka / 100) % 10
		c := (angka / 10) % 10
		d := angka % 10

		if a < b && b < c && c < d {
			fmt.Println("digit pada angka", angka, "terurut membesar")
		} else if a > b && b > c && c > d {
			fmt.Println("digit pada angka", angka, "terurut mengecil")
		} else {
			fmt.Println("digit pada angka", angka, "tidak terurut")
		}
	} else {
		fmt.Println("Angka tidak terdiri dari 4 digit")
	}
}