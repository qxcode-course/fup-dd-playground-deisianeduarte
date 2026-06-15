package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	delta := math.Pow(b, 2) - 4*a*c

	if delta < 0 {
		fmt.Println("nao ha raiz real")
	} else if delta == 0 {
		x := -b / (2 * a)
		if x == 0 {
			x = 0
		}
		fmt.Printf("%.2f\n", x)
	} else {
		x1 := (-b + math.Sqrt(delta)) / (2 * a)
		x2 := (-b - math.Sqrt(delta)) / (2 * a)
		if x1 == 0 {
			x1 = 0
		}
		if x2 == 0 {
			x2 = 0
		}

		fmt.Printf("%.2f\n", x1)
		fmt.Printf("%.2f\n", x2)
	}
}
