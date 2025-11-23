package main

import "fmt"

func main() {
	var x string
	var huruf, kecil, besar bool
	fmt.Scan(&x)
	huruf = (x >= "a" && x <= "z") || (x >= "A" && x <= "Z")
	kecil = x == "a" || x == "i" || x == "u" || x == "e" || x == "o"
	besar = x == "A" || x == "I" || x == "U" || x == "E" || x == "O"
	if huruf && (kecil || besar) {
		fmt.Println("vokal")
	}else 
		if huruf && !(kecil || besar){
			fmt.Println("konsonan")
		}else{
			fmt.Println("bukan huruf")
	}
}