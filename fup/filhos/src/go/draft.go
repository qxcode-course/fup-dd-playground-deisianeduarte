package main
import "fmt"
func main() {
  var idade, qntd int
  fmt.Scan(&idade, &qntd)
  
  soma:= idade
  fmt.Println(soma)

  for i := 1; i < qntd; i++ {
    soma = soma + 2
    fmt.Println(soma)
  }

}