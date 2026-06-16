package main
import "fmt"

func main() {
	var p1, p2, p3, trabalho float64

	fmt.Scan(&p1)
	fmt.Scan(&p2)
	fmt.Scan(&p3)
	fmt.Scan(&trabalho)

	menor := p1
	if p2 < menor {
		menor = p2
	}
	if p3 < menor {
		menor = p3
	}

	media := (p1 + p2 + p3 - menor + trabalho) / 3.0

	if media >= 7.0 {
		fmt.Printf("Aprovado com %.1f\n", media)
	} else {
		fmt.Printf("Final com %.1f\n", media)
	}
}