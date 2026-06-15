package main

import "fmt"

func floor(num float64) int {
	return int(num)
}

func ceil(num float64) int {
	if num == float64(int(num)) {
		return int(num)
	}
	return int(num) + 1
}

func round(num float64) int {
	parteInteira := int(num)
	parteFracionaria := num - float64(parteInteira)

	if parteFracionaria >= 0.5 {
		return parteInteira + 1
	}
	return parteInteira
}

func main() {
	var op string
	var num float64

	fmt.Scan(&op)
	fmt.Scan(&num)

	switch op {
	case "c":
		fmt.Println(ceil(num))
	case "f":
		fmt.Println(floor(num))
	case "r":
		fmt.Println(round(num))
	}
}