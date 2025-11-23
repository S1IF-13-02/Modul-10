package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)
	a := n / 1000
	b := (n / 100) % 10
	c := (n / 10) % 10
	d := n % 10
	fmt.Printf("Digit pada bilangan %d ", n)
	if a < b && b < c && c < d {
		fmt.Println("terurut membesar")
	} else if a > b && b > c && c > d {
		fmt.Println("terurut mengecil")
	} else {
		fmt.Println("tidak terurut")
	}
}
