package main

import "fmt"

func main() {

	var umur int 		
	var kartuKeluarga bool

	fmt.Print("Masukkan umur anda: ")
	fmt.Scan(&umur)
	fmt.Print("Apakah punya kartu keluarga? (true/false): ")
	fmt.Scan(&kartuKeluarga)

	if umur >= 17 && kartuKeluarga == true {
		fmt.Println("Anda bisa membuat KTP")
	} else {
		fmt.Println("Anda tidak bisa membuat KTP")
	}
}

