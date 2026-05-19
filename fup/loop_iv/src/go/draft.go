package main
import "fmt"
func main() {
    var a, b int
    fmt.Scan(&a, &b)

        fmt.Print("[ ")
    if a > b {
        for i := a; i > b; i -- {
            fmt.Print(i, " ")
        }
    } 
        if a < b {
        for i := a; i < b; i ++ {
            fmt.Print(i, " ")
        }
    }
    fmt.Println("]")
}
