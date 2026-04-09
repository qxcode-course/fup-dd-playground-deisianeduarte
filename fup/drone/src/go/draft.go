package main

import "fmt"

func main() {
    var A, B, C, H, L int
    fmt.Scan(&A)
    fmt.Scan(&B)
    fmt.Scan(&C)
    fmt.Scan(&H)
    fmt.Scan(&L)

    pass := false

    if (A <= H && B <= L) || (A <= L && B <= H) {
        pass = true
    }
    if (A <= H && C <= L) || (A <= L && C <= H) {
        pass = true
    }
    if (B <= H && C <= L) || (B <= L && C <= H) {
        pass = true
    }

    if pass {
        fmt.Println("S")
    } else {
        fmt.Println("N")
    }
}

