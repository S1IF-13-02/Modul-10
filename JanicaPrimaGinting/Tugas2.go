package main
import "fmt"
func main() {
	var nam float64
	var nmk string
	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scan(&nam)
	if nam > 80 {
		nmk = "A"
	}
	if 72.5 < nam && nam <= 80 {
		nmk = "AB"
	}
	if 65 < nam && nam <= 72.5 {
		nmk = "B"
	}
	if 57.5 < nam && nam <= 65 {
		nmk = "BC"
	}
	if 50 < nam && nam <= 57.5 {
		nmk = "C"
	}
	if 40 < nam && nam <= 50 {
		nmk = "D"
	}else
		if nam <= 40{
		nmk = "E"
	}
	fmt.Println("Nilai mata kuliah:",nmk)
}