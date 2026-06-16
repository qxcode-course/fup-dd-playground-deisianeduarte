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
		if pos+salto >= P {
			fmt.Printf("%d saiu\n", pos)
			break
		}

		fmt.Printf("%d %d\n", pos, pos+salto)

		pos = pos + salto - E

		salto -= 10

		if pos < 0 {
			fmt.Printf("%d morreu\n", pos)
			break
		}
	}
}