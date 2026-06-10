package main
import "fmt"

type Carta struct {
    numero int
    naipe int
}

func toString(carta Carta) string {
    var symbols = []string{"♠", "♥", "♦", "♣"}
    var numeros = []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
    return fmt.Sprintf("%v%v", numeros[carta.numero], symbols[carta.naipe])
}

func main() {
    for num := 1; num <= 13; num++ {
        for naipe := 0; naipe <= 3; naipe++ {
            fmt.Printf("%v\n", toString(Carta{numero: num, naipe: naipe}))
        }
    }
}