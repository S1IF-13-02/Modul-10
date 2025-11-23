package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Print("Masukkan Alfabet: ")
	fmt.Scan(&input)

	alf := strings.ToUpper(input)

	if alf == "A" || alf == "I" || alf == "U" || alf == "E" || alf == "O" {
		fmt.Println("vokal")
	} else if alf >= "A" && alf <= "Z" {
		fmt.Println("konsonan")
	} else {
		fmt.Println("bukan huruf")
	}
}
