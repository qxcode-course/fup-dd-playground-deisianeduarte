package main

import "fmt"

func main() {
	var l, c int

	fmt.Scan(&l)
	fmt.Scan(&c)

	if (l+c)%2 == 0 {
		fmt.Println(1) // branca
	} else {
		fmt.Println(0) // preta
	}
}