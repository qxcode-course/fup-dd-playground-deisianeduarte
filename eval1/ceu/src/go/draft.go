package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    fmt.Print("[ ")
    first := true

    for i := 0; i <= 10; i++ {
        if i == n {
            continue
        }
        if !first {
            fmt.Print(" ")
        }
        first = false

        if i == 10 {
            fmt.Print("ceu")
        } else {
            fmt.Print(i)
        }
    }

    fmt.Println(" ]")
}

