package main

import "fmt"

func main() {
	preco, quantidade := 6., 130
	lucroMax, precoLMax, quantidadeMax := .0, .0, 0

	for i := 0; preco >= 1; i++ {
		fmt.Printf("Lucro esperado(R$ %.2v): R$ %.2f\n", preco, preco*float64(quantidade)-300)

		if preco*float64(quantidade)-300 > float64(lucroMax) {
			lucroMax = preco*float64(quantidade) - 300
			precoLMax = preco
			quantidadeMax = quantidade
		}

		preco -= .6
		quantidade += 30
	}

	fmt.Printf("Lucro máximo: %.2f\n", lucroMax)
	fmt.Printf("Preço para ter lucro máximo: %.2f\n", precoLMax)
	fmt.Println("Quantidade de ingressos para lucro máximo:", quantidadeMax)

}
