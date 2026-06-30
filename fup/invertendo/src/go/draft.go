package main
import (
        "fmt"
        "bufio"
        "os"
        "strconv"
        "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)

    for scanner.Scan() {
        n, _ := strconv.Atoi(scanner.Text())
        elementos := make([]string, n)
        for i := 0; i < n; i++ {
            scanner.Scan()
            elementos[i] = scanner.Text()
        }
        for i, j := 0, len(elementos)-1; i < j; i, j = i+1, j - 1 {
            elementos[i], elementos[j] = elementos[j], elementos[i]
        }
        if len(elementos) == 0 {
            fmt.Println("[ ]")
        } else {
            fmt.Printf("[ %s ]\n", strings.Join(elementos, " "))
        }
    }

}