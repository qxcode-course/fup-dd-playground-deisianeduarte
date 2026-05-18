package main
import "fmt"
func main() {
    var a, b, c int
    fmt.Scan(&a)
    fmt.Scan(&b)
    fmt.Scan(&c)

    if a - b > a - c {
        fmt.Println("segundo")
    } else if a - b < a - c {
        fmt.Println("terceiro")
    } else {
        fmt.Println("empate")
    }
}
