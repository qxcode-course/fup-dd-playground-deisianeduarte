package main
import "fmt"
import "math"
func main() {
    var a, b, c, p, s, at float64
    fmt.Scan(&a, &b, &c)
    p = (a + b + c) / 2
    s = p * (p - a) * (p - b) * (p - c)
    at = math.Sqrt(s)
    fmt.Printf("%.2f\n", at)
}
