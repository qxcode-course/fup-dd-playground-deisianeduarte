package main
import "fmt"
func main() {
    var p, d1, d2 int64
    fmt.Scan(&p)
    fmt.Scan(&d1)
    fmt.Scan(&d2)
    soma := d1 + d2
    
    // se a soma for par

    if soma%2 == 0 && p == 0 {
        fmt.Println(0)
        }
        
        //a Alice ganha
      
    if soma%2 == 0 && p == 1 {
        fmt.Println(1)

    } 
    if soma%2 != 0 && p == 0 {
        fmt.Println(1)
        }
    if soma%2 != 0 && p == 1 {
        fmt.Println(0)
        }
}
