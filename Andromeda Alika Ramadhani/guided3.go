package main

import "fmt"

func main() {
	var bilangan, ratusan, ribuan, puluhan, satuan int
	fmt.Print("Bilangan: ")
	fmt.Scan(&bilangan)

	ribuan = bilangan / 1000
	ratusan = (bilangan / 100) % 10
	puluhan = (bilangan / 10) % 10
	satuan = bilangan % 10

	if ribuan < ratusan && ratusan < puluhan && puluhan < satuan {
		fmt.Print("Bilangan terurut membesar")
	} else if ribuan > ratusan && ratusan > puluhan && puluhan > satuan {
		fmt.Print("Bilangan terurut mengecil")
	} else {
		fmt.Print("Bilangan tidak terurut")
	}
}
