package main 

import "fmt"

func main(){
	var gram,bSisa int
	fmt.Print("Masukkan berat parsel: ")
	fmt.Scan(&gram)
	fmt.Println("Berat parsel (gram):", gram)
	kg := gram / 1000
	sGram := gram % 1000
	bPerkg := kg * 10000
	fmt.Println("Detail berat:", kg ,"kg +", sGram, "gr")
	if sGram >= 500 {
		bSisa = sGram * 5
		fmt.Println("Detail biaya: Rp.", bPerkg ,"+ Rp.", bSisa)
	}else{
		bSisa = sGram * 15 
		fmt.Println("Detail biaya: Rp.", bPerkg ,"+ Rp.", bSisa)
	}
	if kg > 10 {
		bSisa = 0
	}
	tBiaya := bPerkg + bSisa

	fmt.Println("Total biaya: Rp.", tBiaya)
}