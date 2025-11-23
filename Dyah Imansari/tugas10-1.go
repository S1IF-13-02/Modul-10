package main
import "fmt"
func main() {
	var b_g, b_kg, sisa int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scanln(&b_g)
	b_kg = b_g / 1000
	sisa = b_g % 1000
	fmt.Println("Detail berat:", b_kg, "kg +", sisa, "gr")
	if b_kg > 10 {
		sisa = 0
		fmt.Println("Total biaya: Rp.", b_kg*10000 + sisa)
	} else if sisa >= 500 {
		fmt.Println("Detail biaya: Rp.", b_kg*10000, "+ Rp.", sisa*5)
		fmt.Println("Total biaya: Rp.", b_kg*10000 + sisa*5)
	} else {
		fmt.Println("Detail biaya: Rp.", b_kg*10000, "+ Rp.", sisa*15)
		fmt.Println("Total biaya: Rp.", b_kg*10000 + sisa*15)
	}
}