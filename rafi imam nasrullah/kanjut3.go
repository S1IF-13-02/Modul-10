package main

import "fmt"

func main() {
    var n int

    fmt.Print("Masukkan bilangan: ")
    fmt.Scan(&n)

    d1 := n / 1000        
    d2 := (n / 100) % 10 
    d3 := (n / 10) % 10  
    d4 := n % 10 

    membesar := (d1 < d2) && (d2 < d3) && (d3 < d4)
    mengecil := (d1 > d2) && (d2 > d3) && (d3 > d4)

    if membesar {
        fmt.Println("Digit pada bilangan", n ,"terurut membesar")
    } else if mengecil {
        fmt.Println("Digit pada bilangan", n ," terurut mengecil")
    } else {
        fmt.Println("Digit pada bilangan", n ,"tidak terurut")
    }
}
