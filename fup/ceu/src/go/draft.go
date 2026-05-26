package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    fmt.Print("[ ")

    for i := 0; i < n; i++ {
        fmt.Print(i, " ")
    }
    for i:= n+1; i <= 9; i++ {
        fmt.Print(i, " ")
    }
            if n != 10 {
        fmt.Print("ceu ")
    }
        fmt.Println("]")
} 
