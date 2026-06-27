package main
import "fmt"
func main() {
    var n int 
    fmt.Scan(&n)

    contador := 0
    
    for i := 1; i <= n; i++ {
        if i*i == n {
            contador++
        }
    }

    if contador > 0 {
        fmt.Println("sim")   
    } else {
        fmt.Println("nao")
    }
}