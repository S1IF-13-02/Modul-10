package main

import "fmt"

func main(){
	var angka int
	var d1,d2,d3,d4 int 
	var hasil string
	fmt.Scan(&angka)

	d1 = (angka / 1000)
	d2 = (angka % 1000 / 100)
	d3 = (angka % 100 / 10)
	d4 = (angka % 10)
	
	if d1 <= d2 && d2 <= d3 && d3 <= d4 {
		hasil = "terurut membesar"
	}else
		if d1 >= d2 && d2 >= d3 && d3 >= d4 {
			hasil = "terurut mengecil"
		}else{
			hasil = "tidak terurut"
		}
		fmt.Println("Digit pada bilangan",angka,hasil)

}