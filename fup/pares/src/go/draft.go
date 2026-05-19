package main
import "fmt"
func main() {
    var a, b int 
    fmt.Scan(&a, &b)

    soma := 0

    if a > b {
        fmt.Println("invalido")
    } else {
    for i := a; i <= b; i ++ {
        if i % 2 == 0 {
        soma = soma + i
    }
}
    fmt.Println(soma)
}
}