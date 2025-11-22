package main

import (
    "fmt"
)

func main() {
    var berat int
    fmt.Print("Berat parsel (gram): ")
    fmt.Scan(&berat)

    kg := berat / 1000
    sisa := berat % 1000

    fmt.Printf("Detail berat: %d kg + %d gr\n", kg, sisa)


    biayaKg := kg * 10000

    biayaSisa := 0

    if berat <= 10000 {
        if sisa > 0 {
            if sisa >= 500 {
                biayaSisa = sisa * 5
            } else {
                biayaSisa = sisa * 15
            }
        }
    } else {
        biayaSisa = 0
    }

    fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaSisa)
    fmt.Printf("Total biaya: Rp. %d\n", biayaKg+biayaSisa)
}
