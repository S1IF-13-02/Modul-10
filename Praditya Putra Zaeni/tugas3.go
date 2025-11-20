package main

import "fmt"

func main() {
    var b int
    fmt.Print("Bilangan: ")
    fmt.Scan(&b)

    fmt.Print("Faktor: ")

    jumlahFaktor := 0
    for i := 1; i <= b; i++ {
        if b%i == 0 {
            fmt.Print(i, " ")
            jumlahFaktor++
        }
    }

    prima := false
    if jumlahFaktor == 2 {
        prima = true
    }

    fmt.Printf("\nPrima: %v\n", prima)
}
