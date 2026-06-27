package main

import "fmt"

func ConsegueSair(p, e, s int) bool {
    pos := 0 
    salto := s 

    for {
        if pos+salto >= p {
            return true
        }

        pos = pos + salto - e
        if pos < 0 {
            return false
        }
        salto -= 10
    }
}
func main() {
    var p, e int
    fmt.Scan(&p)
    fmt.Scan(&e)
    
    s := 0
    for {
        if ConsegueSair(p, e, s) {
            fmt.Println(s)
            break
        }
        s++
    }
}
