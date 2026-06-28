package main
import "fmt"
func main() {
    var n int
    fmt.Scan (&n)

    largura := 2*n - 1

    for i := 1; i <= n; i++ {
        copos := i
        pontosEsq := n - i 
        
            for j := 0; j < pontosEsq; j++ {
                fmt.Print(".")
            }

            for j := 0; j < copos; j++ {
                fmt.Print(n)
                if j != copos-1 {
                    fmt.Print(".")
                }
            }
            usados := pontosEsq + copos + (copos - 1)
            for j := usados; j < largura; j++ {
                fmt.Print(".")
            }
            fmt.Println()
    }
}