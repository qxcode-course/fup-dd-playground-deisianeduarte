package main

import (
	"fmt"
)
func main() {
    var dia, hora int
    fmt.Scan(&dia)
    fmt.Scan(&hora)

     if dia == 1 {
        fmt.Println("NAO")
    } else if dia >= 2 && dia < 7 && hora >= 8 && hora <= 11 {
        fmt.Println("SIM") 
    } else if dia >= 2 && dia < 7 && hora >= 14 && hora <= 17 {
        fmt.Println("SIM")
    } else if dia == 7 && hora >= 8 && hora <= 11 {
        fmt.Println("SIM")
    } else if dia == 7 && hora > 11 {
        fmt.Println("NAO")
    } else {
        fmt.Println("NAO")
    }
}