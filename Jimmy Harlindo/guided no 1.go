package main
import "fmt"

func main(){
	var usia int
	var punyaKK bool

	fmt.Scan(&usia,&punyaKK)
	

	if usia >= 17 && punyaKK == true{
		fmt.Println("bisa membuat KTP")
	} else {
		fmt.Println("belum bisa membuat KTP")
	}

}