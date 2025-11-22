package main

import "fmt"

func main() {
	var n string

	fmt.Print("masukan apa aja: ")
	fmt.Scan(&n)

	if n == "a" || n == "i" || n == "u" || n == "e" || n == "o" {
		fmt.Print("vokal")
	} else if (n >= "b" && n <= "z") || (n >= "B" && n < "z") {
		fmt.Print("konsonan")
	} else {
		fmt.Print("bukan huruf")
	}

}