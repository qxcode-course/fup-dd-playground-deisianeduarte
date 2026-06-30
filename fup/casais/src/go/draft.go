package main
import (
        "fmt"
        "bufio"
        "os"
        "strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)

    for scanner.Scan() {
        n, _ := strconv.Atoi(scanner.Text())
        esperando := make(map[int]int)
        casais := 0

        for i := 0; i < n; i++ {
            scanner.Scan()
            animal, _ := strconv.Atoi(scanner.Text())
            parOposto := animal * -1

            if esperando[parOposto] > 0 {
                casais ++
                esperando[parOposto]--
            } else {
                esperando[animal]++
            }
        }
        fmt.Println(casais)
    }
}