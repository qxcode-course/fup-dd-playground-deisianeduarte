package main
import "fmt"
func main() {
    var n int
    var pe rune = 'd'
    fmt.Scan(&n)
    fmt.Scan(&pe)

    fmt.Print("[ ")

    if pe == 'd' {
    
        if n != 0 {
            fmt.Print("0d ")
    for i := 1; i < n; i++ {
        fmt.Print(i)

        if i % 2 == 0 {
            fmt.Print("d ")

        } else {
            fmt.Print("e ")
        }
        }
    }
    
    for i:= n+1; i <= 9; i++ {
        fmt.Print(i)

        if i % 2 == 0 {
            fmt.Print("d ")

        } else {
            fmt.Print("e ")
        }
    }

            if n != 10 {
        fmt.Print("ceu ")
            }
        
    }
        if pe == 'e' {
    
        if n != 0 {
            fmt.Print("0e ")
    for i := 1; i < n; i++ {
        fmt.Print(i)

        if i % 2 == 0 {
            fmt.Print("e ")

        } else {
            fmt.Print("d ")
        }
        }
    }
    
    for i:= n+1; i <= 9; i++ {
        fmt.Print(i)
        
        if i % 2 == 0 {
            fmt.Print("e ")

        } else {
            fmt.Print("d ")
        }
    }

            if n != 10 {
        fmt.Print("ceu ")
            }
        
    }
        fmt.Println("]")
} 

