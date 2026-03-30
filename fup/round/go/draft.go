package main

import "fmt"

func floor(x float64) int {
	i := int(x)
	if x < 0 && x != float64(i) {
		return i - 1
	}
	return i
}

func ceil(x float64) int {
	i := int(x)
	if x > 0 && x != float64(i) {
		return i + 1
	}
	return i
}

func round(x float64) int {
	i := int(x)
	if x >= 0 {
		if x-float64(i) >= 0.5 {
			return i + 1
		}
		return i
	}
	if float64(i)-x > 0.5 {
		return i - 1
	}
	return i
}

func main() {
	var op byte
	var x float64

	fmt.Scanf("%c %f", &op, &x)

	if op == 'f' {
		fmt.Println(floor(x))
	} else if op == 'c' {
		fmt.Println(ceil(x))
	} else {
		fmt.Println(round(x))
	}
}


// x = numero quebrado
// i = numero inteiro 
