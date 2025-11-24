package main

import "fmt"

func main() {
	var beratKilogram, beratGram, hargaPerKilogram, biayaJasaTambahan int
	var beratTotalGram int

	fmt.Println("=== Program Perhitungan Biaya Pengiriman Parsel ===")
	fmt.Print("Masukkan berat total parsel (gram): ")
	fmt.Scan(&beratTotalGram)

	// Konstanta harga
	hargaPerKilogram = 10000
	tarifGramDiatas500 := 5
	tarifGramDibawah500 := 15

	// Konversi gram ke kilogram dan sisa gram
	beratKilogram = beratTotalGram / 1000
	beratGram = beratTotalGram % 1000

	// Hitung biaya berdasarkan kilogram
	biayaPerKilogram := beratKilogram * hargaPerKilogram

	// Hitung biaya jasa tambahan untuk sisa gram
	if beratKilogram >= 10 {
		biayaJasaTambahan = 0
		fmt.Println("✓ Gratis biaya tambahan (berat >= 10 kg)")
	} else {
		if beratGram >= 500 {
			biayaJasaTambahan = beratGram * tarifGramDiatas500
			fmt.Printf("✓ Tarif sisa gram: Rp.%d per gram (>= 500 gram)\n", tarifGramDiatas500)
		} else {
			biayaJasaTambahan = beratGram * tarifGramDibawah500
			fmt.Printf("✓ Tarif sisa gram: Rp.%d per gram (< 500 gram)\n", tarifGramDibawah500)
		}
	}

	totalBiayaPengiriman := biayaPerKilogram + biayaJasaTambahan

	// Tampilkan detail perhitungan
	fmt.Println("\n--- Rincian Biaya Pengiriman ---")
	fmt.Printf("Berat parsel    : %d kg + %d gram\n", beratKilogram, beratGram)
	fmt.Printf("Biaya kilogram  : Rp.%d (Rp.%d × %d kg)\n", biayaPerKilogram, hargaPerKilogram, beratKilogram)
	fmt.Printf("Biaya tambahan  : Rp.%d\n", biayaJasaTambahan)
	fmt.Println("--------------------------------")
	fmt.Printf("Total biaya     : Rp.%d\n", totalBiayaPengiriman)
}
