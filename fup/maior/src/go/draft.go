package main

import "fmt"

func main() {
	var chute, valor float64
	var escolha string

	fmt.Scan(&chute)
	fmt.Scan(&escolha)
	fmt.Scan(&valor)

	if valor == chute {
		fmt.Println("primeiro")
	} else if escolha == "M" {
		if valor > chute {
			fmt.Println("segundo")
		} else {
			fmt.Println("primeiro")
		}
	} else { // escolha == "m"
		if valor < chute {
			fmt.Println("segundo")
		} else {
			fmt.Println("primeiro")
		}
	}
}