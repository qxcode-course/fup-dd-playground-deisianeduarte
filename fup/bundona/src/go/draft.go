package main

import "fmt"

func main() {
	var h, m, d int
	var s string

	fmt.Scan(&h)
	fmt.Scan(&m)
	fmt.Scan(&s)
	fmt.Scan(&d)

	// Cada cm equivale a 10 minutos
	posicao := h*60 + m

	if s == "H" {
		posicao += d * 10
	} else {
		posicao -= d * 10
	}

	// Relógio de 12 horas = 720 minutos
	posicao = ((posicao % 720) + 720) % 720

	hora := posicao / 60
	minuto := posicao % 60

	fmt.Printf("%02d %02d\n", hora, minuto)
}