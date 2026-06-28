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
        totalAlbum, _ := strconv.Atoi(scanner.Text())

        scanner.Scan()
        possuiQtd, _ := strconv.Atoi(scanner.Text())

        frequencia :=  make(map[int]int)
        var repetidas []string

        for i := 0; i< possuiQtd; i++ {
            scanner.Scan()
            fig, _ := strconv.Atoi(scanner.Text())
            frequencia[fig]++

            if frequencia[fig] > 1 {
                repetidas = append(repetidas, strconv.Itoa(fig))
            }
        }
        var faltando []string
        for i := 1; i <= totalAlbum; i++ {
            if frequencia[i] == 0 {
                faltando = append(faltando, strconv.Itoa(i))
            }
        }

        imprimirResultado(repetidas)
        imprimirResultado(faltando)

    }
}

func imprimirResultado(lista []string) {
    if len(lista) == 0 {
        fmt.Println("[ ]")
    } else {
        fmt.Printf("[ %s ]\n", strings.Join(lista, " "))
    }
}