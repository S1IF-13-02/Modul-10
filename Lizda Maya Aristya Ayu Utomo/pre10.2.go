package main

import "fmt"

func main() {
	var alfabet string

	fmt.Scan(&alfabet)

	if (alfabet >= "A" && alfabet <= "Z") || (alfabet >= "a" && alfabet <= "z") {
		if alfabet == "A" || alfabet == "a" ||
			alfabet == "I" || alfabet == "i" ||
			alfabet == "U" || alfabet == "u" ||
			alfabet == "E" || alfabet == "e" ||
			alfabet == "O" || alfabet == "o" {
			fmt.Println("vokal")
		} else {
			fmt.Println("konsonan")
		}

	} else {
		fmt.Println("bukan huruf")
	}
}
