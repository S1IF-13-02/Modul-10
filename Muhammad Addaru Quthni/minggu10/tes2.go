package main
import "fmt"
func main()	{
	var alpa string
	fmt.Print("Masukkan sebuah huruf: ")
	fmt.Scan(&alpa)
	if alpa == "a" || alpa == "i" || alpa == "u" || alpa == "e" || alpa == "o" ||
	   alpa == "A" || alpa == "I" || alpa == "U" || alpa == "E" || alpa == "O" {
		fmt.Println("Vokal")
	} else if (alpa >= "a" && alpa <= "z") || (alpa >= "A" && alpa <= "Z") {
		fmt.Println("Konsonan")
	} else {
		fmt.Println("Bukan huruf")
	}
}