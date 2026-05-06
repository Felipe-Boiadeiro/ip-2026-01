package main

import "fmt"

func main() {

	sequencia := []int{}
	posicao := 0
	for i := 1; i <= 10; i++ {
		n := 0
		fmt.Printf("Informe o %d° número: ", i)
		fmt.Scan(&n)
		sequencia = append(sequencia, n)
	}

	menor := sequencia[0]
	for i := range sequencia {
		if sequencia[i] < menor {
			menor = sequencia[i]
			posicao = i

		}

	}
	fmt.Printf("O menor elemento do vetor é %d e sua posição dentro do vetor é %d", menor, posicao)
}
