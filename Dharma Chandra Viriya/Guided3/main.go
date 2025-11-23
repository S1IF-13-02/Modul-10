package main

import "fmt"

func main() {
	var n int
	var tampung []int
	var menaik bool = true
	var menurun bool = true

	fmt.Print("Masukkan n: ")
	fmt.Scan(&n)

	tampung = append(tampung, n%10)

	for i := len(fmt.Sprintf("%d", n)) - 1; i >= 1; i-- {
		n = n / 10
		nTampung := n % 10

		tampung = append(tampung, nTampung)
	}

	for i := 0; i < len(tampung)-1; i++ {
		if tampung[i] < tampung[i+1] {
			menaik = false
		}
		if tampung[i] > tampung[i+1] {
			menurun = false
		}
	}

	if menaik {
		fmt.Println("Digit terurut membesar")
	} else if menurun {
		fmt.Println("Digit terurut mengecil")
	} else {
		fmt.Println("Digit tidak berurut")
	}

}
