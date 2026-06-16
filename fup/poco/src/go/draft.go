package main

import "fmt"

func main() {
	var P, S, E int

	fmt.Scan(&P)
	fmt.Scan(&S)
	fmt.Scan(&E)

	pos := 0
	salto := S

	for {
		// Saiu do poço
		if pos+salto >= P {
			fmt.Printf("%d saiu\n", pos)
			break
		}

		fmt.Printf("%d %d\n", pos, pos+salto)

		// Nova posição após escorregar
		pos = pos + salto - E

		// Próximo salto é 10 cm menor
		salto -= 10

		// Morreu afogado
		if pos < 0 {
			fmt.Printf("%d morreu\n", pos)
			break
		}
	}
}