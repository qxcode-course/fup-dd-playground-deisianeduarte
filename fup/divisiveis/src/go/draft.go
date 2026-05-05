package main
import "fmt"
func main() {
    var a, b int
fmt.Scan(&a)
fmt.Scan(&b)

if a % 2 == 1 && b % 2 == 1 {
    fmt.Println("sim")
} else {
    fmt.Println("nao")
}
}
