package main
import "fmt"
func main() {
    var b int
    fmt.Print("Bilangan: ")
    fmt.Scan(&b)

    fmt.Print("Faktor: ")

    prima := true

    for i := 1; i <= b; i++ {
        if b%i == 0 {
            fmt.Print(i, " ")

            if i == 1 || i == b {
            } else if i > 1 && i < b {
                prima = false
            }
        }
    }

    if b == 1 {
        prima = false
    }

    fmt.Println()
    fmt.Println("Prima:", prima)
}
