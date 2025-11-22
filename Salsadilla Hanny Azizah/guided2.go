package main

import "fmt"

func main() {
	var huruf string

	fmt.Print("Masukkan = ")
	fmt.Scan(&huruf)

	if huruf < "A" || huruf > "z" {
    fmt.Println("bukan huruf")

	} else if huruf == "A" || huruf == "I" || huruf == "U" || huruf == "E" || huruf == "O" || huruf == "a" || huruf == "i" || huruf == "u" || huruf == "e" || huruf == "o"{
    fmt.Println("vokal")
	
	} else {
    fmt.Println("konsonan")
	}

}