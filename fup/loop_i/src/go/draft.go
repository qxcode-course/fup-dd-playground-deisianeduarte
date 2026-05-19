package main
import "fmt"
func main() {
    var a, b int
    fmt.Scan(&a)
    fmt.Scan(&b)

    for i := a; i < b; i ++ {
        fmt.Print(i,"\n")
    }
}
