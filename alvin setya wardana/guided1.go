package main

import "fmt"

func main() {
	var umur int
	
	fmt.Scan(&umur)

	if umur >= 17 && umur < 20{
		fmt.Println("bisa membuat ktp")
	} else {
		fmt.Println("belum bisa")
	} 
}
	
