package main

import "fmt"

func main() {
	var umur int
	var b bool

	fmt.Print("Masukan umur : ")
	fmt.Scan(&umur)
	fmt.Print(" masukan true/false : ")
	fmt.Scan(&b)
	if umur >= 17 && b == true {
		fmt.Print("bisa membuat KTP")
	} else {
		fmt.Print("tidak bisa membuat KTP")
	}

}
