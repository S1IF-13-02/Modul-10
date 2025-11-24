package main

import "fmt"

func main() {
	var nilaiAkhirMataKuliah float64
	var nilaiHuruf string
	var keterangan string

	fmt.Println("=== Program Konversi Nilai Mata Kuliah ===")
	fmt.Print("Masukkan nilai akhir mata kuliah (0-100): ")
	fmt.Scan(&nilaiAkhirMataKuliah)

	// Validasi input
	if nilaiAkhirMataKuliah < 0 || nilaiAkhirMataKuliah > 100 {
		fmt.Println("❌ Error: Nilai harus berada dalam rentang 0-100")
		return
	}

	// Konversi nilai angka ke huruf
	if nilaiAkhirMataKuliah > 80 {
		nilaiHuruf = "A"
		keterangan = "Lulus dengan nilai sangat baik"
	} else if nilaiAkhirMataKuliah > 72.5 {
		nilaiHuruf = "AB"
		keterangan = "Lulus dengan nilai baik sekali"
	} else if nilaiAkhirMataKuliah > 65 {
		nilaiHuruf = "B"
		keterangan = "Lulus dengan nilai baik"
	} else if nilaiAkhirMataKuliah > 57.5 {
		nilaiHuruf = "BC"
		keterangan = "Lulus dengan nilai cukup baik"
	} else if nilaiAkhirMataKuliah > 50 {
		nilaiHuruf = "C"
		keterangan = "Lulus dengan nilai cukup"
	} else if nilaiAkhirMataKuliah >= 40 {
		nilaiHuruf = "D"
		keterangan = "Tidak lulus"
	} else {
		nilaiHuruf = "E"
		keterangan = "Tidak lulus"
	}

	// Tampilkan hasil konversi
	fmt.Println("\n--- Hasil Konversi Nilai ---")
	fmt.Printf("Nilai Angka : %.2f\n", nilaiAkhirMataKuliah)
	fmt.Printf("Nilai Huruf : %s\n", nilaiHuruf)
	fmt.Printf("Keterangan  : %s\n", keterangan)
}
