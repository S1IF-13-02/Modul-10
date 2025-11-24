package main

import "fmt"

func main() {
	var umurPendaftar int
	var memilikiKartuKeluarga bool

	fmt.Println("=== Program Pengecekan Syarat Pembuatan KTP ===")
	fmt.Print("Masukkan umur: ")
	fmt.Scan(&umurPendaftar)
	fmt.Print("Apakah memiliki Kartu Keluarga? (true/false): ")
	fmt.Scan(&memilikiKartuKeluarga)

	fmt.Println("\n--- Hasil Pengecekan ---")
	if umurPendaftar >= 17 && memilikiKartuKeluarga {
		fmt.Println("Status: Bisa membuat KTP")
		fmt.Printf("Anda berumur %d tahun dan memiliki Kartu Keluarga.\n", umurPendaftar)
	} else {
		fmt.Println("Status: Belum bisa membuat KTP")
		if umurPendaftar < 17 {
			fmt.Println("Alasan: Umur belum mencapai 17 tahun")
		}
		if !memilikiKartuKeluarga {
			fmt.Println("Alasan: Belum memiliki Kartu Keluarga")
		}
	}
}
