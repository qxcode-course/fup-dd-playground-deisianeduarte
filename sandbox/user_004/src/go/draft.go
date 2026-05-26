package main
import "fmt"

//CÓDIGO 01

// type Pessoa struct {
//     nome  string
//     idade int
// }

// func main() {
//     var deisi Pessoa = Pessoa{
//         nome: "Deisi",
//         idade: 18,
//     }
//     fmt.Println(deisi)
//     deisi.nome = "Deisi"
//     deisi.idade += 1
//     fmt.Println(deisi.idade)
// }

//CÓDIGO 02

type Pos struct {
    x, y int
}

func ehPrimeiroQuadrante(pos Pos) bool {
    return pos.x > 0 && pos.y > 0
}

func filtrarPriQuad(lista []Pos) []Pos {
    result := make([]Pos, 0)
    for _, elem := range lista {
        if ehPrimeiroQuadrante(elem) {
            result = append(result, elem)
        }
    }
    return result
}

// func main() {
//     list := make([]Pos, 0)
//     list = append(list, Pos{1, 2})
//     list = append(list, Pos{-1, 2})
//     list = append(list, Pos{1, -2})
//     list = append(list, Pos{-1, -2})
//     fmt.Println(filtrarPriQuad(list))
// }

type Circulo struct {
    centro Pos
    raio   int
}

func main() {
    circulo := Circulo{}
    circulo.centro.x = 5

    circulo.centro.y = 8
    circulo.raio = 10

    fmt.Println(circulo)
}

