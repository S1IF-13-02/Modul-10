package main

import "fmt"

func main() {
    var berat int
    fmt.Print("Berat parsel (gram): ")
    fmt.Scan(&berat)

    kg := berat / 1000
    sisa := berat % 1000

    biayaKG := kg * 10000
    biayaSisa := 0

    if kg > 10 {
        biayaSisa = 0
    } else {
        if sisa >= 500 {
            biayaSisa = sisa * 5
        } else {
            biayaSisa = sisa * 15
        }
    }

    totalBiaya := biayaKG + biayaSisa

    fmt.Printf("Detail berat: %d kg + %d gr\n", kg, sisa)
    fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKG, biayaSisa)
    fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
