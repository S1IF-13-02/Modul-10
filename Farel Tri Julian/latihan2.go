package main

import "fmt"

func main() {
	var nam float64
	var nmk string
	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scan(&nam)
	if nam > 80 {
		nmk = "A"
	} else if nam > 72.5 {
		nmk = "AB"
	} else if nam > 65 {
		nmk = "B"
	} else if nam > 57.5 {
		nmk = "BC"
	} else if nam > 50 {
		nmk = "B"
	} else if nam > 40 {
		nmk = "C"
	} else {
		nmk = "D"
	}
	fmt.Println("Nilai mata kuliah: ", nmk)
}
