package main

import (
	"fmt"
)

func main() {
	var bilanganInput int

	fmt.Println("=== Program Analisis Urutan Digit Bilangan ===")
	fmt.Print("Masukkan bilangan 4 digit (1000-9999): ")
	fmt.Scan(&bilanganInput)

	if bilanganInput >= 1000 && bilanganInput <= 9999 {
		// Ekstraksi setiap digit
		digitPertama := bilanganInput / 1000
		digitKedua := (bilanganInput / 100) % 10
		digitKetiga := (bilanganInput / 10) % 10
		digitKeempat := bilanganInput % 10

		fmt.Println("\n--- Hasil Analisis ---")
		fmt.Printf("Bilangan: %d\n", bilanganInput)
		fmt.Printf("Digit-digit: %d, %d, %d, %d\n", digitPertama, digitKedua, digitKetiga, digitKeempat)

		if digitPertama < digitKedua && digitKedua < digitKetiga && digitKetiga < digitKeempat {
			fmt.Println("Pola Urutan: Bilangan terurut membesar (ascending)")
		} else if digitPertama > digitKedua && digitKedua > digitKetiga && digitKetiga > digitKeempat {
			fmt.Println("Pola Urutan: Bilangan terurut mengecil (descending)")
		} else {
			fmt.Println("Pola Urutan: Bilangan tidak terurut")
		}

	} else {
		fmt.Println("\n❌ Error: Input harus berupa bilangan 4 digit (antara 1000-9999)")
	}
}
