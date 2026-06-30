package main
import ("fmt"
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
        var alunos []string
        var servidores []string

        for i := 0; i < n; i++ {
            scanner.Scan()
            numStr := scanner.Text()
            num, _ := strconv.Atoi(numStr)

            if num%2 == 0 {
                servidores = append(servidores, numStr)
            } else {
                alunos = append(alunos, numStr)
            }
        }
        imprimirLista(alunos)
        imprimirLista(servidores)
    }
}
func imprimirLista(lista []string) {
    if len(lista) == 0 {
        fmt.Println("[ ]")
    } else {
        fmt.Printf("[ %s ]\n", strings.Join(lista, " "))
    }
}