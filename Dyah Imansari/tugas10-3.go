package main
import "fmt"
func main() {
	var b, i int
	var p bool = true
	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&b)
	fmt.Print("Faktor: ")
	for i = 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i, " ")
			if i != 1 && i != b {
				p = false
			}
		}
	}
	fmt.Println()
	fmt.Println("Prima:", p)
}