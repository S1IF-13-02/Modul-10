package main
import "fmt"

func main() {
    var parsel, biaya, sisa, total int
    fmt.Print("Berat parsel (dalam gram): ")
    fmt.Scan(&parsel)

    kg := parsel / 1000       
    gram := parsel % 1000     

    biaya = kg * 10000 
    sisa = 0           

    fmt.Println("Detail Berat Parsel : ", kg, "kg", "+", gram, "gram")

    if kg > 10 {
        sisa = 0
    } else {
        if gram >= 500 {
            sisa = gram * 5
        } else {
            sisa = gram * 15
        }
    }

    total = biaya + sisa

    fmt.Println("Detail Biaya: Rp.", biaya, "+", "Rp.", sisa)
    fmt.Println("Total Biaya Pengiriman: Rp.", total)
}