package main

import (
	"fmt"
)

func main() {
	var huruf string

	fmt.Print("masukan huruf: ")
	fmt.Scan(&huruf)

	if len(huruf) == 0 {
		fmt.Println("Tidak ada input")
		return
	}

	r := []rune(huruf)[0]

	if r >= 'A' && r <= 'Z' {
		r = r + ('a' - 'A') 
	}

	switch r {
	case 'a', 'i', 'u', 'e', 'o':
		fmt.Println("vokal")
	default:
		if (r >= 'a' && r <= 'z') {
			fmt.Println("konsonan")
		} else {
			fmt.Println("bukan huruf")
		}
	}
}
