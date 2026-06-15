package main

import "fmt"

func main() {
	var a, b, c, d int

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&d)

	soma := a + b + c + d

	if soma == 0 {
		fmt.Println("nenhum")
		return
	}

	vencedor := soma % 4

	switch vencedor {
	case 1:
		fmt.Println("jog1")
	case 2:
		fmt.Println("jog2")
	case 3:
		fmt.Println("jog3")
	case 0:
		fmt.Println("jog4")
	}
}