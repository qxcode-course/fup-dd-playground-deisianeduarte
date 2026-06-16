package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	texto, _ := reader.ReadString('\n')

	// Remove o '\n' se existir
	if len(texto) > 0 && texto[len(texto)-1] == '\n' {
		texto = texto[:len(texto)-1]
	}

	runes := []rune(texto)

	for i := len(runes) - 1; i >= 0; i-- {
		fmt.Print(string(runes[i]))
	}
	fmt.Println()
}