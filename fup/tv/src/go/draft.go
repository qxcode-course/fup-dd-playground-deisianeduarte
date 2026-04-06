package main
import "fmt"
func main() {
    var valor float64
    var parcelas float64
    fmt.Scan(&valor, &parcelas)
    valor *= (1 + 0.05*(parcelas-1))
    cada := valor / parcelas
    fmt.Printf("%.2f\n", parcelas)
    fmt.Printf("%.2f\n", cada)
}
