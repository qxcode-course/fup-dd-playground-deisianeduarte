package main
import (
        "bufio"
        "fmt"
        "os"
        "strconv"
        "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        linhaN := strings.TrimSpace(scanner.Text())
        if linhaN == "" {
            continue
        }

        n, err := strconv.Atoi(linhaN)
        if err != nil {
            continue
        }

        if n == 0 {
            fmt.Println("[]")
            continue
        }
        if !scanner.Scan() {
            break
        }
        linhaCartas := strings.TrimSpace(scanner.Text())

        campos := strings.Fields(linhaCartas)

        cartasFormatadas := make([]string, n)

        for i := 0; i < n && i < len(campos); i++ {
            valor, _ := strconv.Atoi(campos[i])

            switch valor {
            case 1:
                cartasFormatadas[i] = "A"
            case 11:
                cartasFormatadas[i] = "J"
            case 12:
                cartasFormatadas[i] = "Q"
            case 13:
                cartasFormatadas[i] = "K"
            default:
                cartasFormatadas[i] = campos[i]       
            }
        }
        fmt.Printf("[%s]\n", strings.Join(cartasFormatadas, ", "))
    }
}