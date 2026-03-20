package main

import "fmt"

var a, b, c, delta float64

func main() {

	fmt.Println("Informe o valor do coeficiente A: ")

	fmt.Scan(&a)

	fmt.Println("Informe o valor do coeficiente B: ")

	fmt.Scan(&b)

	fmt.Println("Informe o valor do coeficiente C: ")

	fmt.Scan(&c)

	delta = (b * b) - (4 * a * c)

	fmt.Printf("O VALOR DE DELTA E = %.2f", delta)

}
