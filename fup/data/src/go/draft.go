package main

import "fmt"

func main() {
	var h, m, d, mo, y int
	if _, err := fmt.Scan(&h); err != nil {
		return
	}
	if _, err := fmt.Scan(&m); err != nil {
		return
	}
	if _, err := fmt.Scan(&d); err != nil {
		return
	}
	if _, err := fmt.Scan(&mo); err != nil {
		return
	}
	if _, err := fmt.Scan(&y); err != nil {
		return
	}
	fmt.Printf("%02d:%02d %02d/%02d/%02d\n", h, m, d, mo, y%100)
}
