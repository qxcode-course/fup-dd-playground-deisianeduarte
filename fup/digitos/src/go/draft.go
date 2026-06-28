package main
import (
        "fmt"
        "io"
        "strings"
)

func main () {
    var a string
    var b string

    for {
        _, err := fmt.Scan(&a, &b)
        if err == io.EOF {
            break
        }
        if err != nil {
            continue
        }

        quantidade := strings.Count(b, a)

        fmt.Println(quantidade)
       
    }
}
