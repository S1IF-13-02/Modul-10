package main

import "fmt"

func main() {

	var x string			
	
	fmt.Print("Masukkan huruf vokal/konsonan: ")
	fmt.Scan(&x)

	if len(x) != 1 {
		fmt.Print("Input harus 1 karakter saja")
		return	
	}

	if x == "a" || x == "i" || x == "u" || x == "e" || x == "o" || x == "A" || x == "I" || x == "U" || x == "E" || x == "O" {
		fmt.Println("Huruf Vokal")
	} else if x >= "a" && x <= "z" || x >= "A" && x <= "Z" {
		fmt.Println("Huruf Konsonan")
	}else {
		fmt.Println("Bukan huruf vokal/konsonan")
	}
}
