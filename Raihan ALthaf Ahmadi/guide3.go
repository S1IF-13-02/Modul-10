package main

import "fmt"
func main (){
	var x, y int
	fmt.Print("Masukan Angka : ")
	fmt.Scan(&x)

	rill := x

	y = x / 1000
	x = x % 1000

	z := x / 100
	x = x % 100

	q := x / 10
	x = x % 10

	r := x / 1
	x = x % 1
	
	if y > z && z > q && q > r {
		fmt.Print("Digit pada bilangan ",rill , " turut mengecil ")
	}else if y < z && z < q && q < r {
		fmt.Print("Digit pada bilangan ",rill, " turut membesar ")
	}else {
		fmt.Print("Digit Pada Bilangan ",rill, " tidak terurut")
	}
}