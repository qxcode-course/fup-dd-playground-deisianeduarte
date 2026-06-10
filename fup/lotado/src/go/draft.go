package main
import "fmt"
func main() {
    var c, m int
    fmt.Scan(&c)
    q := c * 2

    for i := 0; i <= q; i++ {
        if i <= 0 {
            fmt.Println("vazio")
        }
        if i > 0 && i < c {
                fmt.Println("ainda cabe")
        }
        if i >= c && i < q {
                fmt.Println("lotado")
        }
        if i >= q {
                fmt.Println("hora de partir")
        }

        fmt.Scan(&m)
            i = i + m 
        if i < 0 {
            i = 0
        }
        }
    

}
