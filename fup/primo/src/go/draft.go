package main

import "fmt"

func ehPrimo(n int) int {
    if n < 2 {
        return 0
    }

    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return 0
        }
    }
    
    return 1
}

package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    fmt.Println(ehPrimo(n))
}




















func main() {
    
}