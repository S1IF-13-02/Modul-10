package main

import "fmt"

func main() {
	var umur int
	var ktp bool
	fmt.Print("masukan umur")
	fmt.Scan(&umur)

	fmt.Print("memiliki ktp atau tidak?")
	fmt.Scan(&ktp)

	if umur >= 17 && ktp == true {
		fmt.Print("bisa membuat ktp")
	} else {
		fmt.Print("tidak bisa membuat ktp")
	}
}
