package main
import "fmt"
func main() {
    var c, b, g, m, frutas int
    fmt.Scan(&c)
    fmt.Scan(&b)
    fmt.Scan(&g)
    fmt.Scan(&m)

    frutas = b + g + m
    viagem := frutas / c
 
    if frutas %c != 0 {
    viagem ++
  }
    fmt.Println(viagem)
}
