package main
import "fmt"
func main() {
    var a, b, c int
    fmt.Scan(&a)
    fmt.Scan(&b)
    fmt.Scan(&c)

    diferenca1 := a - b 
    diferenca2 := a - c 


    if diferenca1 < 0 {
        diferenca1 = diferenca1 * (-1)
    } 
    if diferenca2 < 0 {
        diferenca2 = diferenca2 * (-1)
    }
    if diferenca1 > diferenca2 {
        fmt.Println("segundo")
    } else if diferenca1 < diferenca2 {
        fmt.Println("primeiro")
    } else if diferenca1 == diferenca2 {
        fmt.Println("empate")
    } 
}
