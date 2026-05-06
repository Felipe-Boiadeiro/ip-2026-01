package main

import "fmt"

func somaArray(arr []float64) float64 {
	soma := .0
	for _, valor := range arr {
		soma += valor
	}
	return soma
}

func main() {
	numeros := []float64{1.5, 2.5, 3.0}
	resultado := somaArray(numeros)
	fmt.Println(resultado)
}
