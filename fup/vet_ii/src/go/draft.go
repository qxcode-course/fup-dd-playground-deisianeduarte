package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    if n == 0 {
        fmt.Println()
    }
    arr := make([]int, n)
    fmt.Print("[")
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }

    for _, valor := range arr {
        fmt.Println( valor)
    }
    fmt.Print("]\n")
}
