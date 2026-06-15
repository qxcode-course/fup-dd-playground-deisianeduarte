package main

import "fmt"

func main() {
	var b, t float64

	fmt.Scan(&b)
	fmt.Scan(&t)

	// Área da parte esquerda (trapézio)
	areaEsquerda := ((b + t) * 70) / 2

	// Metade da área total da nota
	metade := (160 * 70) / 2.0

	if areaEsquerda > metade {
		fmt.Println(1)
	} else if areaEsquerda < metade {
		fmt.Println(2)
	} else {
		fmt.Println(0)
	}
}
