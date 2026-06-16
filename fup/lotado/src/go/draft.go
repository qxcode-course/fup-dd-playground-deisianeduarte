/*package main
import "fmt"
func main() {
    var c, m int
    fmt.Scan(&c)
    q := c * 2

    for i := 0; i <= q; i++ {
        if i <= 0 {
            fmt.Println("vazio")
        }
        if i > 0 && i < c {
                fmt.Println("ainda cabe")
        }
        if i >= c && i < q {
                fmt.Println("lotado")
        }
        if i >= q {
                fmt.Println("hora de partir")
        }

        fmt.Scan(&m)
            i = i + m 
        if i < 0 {
            i = 0
        }
        }
    

}*/

package main

import "fmt"

func main() {
	var c, m int
	passageiros := 0

	fmt.Scan(&c)

	for {
		if _, err := fmt.Scan(&m); err != nil {
			break
		}

		passageiros += m

		if passageiros >= 2*c {
			fmt.Println("hora de partir")
			break
		} else if passageiros == 0 {
			fmt.Println("vazio")
		} else if passageiros < c {
			fmt.Println("ainda cabe")
		} else {
			fmt.Println("lotado")
		}
	}
}