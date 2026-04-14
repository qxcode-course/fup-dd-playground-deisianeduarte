package main
import "fmt"
func main() {
    var C, A int
    fmt.Scan(&C)
    fmt.Scan(&A)
    
viagem := C / A 
resto := C % A

     if viagem == 0 {
        fmt.Println(viagem)
    }   
    if resto != 0 {
     viagem ++
    }   
    fmt.Println(viagem) 
}
