package main
import "fmt"
func main() {
    var n1, n2, nf int
    fmt.Scan(&n1)
    fmt.Scan(&n2)
    fmt.Scan(&nf)

    media := (n1 + n2) / 2
    mf := (media+nf) / 2

    if media >= 7 {
        fmt.Println("aprovado")
    } else if media < 4 {
        fmt.Println("reprovado")
    } else if media > 4 && media < 7 && mf >= 5 {
        fmt.Println("aprovado na final")
    } else {
        fmt.Println("reprovado na final")
    }
}
