package main
import (
	"fmt"
	"math"
)
func main() {
	var n, raiz, quadrado float64
	fmt.Println("Digite um número: ")
	fmt.Scan(&n)

	if n >= 0 {
		raiz = math.Pow(n, 0.5)
		fmt.Printf("A raíz do número é %f\n", raiz)
	} else {
		quadrado = math.Pow(n, 2)
		fmt.Printf("O quadrado do número é %f\n", quadrado)
	}
}
