package main
import "fmt"
func main() {
    var n int 
    fmt.Scan(&n)

    i := 0
    for i*i < n {
        i++
    }

    if i*i == n {
        fmt.Println("sim")   
    } else {
        fmt.Println("nao")
    }
}