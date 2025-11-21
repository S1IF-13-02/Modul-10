package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Bilangan: ")
	fmt.Scan(&n)

	a := n / 1000           
	b := (n / 100) % 10        
	c := (n / 10) % 10        
	d := n % 10              

	if a < b && b < c && c < d {
		fmt.Printf("Digit pada bilangan %d terurut membesar\n", n)
	} else if a > b && b > c && c > d {
		fmt.Printf("Digit pada bilangan %d terurut mengecil\n", n)
	} else {
		fmt.Printf("Digit pada bilangan %d tidak terurut\n", n)
	}
}