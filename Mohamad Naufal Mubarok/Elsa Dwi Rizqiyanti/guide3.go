package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Bilangan: ")
	fmt.Scanln(&n)

	d1 := n / 1000
	d2 := (n / 100) % 10
	d3 := (n / 10) % 10
	d4 := n % 10

	if d1 < d2 && d2 < d3 && d3 < d4 {
		fmt.Printf("Digit pada bilangan %d terurut membesar\n", n)
	} else if d1 > d2 && d2 > d3 && d3 > d4 {
		fmt.Printf("Digit pada bilangan %d terurut mengecil\n", n)
	} else {
		fmt.Printf("Digit pada bilangan %d tidak terurut\n", n)
	}
}
