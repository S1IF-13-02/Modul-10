package main

import "fmt"

func main() {
	var beratParsel int

	fmt.Print("Berat Parsel (gram): ")
	fmt.Scan(&beratParsel)

	if beratParsel <= 0 {
		fmt.Println("Berat parsel tidak boleh kurang atau sama dengan 0 (gram)")
		return
	}

	detailBeratParselKg, detailBeratParselGr := beratParsel/1000, beratParsel%1000
	detailBiayaKg, detailBiayaGr := detailBeratParselKg*10000, 0

	if detailBeratParselGr >= 500 {
		detailBiayaGr = detailBeratParselGr * 5
	} else {
		detailBiayaGr = detailBeratParselGr * 15
	}

	if beratParsel%1000 != 0 {
		fmt.Printf("Detail berat: %d kg + %d gr\n", detailBeratParselKg, detailBeratParselGr)
		fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", detailBiayaKg, detailBiayaGr)
	} else {
		fmt.Printf("Detail berat: %d kg\n", detailBeratParselKg)
		fmt.Printf("Detail biaya: Rp. %d\n", detailBiayaKg)
	}

	totalBiaya := 0

	if detailBeratParselKg < 10 {
		totalBiaya += detailBiayaKg + detailBiayaGr
	} else {
		totalBiaya += detailBiayaKg
	}

	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
