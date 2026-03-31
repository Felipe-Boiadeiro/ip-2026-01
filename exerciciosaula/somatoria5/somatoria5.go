package main

import "fmt"

func main() {
	var valor [5]float64
	var soma float64
	for i := 0; i < 5; i++ {

		fmt.Printf("Informe o valor %d: ", i+1)
		fmt.Scan(&valor[i])

		soma += valor[i]
	}

	fmt.Printf("O valor da soma é %.2f\n", soma)
}
