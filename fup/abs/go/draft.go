package main
import "fmt"
func main() {
    var a, b, d int 
    fmt.Scan(&a)
    fmt.Scan(&b)
    d = a - b
    if d <0 {
    fmt.Println(d * (-1))
    } else {
    fmt.Println(d)
    }
}