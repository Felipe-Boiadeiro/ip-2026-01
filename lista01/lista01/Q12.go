package main

import (
	"fmt"
)

func main() {

	var temp int

	fmt.Println("Digite a quantidade de horas utilizadas: ")

	fmt.Scan(&temp)

	var valor float64

	par := temp / 3

	resto := temp % 3

	valor = float64(par * 10)

	valor += float64(resto * 5)

	fmt.Printf("O VALOR A PAGAR E = %.2f\n", valor)

}
