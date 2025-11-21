package main

import "fmt"

func main() {
	var usia int
	var status bool

	fmt.Scanln(&usia)
	fmt.Scanln(&status)

	if usia >= 17 && status {
		fmt.Println("bisa membuat KTP")
	} else {
		fmt.Println("belum bisa membuat KTP")
	}
}
