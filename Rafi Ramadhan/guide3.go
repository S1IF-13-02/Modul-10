package main

import (
	"fmt"
)

func main() {
	var bilangan int

	fmt.Print("Masukkan 4 digit bilangan:")
	fmt.Scan(&bilangan)

	if bilangan >= 1000 && bilangan <= 9999 {
		digit1 := bilangan / 1000
		digit2 := (bilangan / 100) % 10
		digit3 := (bilangan / 10) % 10
		digit4 := bilangan % 10

	if digit1 < digit2 && digit2 < digit3 && digit3 < digit4 {
		fmt.Println("Bilangan terurut membesar")
	} else if digit1 > digit2 && digit2 > digit3 && digit3 > digit4 {
		fmt.Println("Bilangan terurut mengecil")
	} else {
		fmt.Println("Bilangan terurut tidak membesar atau mengecil")
	}
	
	}else{
		fmt.Println("Input harus antara 1000-9999")
	}
}
