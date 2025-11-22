package main

import "fmt"

func main() {
	var n int

	fmt.Print("masukan 4 angka: ")
	fmt.Scan(&n)

	if n >= 1000 && n <= 9999 {
		d1 := n / 1000
		d2 := (n / 100) % 10
		d3 := (n / 10) % 10
		d4 := n % 10

		if d1 < d2 && d2 < d3 && d3 < d4 {
			fmt.Print("bilangan terurut membesar")
		} else if d1 > d2 && d2 > d3 && d3 > d4 {
			fmt.Print("bilangan terurut mengecil")
		} else {
			fmt.Print("bilangan tidak terurut")

		}
	} else {
		fmt.Print("masukan antara 1000-9999")
	}
}
