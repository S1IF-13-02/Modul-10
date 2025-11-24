package main

import (
	"fmt"
	"strings"
)

func main() {
	var inputKarakter string

	fmt.Println("=== Program Pengecekan Huruf Vokal/Konsonan ===")
	fmt.Print("Masukkan satu karakter: ")
	fmt.Scan(&inputKarakter)

	// Konversi ke lowercase untuk pengecekan yang lebih mudah
	karakterLower := strings.ToLower(inputKarakter)

	fmt.Println("\n--- Hasil Analisis ---")
	if karakterLower == "a" || karakterLower == "i" || karakterLower == "u" || karakterLower == "e" || karakterLower == "o" {
		fmt.Printf("Karakter '%s' adalah: Huruf Vokal\n", inputKarakter)
	} else if karakterLower >= "a" && karakterLower <= "z" {
		fmt.Printf("Karakter '%s' adalah: Huruf Konsonan\n", inputKarakter)
	} else {
		fmt.Printf("Karakter '%s' adalah: Bukan Huruf\n", inputKarakter)
	}
}
