package main

import "fmt"

func main() {
	var parsel int
	var sisa int
	fmt.Print("Masukkan berat parsel: ")
	fmt.Scan(&parsel)
	sisa = parsel % 1000
	if parsel < 10000 && sisa >= 500 {
		berat := (parsel - sisa) + sisa
		biaya := (parsel-sisa)*10 + sisa*5
		fmt.Println("Detail berat: ", berat)
		fmt.Println("Total biaya: ", biaya)
	} else if parsel < 10000 && sisa < 500 {
		berat := (parsel - sisa) + sisa
		biaya := (parsel-sisa)*10 + sisa*15
		fmt.Println("Detail berat: ", berat)
		fmt.Println("Total biaya: ", biaya)
	} else if parsel > 10000 && sisa < 1000 {
		berat := (parsel - sisa) + sisa
		biaya := (parsel - sisa) * 10
		fmt.Println("Detail berat: ", berat)
		fmt.Println("Total biaya: ", biaya)
	}

}
