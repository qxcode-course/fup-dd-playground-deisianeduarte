package main

import (
	"fmt"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var chico, cebolinha, n int

	fmt.Scan(&chico)
	fmt.Scan(&cebolinha)
	fmt.Scan(&n)

	totalPatas := 0

	for i := 0; i < n; i++ {
		var animal string
		fmt.Scan(&animal)

		switch animal {
		case "v":
			totalPatas += 4
		case "c":
			totalPatas += 4
		case "g":
			totalPatas += 2
		}
	}

	fmt.Println(totalPatas)

	diffChico := abs(chico - totalPatas)
	diffCebolinha := abs(cebolinha - totalPatas)

	if diffChico < diffCebolinha {
		fmt.Println("Chico Bento")
	} else if diffCebolinha < diffChico {
		fmt.Println("Cebolinha")
	} else {
		fmt.Println("empate")
	}
}