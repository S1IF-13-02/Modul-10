package main

import "fmt"

func main() {
    var berat int          
    var biayaKg int
    var biayaSisa int
    var kg, sisa int

    fmt.Print("Masukkan berat parsel (gram): ")
    fmt.Scan(&berat)

    kg = berat / 1000
    sisa = berat % 1000

    biayaKg = kg * 10000

    biayaSisa = 0
    if berat > 10000 {
        biayaSisa = 0
    } else {
        if sisa >= 500 {
            biayaSisa = sisa * 5
        } else {
            biayaSisa = sisa * 15
        }
    }

    total := biayaKg + biayaSisa

    fmt.Println("Berat parsel (gram):", berat)
    fmt.Printf("Detail berat: %d kg + %d gr\n", kg, sisa)
    fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaSisa)
    fmt.Printf("Total biaya: Rp. %d\n", total)
}
