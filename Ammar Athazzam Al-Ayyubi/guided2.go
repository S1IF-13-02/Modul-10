package main

import (
	"fmt"
	"strings"
)

func main() {
	var n string
	fmt.Scanln(&n)
	if len(n) != 1 {
		fmt.Println("bukan huruf")
		return
	}

	n = strings.ToLower(n)
	if n >= "a" && n <= "z" {
		if n == "a" || n == "i" || n == "u" || n == "e" || n == "o" {
			fmt.Println("vokal")
		} else {
			fmt.Println("konsonan")
		}
	} else {
		fmt.Println("bukan huruf")
	}
}
