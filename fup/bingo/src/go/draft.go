package main

import "fmt"

// func main() {
//     matriz := make([][]int, 0)
//     matriz = append(matriz, make([]int, 4))
//     matriz = append(matriz, make([]int, 4))
//     matriz = append(matriz, make([]int, 4))

//     for _, fila := range matriz {
//         fmt.Println(i, fila)
//     }
// }

func main() {
    var matriz [][]int {
        {1, 9, 27, 23},
        {34, 20, 37, 47},
        {30, 87, 55, 69},
        {13, 60, 99, 66},
    }
    for _, linha := range matriz {
        fmt.Println(linha)
    }
    }
