package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    if n == 0 {
        fmt.Println()
    }
    arr := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }

    for _, valor := range arr {
        fmt.Println( valor)
    }

}
