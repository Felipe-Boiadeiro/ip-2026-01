package main

import "fmt"

func main() {
	dados := map[int]int{}
	numeros := []int{}

	for i := 0; i < 20; i++ {
		var valor int
		fmt.Printf("Digite o valor do dado %d: ", i+1)
		fmt.Scan(&valor)

		numeros = append(numeros, valor)
		dados[valor]++
	}

	fmt.Println("\nNúmeros sorteados:")
	for i := 0; i < len(numeros); i++ {
		fmt.Print(numeros[i], " ")
	}

	fmt.Println("\n\nFrequência:")
	for i := 1; i <= 6; i++ {
		fmt.Printf("Número %d apareceu %d vezes\n", i, dados[i])
	}
}
