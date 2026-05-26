package main

import "fmt"

type Bota struct {
    pe string
    tam int
}

func contPares(botas []Bota) int {
    par := Bota{tam: bota.tam}
    if bota.pe == "D" {
        par.pe = "E"
    } else {
        par.pe = "D"
    }
    return par
}
    func main() {
    qtd := 0
    fmt.Scan(&qtd)
    botas := make([]Bota, qtd)
    for i := range botas {
        fmt.Scan(&botas[i].tam, &botas[i].pe)
    }
    fmt.Println(botas)
}

cont_pares := 0
for i, elem := range botas {
    par : = cont_Pares(elem)
    pos := procurarBota(botas, par)
    if pos != -1 {
        cont_pares += 1
        botas[i].pe = "X"
        botas[pos].pe = "X"
    }