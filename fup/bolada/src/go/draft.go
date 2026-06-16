package main

import "fmt"

func main() {
	var tipo string
	var forca int

	fmt.Scan(&tipo)
	fmt.Scan(&forca)

	var t int

	if tipo == "b" {
		t = 20
	} else { // tipo == "c"
		t = 18
	}

	poder := ((forca * t) - 80) / 10

	if poder < 150 {
		fmt.Println("Fraco, nem passou")
	} else if poder < 180 {
		fmt.Println("Perfeito")
	} else if poder <= 210 {
		fmt.Println("Satisfeito")
	} else {
		fmt.Println("Muito forte, bola fora")
	}
}