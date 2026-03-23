package main
import "fmt"
func main() {
    var a, b int
    var operador string
    fmt.Scan(&a)
    fmt.Scan(&b)
    fmt.Scan(&operador)

    resultado := 0
    valid := true

    switch operador{
    case "+":
    resultado = a + b 
    case "-":
    resultado = a - b
    case "*": 
    resultado = a * b 
    case "/":
        if b == 0 {
            fmt.Println("ERRO, divisao por zero") 
            valid = false
        } else {
            resultado = a / b
        }
    
         default:
		fmt.Println("Operação inválida!")
		valid = false
         }

	if valid {
		fmt.Println(resultado)
	}
    }
