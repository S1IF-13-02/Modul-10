package main

import "fmt"

func main() {
	var umur int
	var punyakk bool
	fmt.Print("Masukkan umur: ")
	fmt.Scan(&umur)
	fmt.Print("Punya KK? ")
	fmt.Scan(&punyakk)
	if umur >= 17 && punyakk {
		fmt.Println("bisa membuat KTP")
	} else {
		fmt.Println("belum bisa membuat KTP")
	}
}
