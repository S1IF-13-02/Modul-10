package main

import "fmt"

func main() {
	var bilanganInput int
	var faktorBilangan []int
	var jumlahFaktor int

	fmt.Println("=== Program Analisis Bilangan Prima dan Faktor ===")
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&bilanganInput)

	// Validasi input
	if bilanganInput <= 0 {
		fmt.Println("❌ Error: Masukkan bilangan bulat positif (> 0)")
		return
	}

	fmt.Println("\n--- Analisis Bilangan ---")
	fmt.Printf("Bilangan: %d\n", bilanganInput)

	// Cari semua faktor bilangan
	fmt.Print("Faktor: ")
	for pembagi := 1; pembagi <= bilanganInput; pembagi++ {
		if bilanganInput%pembagi == 0 {
			fmt.Print(pembagi, " ")
			faktorBilangan = append(faktorBilangan, pembagi)
			jumlahFaktor++
		}
	}
	fmt.Println()

	// Tentukan apakah bilangan prima
	// Bilangan prima memiliki tepat 2 faktor: 1 dan dirinya sendiri
	adalahBilanganPrima := (jumlahFaktor == 2 && bilanganInput > 1)

	// Tampilkan hasil analisis
	fmt.Printf("Jumlah faktor: %d\n", jumlahFaktor)
	fmt.Printf("Bilangan prima: %v\n", adalahBilanganPrima)

	if adalahBilanganPrima {
		fmt.Printf("✓ %d adalah bilangan prima\n", bilanganInput)
	} else {
		fmt.Printf("✗ %d bukan bilangan prima\n", bilanganInput)
		if bilanganInput == 1 {
			fmt.Println("  (1 bukan termasuk bilangan prima)")
		}
	}
}
