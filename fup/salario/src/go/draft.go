package main
import "fmt"
func main() {
    var salario float64
    fmt.Scan(&salario)

    if salario <= 1000.00 {
        a := salario + (salario * 0.20)
        fmt.Printf("%.2f\n", a)
    } else if salario > 1000.00 && salario <= 1500.00 {
        a := salario + (salario * 0.15)
        fmt.Printf("%.2f\n", a)
    } else if salario > 1500.00 && salario <= 2000.00 {
        a := salario + (salario * 0.10)
        fmt.Printf("%.2f\n", a)
    } else if salario > 2000.00 {
        a := salario + (salario * 0.05)
        fmt.Printf("%.2f\n", a)
    }
}