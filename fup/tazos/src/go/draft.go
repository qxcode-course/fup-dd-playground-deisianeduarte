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
        vetor := make([]int, n)
        for i := 0; i < n; i ++ {
            scanner.Scan()
            vetor[i], _= strconv.Atoi(scanner.Text())
        }
        var maisRepetidos []string
        maxFrequencia := 0
        contaAtual := 1

        for i := 1; i <= n; i++ {
            if i == n || vetor[i] != vetor[i-1] {
                if contaAtual > maxFrequencia {
                    maxFrequencia = contaAtual
                    maisRepetidos = []string{strconv.Itoa(vetor[i-1])}
                } else if contaAtual == maxFrequencia{
                    maisRepetidos = append(maisRepetidos,strconv.Itoa(vetor[i-1]))
                }
                contaAtual = 1
            } else {
                contaAtual++
            }
        }
        if len(maisRepetidos) == 0 {
            fmt.Println("[ ]")
        } else {
            fmt.Printf("[ %s ]\n", strings.Join(maisRepetidos, " "))
        }
    }
}
