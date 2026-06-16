package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	chocolate, limao := 0, 0
	manha, tarde := 0, 0

	for i := 0; i < n; i++ {
		var sabor, turno string
		fmt.Scan(&sabor, &turno)

		if sabor == "c" {
			chocolate++
		} else if sabor == "l" {
			limao++
		}

		if turno == "m" {
			manha++
		} else if turno == "t" {
			tarde++
		}
	}

	// Sabor mais vendido
	if chocolate > limao {
		fmt.Println("c")
	} else if limao > chocolate {
		fmt.Println("l")
	} else {
		fmt.Println("empate")
	}

	// Turno mais vago (menos vendas)
	if manha < tarde {
		fmt.Println("m")
	} else if tarde < manha {
		fmt.Println("t")
	} else {
		fmt.Println("empate")
	}
}