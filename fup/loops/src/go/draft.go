package main

import "fmt"

func main() {
	var angulo int

	fmt.Scan(&angulo)

	resultado := ((angulo % 360) + 360) % 360

	fmt.Println(resultado)
}