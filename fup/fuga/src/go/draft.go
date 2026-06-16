package main

import "fmt"

func main() {
	var h, p, f, d int

	fmt.Scan(&h, &p, &f, &d)

	pos := f

	for {
		if pos == h {
			fmt.Println("S")
			return
		}

		if pos == p {
			fmt.Println("N")
			return
		}

		pos = (pos + d + 16) % 16
	}
}