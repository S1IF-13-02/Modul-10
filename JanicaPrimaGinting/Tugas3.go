package main

import "fmt"

func main() {
    var b int
    fmt.Print("Masukkan bilangan bulat: ")
    fmt.Scan(&b)
    fmt.Println("Bilangan:",b)

    jumlahFaktor := 0
    fmt.Print("Faktor: ")
    for i := 1; i <= b; i++ {
        if b % i == 0 {
            fmt.Print(i, " ")
            jumlahFaktor++
        }
    }
    fmt.Println()

    if jumlahFaktor == 2 {
        fmt.Println("Prima:",true)
    } else {
        fmt.Println("Prima:",false)
    }
}