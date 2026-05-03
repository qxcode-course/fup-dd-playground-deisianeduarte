package main
import "fmt"
func main() {
    var a int
    fmt.Scan(&a)

    if a > 0 {
        fmt.Println("positivo")
    
    } else if a < 0 {
        fmt.Println("negativo")

    } else {
        fmt.Println("nulo")
}
}
