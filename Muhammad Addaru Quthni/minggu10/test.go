package main
import "fmt"

func main()	{
	var usia int
	var bisaBuatKTP bool
	fmt.Print("Masukkan usia Anda: ")
	fmt.Scan(&usia)
	fmt.Scan(&bisaBuatKTP)
	if usia >= 17 && bisaBuatKTP == true {
		fmt.Println("bisa membuat KTP")
	} else {
		fmt.Println("belum bisa membuat KTP")
	}
}