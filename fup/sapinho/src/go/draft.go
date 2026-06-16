package main

import "fmt"

func main() {
	var P, S, E int

	fmt.Scan(&P)
	fmt.Scan(&S)
	fmt.Scan(&E)

	pos := 0

	for {
		if pos+S >= P {
			fmt.Printf("%d saiu\n", pos)
			break
		}

		fmt.Printf("%d %d\n", pos, pos+S)
		pos = pos + S - E
	}
}