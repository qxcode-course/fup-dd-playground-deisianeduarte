package main
import "fmt"
func main() {
    var t, h, m, s, r int
    fmt.Scan(&t)
    h = t / 3600
    r = t % 3600
    m = r / 60
    s = r % 60

    fmt.Printf("%d:%d:%d\n", h, m, s)

}

// tempo em segundos = t
// horas = h
// resto = r (o que sobra em segundos depois de tirar as hs)
// minutos = m
// segundos = s

// vai entrar um numero em segundos, ai tem que tirar a hora e vai sobrar um valor (r).
// Agora tem que tirar os minutos, vai sobrar de novo, 
// agora o r vai ser os segundos e no final o resultado tem que ser h:m:s