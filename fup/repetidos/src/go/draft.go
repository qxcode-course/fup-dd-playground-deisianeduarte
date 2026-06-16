package main

import "fmt"

func main() {
	var p, n int
	fmt.Scan(&p, &n)

	contador := 0

	for i := 0; i < n; i++ {
		var valor int
		fmt.Scan(&valor)

		if valor == p {
			contador++
		}
	}

	fmt.Println(contador)
}