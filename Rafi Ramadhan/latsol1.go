package main
import "fmt"
func main() {
    var berat, kg, lebih, biayaLebih, biayaKg int
    fmt.Print("Berat parsel (gram): ")
    fmt.Scan(&berat)

    kg = berat / 1000
    lebih = berat % 1000

    biayaKg = kg * 10000
    biayaLebih = 0

    if berat > 10000 {
        biayaLebih = 0
    } else if lebih >= 500 {
        biayaLebih = lebih * 5
    } else if lebih > 0 {
        biayaLebih = lebih * 15
    } else {
        biayaLebih = 0
    }

    total := biayaKg + biayaLebih

    fmt.Println("Detail berat:", kg, "kg +", lebih, "gr")
    fmt.Println("Detail biaya: Rp.", biayaKg, "+ Rp.", biayaLebih)
    fmt.Println("Total biaya: Rp.", total)
}
