package main
import "fmt"
func main() { // lista estaitica

    
    // var lista []int = []int{1, 2, 3, 4, 5, 6, 7}
    // var nomes []string = []string{"uva", "ovo", "eva"}

    // fmt.Println(nomes[2]) // uva

    // nomes[2] = macarrao // alterar
    // fmt.Println(lista)
    // fmt.Println(nomes)



// var qtd int
// fmt.Scan(&qtd)
// var idades [] int = make([]int, qtd)
// fmt.Println(idades) 



arr := []int{1, 2, 3, 4, 5}
fmt.Print("[ ")
for _, valor := range arr {
    fmt.Printf("%d ", valor)
    }
fmt.Print("]\n")
}