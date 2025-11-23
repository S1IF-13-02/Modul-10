package main

import "fmt"

func main() {
	var alfabet string

	fmt.Scan(&alfabet)

	if (alfabet >= "a" && alfabet <= "z") || (alfabet >= "A" && alfabet <= "Z") {
		if alfabet == "a" || alfabet == "A" ||
			alfabet == "i" || alfabet == "I" ||
			alfabet == "u" || alfabet == "U" ||
			alfabet == "e" || alfabet == "E" ||
			alfabet == "o" || alfabet == "O" {
			fmt.Println("vokal")
		} else {
			fmt.Println("konsonan")
		}

	} else {
		fmt.Println("bukan huruf")
	}
}
