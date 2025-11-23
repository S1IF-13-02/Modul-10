package main

import "fmt"

func main() {
	var usia int
	var ktp bool

	fmt.Print("masukan usia: ")
	fmt.Scan(&usia)
	fmt.Scan(&ktp)

	if usia >= 17 && ktp {
		fmt.Println("bisa membuat ktp")
	} else {
		fmt.Println(" blm bisa membuat ktp")

	}
}
