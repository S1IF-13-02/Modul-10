package main

import "fmt"

func main() {
	 
	var beratParsel, sisaBiaya int

	fmt.Print("Masukkan berat parsel (gram): ")
	fmt.Scan(&beratParsel)

	beratKG := beratParsel / 1000
	beratGram := beratParsel % 1000

	biaya := beratKG * 10000

	if beratKG > 10{
		sisaBiaya = 0
	}else if beratGram >= 500 {
		sisaBiaya = beratGram * 5
	}else{
		sisaBiaya = beratGram * 15
	}

	totalBiaya := biaya + sisaBiaya

	fmt.Printf("Detail berat: %d kg + %d gram\n", beratKG, beratGram)
	fmt.Printf("Detail biaya: Rp.%d + Rp.%d\n", biaya, sisaBiaya)
	fmt.Printf("Total biaya pengiriman: Rp.%d\n", totalBiaya)

}

	





