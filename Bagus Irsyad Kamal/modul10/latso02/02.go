package main
import "fmt"

func main() {
    var nam float64
    var nmk string

    fmt.Print("Masukkan nilai akhir mata kuliah (NAM): ")
    fmt.Scan(&nam)

    if nam > 88 {
        nmk = "A"
    } else if nam > 72.5 {
        nmk = "AB"
    } else if nam > 65 {
        nmk = "B"
    } else if nam > 57.5 {
        nmk = "BC"
    } else if nam > 50 {
        nmk = "C"
    } else if nam >= 48 {
        nmk = "D"
    } else {
        nmk = "E"
    }

    fmt.Println("Nilai Mutu Kuliah (NMK):", nmk)
}
