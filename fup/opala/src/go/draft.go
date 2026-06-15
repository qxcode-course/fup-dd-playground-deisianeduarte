
package main

import "fmt"

func main() {
	var velocidade, tempoMinutos, consumo float64

	fmt.Scan(&velocidade)
	fmt.Scan(&tempoMinutos)
	fmt.Scan(&consumo)

	tempoHoras := tempoMinutos / 60.0
	distancia := velocidade * tempoHoras
	desempenho := distancia / consumo

	fmt.Printf("%.2f\n", desempenho)
}
