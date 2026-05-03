package main
import "fmt"
func main() {
    var n, d, a int
    fmt.Scan(&n)
    fmt.Scan(&d)
    fmt.Scan(&a)

    m := (d - a + n)%n

    fmt.Println(m)
}
