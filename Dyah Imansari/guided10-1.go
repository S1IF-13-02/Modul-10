package main
import "fmt"
func main() {
	var u int
	var tes bool
	fmt.Print("Usia = ")
	fmt.Scanln(&u)
	fmt.Print("ada KK? ")
	fmt.Scanln(&tes)
	if u >= 17 && tes == true {
		fmt.Println("bisa membuat KTP")
	} else {
		fmt.Println("belum bisa membuat KTP")
	}
}