package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan Digit bilangan: ")
	fmt.Scan(&n)

	a := n / 1000
	b := (n % 1000) / 100
	c := (n % 100) / 10
	d := n % 10

	if a < b && b < c && c < d {
		fmt.Printf("digit Bilangan Pada %d Tururut Membesar", n)
	} else if a > b && b > c && c > d {
		fmt.Printf("digit Bilangan Pada %d Tururut Mengecil", n)
	} else if a == b && b == c && c == d {
		fmt.Printf("digit Bilangan pada %d Sama", n)
	} else {
		fmt.Printf("digit Bilangan pada %d Tidak Terurut", n)
	}
}
