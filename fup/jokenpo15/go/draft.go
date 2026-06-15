package main

import "fmt"

func main() {
	var a, b int

	fmt.Scan(&a)
	fmt.Scan(&b)

	if a == b {
		fmt.Println("Empate")
		return
	}

	// Distância de A para B no círculo de 15 elementos
	dist := (b - a + 15) % 15

	if dist >= 1 && dist <= 7 {
		fmt.Println("Jogador 1")
	} else {
		fmt.Println("Jogador 2")
	}
}


//     import "fmt"
//     func main() {    
//     x := 0
//     for { //lop infinito 
//     x += 1
//     if x > 10 {
//     break //sai fora
//     }
//     fmt.Println(x)
// }