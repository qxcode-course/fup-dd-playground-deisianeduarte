package main

import "fmt"

func main() {
	var r1, r2, r3, r4 string
	acertos := 0

	fmt.Scan(&r1)
	fmt.Scan(&r2)
	fmt.Scan(&r3)
	fmt.Scan(&r4)

	if r1 == "d" {
		acertos++
	}
	if r2 == "a" {
		acertos++
	}
	if r3 == "c" {
		acertos++
	}
	if r4 == "d" {
		acertos++
	}

	switch acertos {
	case 0:
		fmt.Println("Nunca assistiu")
	case 1:
		fmt.Println("Ja ouviu falar")
	case 2:
		fmt.Println("Interessado no assunto")
	case 3:
		fmt.Println("Fa")
	case 4:
		fmt.Println("Super Fa")
	}
}