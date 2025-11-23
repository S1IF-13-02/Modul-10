package main
import "fmt"
func main() {
	var x string
	fmt.Print("Masukkan huruf = ")
	fmt.Scanln(&x)
	if x == "A" || x == "a" || x == "I" || x == "i" || x == "U" || x == "u" || x == "E" || x == "e" || x == "O" || x == "o" {
		fmt.Println("vokal")
	} else if (x >= "A" && x <= "Z") || (x >= "a" && x <= "z") {
		fmt.Println("konsonan")
	} else {
		fmt.Println("bukan huruf")
	}
}