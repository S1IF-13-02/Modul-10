package main

import "fmt"
func main (){
	var usia int
	var kk bool 
	fmt.Print("Masukan Usia : ")
	fmt.Scan(&usia)
	fmt.Print("Apakah mempunyai KK?, (masukan True/False) : ")
	fmt.Scan(&kk)

	if usia >= 17 && kk == true {
		fmt.Print("Bisa Membuat ktp")
	}else {
		fmt.Print("Belum bisa membuat ktp")
	}
}