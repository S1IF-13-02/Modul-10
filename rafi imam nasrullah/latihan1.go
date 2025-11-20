package main

import "fmt"

func main(){
	var umur int
	var kk bool

	fmt.Print("masukan umur : ")
	fmt.Scan(&umur)
	
	fmt.Print("true/false : ")
	fmt.Scan(&kk)

	if umur >= 17 && kk == true {
		fmt.Print("bisa membuat ktp")
	}else {
		fmt.Print("belum bisa membuat ktp")
	}
}