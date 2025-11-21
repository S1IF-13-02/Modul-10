package main

import "fmt"

func main() {

	var bilangan int		
	var teks string

	fmt.Print("Masukkan deret bilangan (1000-9999):")
	fmt.Scan(&bilangan)

	digit1 := bilangan / 1000
	digit2 := (bilangan / 100) % 10
	digit3 := (bilangan / 10) % 10
	digit4 := bilangan % 10


	if digit1 < digit2 && digit2 < digit3 && digit3 < digit4 {
		teks = " terurut membesar"
	} else if digit1 > digit2 && digit2 > digit3 && digit3 > digit4 {
		teks = " terurut mengecil"
	} else {
		teks = " tidak terurut"
	}
	fmt.Print("Digit pada bilangan ", bilangan, teks)

}
