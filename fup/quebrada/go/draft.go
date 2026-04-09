package main
import "fmt"
func main() {
    var n1, n2 int
    fmt.Scan(&n1, &n2)
    r1 := n1/n2
    r2 := n1 % n2
    r3 := float64(n1) / float64(n2)
    fmt.Println(r1)
    fmt.Println(r2)
    fmt.Printf("%.2f\n", r3)
}
