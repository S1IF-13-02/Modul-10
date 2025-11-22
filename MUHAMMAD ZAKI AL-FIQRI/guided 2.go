package main

import "fmt"

func main() {
	var A string
	fmt.Print("masukan huruf/angka/simbol : ")
	fmt.Scan(&A)
	if A == "a" || A == "i" || A == "u" || A == "e" || A == "o" || A == "A" || A == "I" || A == "U" || A == "E" || A == "O" {
		fmt.Print("Vokal")
	} else if (A >= "b" && A <= "z") || (A >= "B" && A <= "Z") {
		fmt.Print("Konsonan")
	} else {
		fmt.Print("Bukan huruf")
	}
}
