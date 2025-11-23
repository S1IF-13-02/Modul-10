package main

import "fmt"

func main() {
	var usia int
	var KK bool
	fmt.Print("Masukkan usia dan Memiliki KK Atau TIdak(true/false): ")
	fmt.Scan(&usia, &KK)


	if usia >= 17 && KK {
		fmt.Println("bisa membuat ktp")
	} else {
		fmt.Println("tidak bisa membuat ktp")
	}
}
