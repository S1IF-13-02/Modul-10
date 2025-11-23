package main

import "fmt"

func main() {
	var usia int
	var kk bool
	fmt.Print("USIA: ")
	fmt.Scan(&usia)
	fmt.Print("Status KK: ")
	fmt.Scan(&kk)
	if usia >= 17 && kk {
		fmt.Print("Bisa membuat KTP")
	} else {
		fmt.Print("Tidak bisa membuat KTP")
	}
}
