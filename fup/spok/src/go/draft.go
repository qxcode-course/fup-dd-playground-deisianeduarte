package main

import "fmt"

func inverter(n int) int {
	invertido := 0

	for n > 0 {
		digito := n % 10
		invertido = invertido*10 + digito
		n /= 10
	}

	return invertido
}

func main() {
	var n int

	fmt.Scan(&n)

	if n == inverter(n) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}