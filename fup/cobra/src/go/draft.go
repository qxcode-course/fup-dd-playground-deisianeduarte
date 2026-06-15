package main

import "fmt"

func main() {
	var n, x, y, s int
	var c string

	fmt.Scan(&n)
	fmt.Scan(&x)
	fmt.Scan(&y)
	fmt.Scan(&c)
	fmt.Scan(&s)

	switch c {
	case "R":
		x = (x + s) % n
	case "L":
		x = (x - s%n + n) % n
	case "D":
		y = (y + s) % n
	case "U":
		y = (y - s%n + n) % n
	}

	fmt.Printf("%d %d\n", x, y)
}
