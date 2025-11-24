package main
import "fmt"

func main() {
	var b int
	var f int = 1
	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scan(&b)
	fmt.Print("Faktor: ")
	for i := 1; i <= b; i++ {
		if b % i == 0 {
			fmt.Print(i, " ")
			f++
	}
}
	fmt.Println()
	fmt.Print("Prima : ")
	if f == 3 {
		fmt.Print("True")
	} else {
		fmt.Print("False")
	}
}