package main

import "fmt"

func main() {
	var usia int
	var ktp bool

	fmt.Print("Masukan Usia : ")
	fmt.Scan(&usia)
	fmt.Print("Bisa buat Ktp atau tidak? ")
	fmt.Scan(&ktp)
	
	if usia >= 17 && ktp {
		fmt.Print("Bisa buat ktp")
	}else{
		fmt.Print("Belum bisa buat ktp")
	}
}