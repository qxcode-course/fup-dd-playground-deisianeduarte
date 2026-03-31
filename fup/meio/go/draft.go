package main
import "fmt"
func main() {
    var a, b, c int64
    fmt.Scan(&a)
    fmt.Scan(&b)
    fmt.Scan(&c)
    if a > b && a < c || a < b && a > c{
    fmt.Println(a)
    }
    if b > a && b < c || b < a && b > c {
        fmt.Println(b)
    } 
    if c > a && c < b || c < a && c > b {
        fmt.Println(c)
    }
}
