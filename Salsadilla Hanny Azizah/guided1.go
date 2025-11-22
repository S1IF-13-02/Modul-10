package main

import "fmt"

func main() {
	var umur int
	var KK bool

	fmt.Print("Masukkan umur = ")
	fmt.Scan(&umur)
	fmt.Print("Status KK = ")
	fmt.Scan(&KK)

	if umur >= 17 && KK {
		fmt.Print("bisa membuat ktp") 

		} else{
			fmt.Print("tidak bisa membuat ktp") 
		}
	}
		 
