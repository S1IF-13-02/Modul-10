package main

import "fmt"

func main() {
	var bulat, i int
	var prima bool
	var faktor int = 0

	fmt.Print("Bilangan: ")
	fmt.Scan(&bulat)

	for i = 1 ; i <= bulat ; i++{
		if  bulat % i == 0 {
			fmt.Print(i, " ")
			faktor++
		} 
	}
	fmt.Println()
	prima = faktor != 2
	fmt.Println(prima)
}