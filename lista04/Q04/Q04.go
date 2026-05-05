package main

import "fmt"

func main() {
	sequencia := []int{}
	frequencia := map[int]int{}

	for i := 1; i <= 10; i++ {
		n := 0
		fmt.Printf("Digite o %d° número: ", i)
		fmt.Scan(&n)
		sequencia = append(sequencia, n)
		frequencia[n]++
	}

	for valor, qtd := range frequencia {
		if qtd > 1 {
			fmt.Printf("O número %d aparece %d vezes\n", valor, qtd)
		}

	}
}
