package main

import (
        "fmt"
        "bufio"
        "os"
        "strconv"
)

func fibonacci(n int) uint64 {
    if n == 0 {
        return 0 
    }
    if n == 1 || n == 2 {
        return 1 
    }

    var a, b uint64 = 1, 1
    for i := 3; i <= n; i++ {
            a, b = b, a+b
    }
    return b

}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {    
            linha := scanner.Text()
            if linha == "" {
                continue
        }

            n, _ := strconv.Atoi(linha)
            fmt.Println(fibonacci(n))
    }
    
}