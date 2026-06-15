package main

import "fmt"

func main() {
	var A, G, Ra, Rg float64

	fmt.Scan(&A)
	fmt.Scan(&G)
	fmt.Scan(&Ra)
	fmt.Scan(&Rg)

	custoAlcool := A / Ra
	custoGasolina := G / Rg

	if custoAlcool < custoGasolina {
		fmt.Println("A")
	} else {
		fmt.Println("G")
	}
}