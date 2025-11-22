package main

import "fmt"

func main() {
	var usia int
	var kk bool

	fmt.Print("Masukkan usia: ")
	fmt.Scan(&usia)

	fmt.Print("Apakah punya KK? (true/false): ")
	fmt.Scan(&kk)

	if usia >= 17 && kk {
		fmt.Println("bisa membuat KTP")
	} else {
		fmt.Println("belum bisa membuat KTP")
	}
}
