package main
import "fmt"
func main() {
    var C, A int
    fmt.Scan(&C)
    fmt.Scan(&A)
    

    C --
    viagem := A / C
    resto := A % C
     if resto == 0 {
        fmt.Println(viagem)
    }   
    if resto != 0 {
     viagem ++
     fmt.Println(viagem) 
    }   

}
