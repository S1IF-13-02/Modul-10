package main

import "fmt"

func main() {
    var berat, biaya, total int

    fmt.Print("Berat Parsel (Gram) = ")
    fmt.Scan(&berat)

    kg := berat / 1000
    sisa := berat % 1000

    biayaKg := kg * 10000

    if sisa >= 500 {
        biaya = sisa * 5
    } else {
        biaya = sisa * 15
    }

    if kg > 10 {
        total = biayaKg
    } else {
        total = biayaKg + biaya
    }

    fmt.Printf("Detail berat: %d kg + %d gr\n", kg, sisa)
    fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biaya)
    fmt.Printf("Total biaya: Rp. %d\n", total)
}
