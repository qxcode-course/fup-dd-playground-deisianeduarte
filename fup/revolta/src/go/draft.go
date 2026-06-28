package main
import (
        "fmt"
        "io"
)
   
func main(){
    var n int 

    for {
        _, err := fmt.Scan(&n)
        if err == io.EOF {
            break
        }
        if err != nil {
            continue
        }

        somaSoldados := 0
        somaRebeldes := 0

        for i := 0; i < n; i++ {
            var x int
            fmt.Scan(&x)

            if x%2 == 0 {
                somaRebeldes += x
            } else {
                somaSoldados += x
            }
        }
    

    if somaSoldados > somaRebeldes {
        fmt.Println("soldados")
    } else if somaRebeldes > somaSoldados {
        fmt.Println("rebeldes")
    } else {
        fmt.Println("empate")
    }
}
} 