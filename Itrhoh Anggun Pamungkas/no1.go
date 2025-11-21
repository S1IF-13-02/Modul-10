package main
import "fmt"

func main() {
	var berat, kg, sisa int
	var biayaPerKg, biayaSisa, total int

	biayaPerKg = 10000

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&berat)

	kg = berat / 1000
	sisa = berat % 1000

	fmt.Println("Detail berat:", kg, "kg +", sisa, "gr")

	if kg > 10 {
		biayaSisa = 0
	} else {
		if sisa >= 500 {
			biayaSisa = sisa * 5
		} else {
			biayaSisa = sisa * 15
		}
	}

	total = kg*biayaPerKg + biayaSisa

	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", kg*biayaPerKg, biayaSisa)
	fmt.Println("Total biaya: Rp.", total)
}