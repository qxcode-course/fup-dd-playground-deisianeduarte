package main

import "fmt"

func main() {
	var num, menor int

	fmt.Scan(&menor)

	for i := 0; i < 4; i++ {
		fmt.Scan(&num)

		if num < menor {
			menor = num
		}
	}

	fmt.Println(menor)
}