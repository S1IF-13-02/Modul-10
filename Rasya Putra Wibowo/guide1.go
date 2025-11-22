package main 
import "fmt" 

func main() {
var usia int
var punyaKK bool 
fmt.Scan(&usia) 
fmt.Scan(&punyaKK)
if usia >= 17 && punyaKK {
fmt.Println("bisa buat KTP")
} else {
fmt.Println("belum bisa membuat ktp")
}
}
