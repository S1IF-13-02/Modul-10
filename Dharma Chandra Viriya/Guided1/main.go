package main

import "fmt"

func main() {
	var umur int
	var memilikiKtp bool

	fmt.Print("Masukkan umur & apakah sudah memiliki ktp (true/false): ")
	fmt.Scanln(&umur, &memilikiKtp)
	if umur >= 17 && memilikiKtp {
		fmt.Println("bisa membuat KTP")
	} else {
		fmt.Println("belum bisa membuat KTP")
	}
}
