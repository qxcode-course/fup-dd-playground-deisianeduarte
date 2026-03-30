package main
import "fmt"
func main() {
    var p, d1, d2 int64
    fmt.Scan(&p)
    fmt.Scan(&d1)
    fmt.Scan(d2)
    soma := d1 + d2
    // se a soma for par
    if soma%2 == 0 {
    if p == 0 {
        fmt.Println(0)
        }//a Alice ganha
    } else{
        //se a soma for ímpar
        fmt.Println(1)
        //o Bob ganha 

    } else {
        if p == 0 {
            fmt.Println(1)
        } else {
            fmt.Println(0)
        }
    }
    
    
}
