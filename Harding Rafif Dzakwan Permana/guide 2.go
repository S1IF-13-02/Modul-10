package main

import "fmt"

func main() {
    var c rune

    fmt.Print("Masukkan satu karakter: ")
    fmt.Scanf("%c", &c)   

    if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {

        if c == 'a' || c == 'i' || c == 'u' || c == 'e' || c == 'o' ||
		   c == 'A' || c == 'I' || c == 'U' || c == 'E' || c == 'O' {

            fmt.Println("vokal")
        } else {
            fmt.Println("konsonan")
        }

    } else {
        fmt.Println("bukan huruf")
    }
}
