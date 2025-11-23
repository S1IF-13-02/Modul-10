package main

import "fmt"

func main() {
	var huruf string
	fmt.Print("Masukkan huruf: ")
	fmt.Scan(&huruf)
	if huruf == "i" || huruf == "a" || huruf == "u" || huruf == "e" || huruf == "o" ||
		huruf == "I" || huruf == "A" || huruf == "U" || huruf == "E" || huruf == "O" {
		fmt.Print("vokal")
	} else if huruf >= "A" && huruf <= "Z" || huruf >= "a" && huruf <= "z" {
		fmt.Print("konsonan")
	} else {
		fmt.Print("bukan huruf")
	}
}
