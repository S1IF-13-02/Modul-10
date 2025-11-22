package main

import "fmt"

func main() {
	var usia int
	var n bool

	fmt.Print("masukan usia: ")
	fmt.Scan(&usia)
	fmt.Scan(&n)

	if usia >= 17 && n == true {
		fmt.Print("bisa buat KTP")
	} else {
		fmt.Print("tidak bisa buat KTP")
	}
}
