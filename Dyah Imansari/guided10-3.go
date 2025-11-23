package main
import "fmt"
func main() {
	var x, a, b, c, d int
	fmt.Print("Bilangan: ")
	fmt.Scanln(&x)
	a = x / 1000
	b = (x % 1000) / 100
	c = (x % 1000 % 100) / 10
	d = x % 1000 % 100 % 10
	if x >= 1000 && x<= 9999 && a>b && b>c && c>d {
		fmt.Println("Digit pada bilangan", x, "terurut mengecil")
	} else if x >= 1000 && x<= 9999 && a<b && b<c && c<d {
		fmt.Println("Digit pada bilangan", x, "terurut membesar")
	} else {
		fmt.Println("Digit pada bilangan", x, "tidak terurut")
	}
}