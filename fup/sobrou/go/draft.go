package main
import "fmt"
func main() {
    var a, b, c, w, y, z, h float64
     fmt.Scan(&a)
     fmt.Scan(&b)
     fmt.Scan(&c)
     fmt.Scan(&w)
     fmt.Scan(&y)
     fmt.Scan(&z)
     fmt.Scan(&h)
    fmt.Printf("%.2f\n", ((a*w) + (b*y) + (c*z) - h)*(-1))

}
