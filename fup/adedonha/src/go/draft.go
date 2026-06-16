package main

import "fmt"

func main() {
	var soma int

	fmt.Scan(&soma)

	if soma == 0 {
		fmt.Println("joguem de novo")
		return
	}

	letra := 'a' + rune((soma-1)%26)

	fmt.Printf("%c\n", letra)
}