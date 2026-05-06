package main

import "fmt"

func main() {
	vetor := []int{}
	vetorInverso := []int{}
	n := 0

	for i := 1; i <= 10; i++ {
		fmt.Printf("Informe o %dº elemento do vetor:\n", i)
		fmt.Scan(&n)
		vetor = append(vetor, n)
	}

	codigo := 0
	fmt.Print("Informe o código: ")
	fmt.Scan(&codigo)

	switch codigo {
	case 0:
		return

	case 1:
		fmt.Println(vetor)

	case 2:
		for i := 9; i >= 0; i-- {
			vetorInverso = append(vetorInverso, vetor[i])
		}
		fmt.Println(vetorInverso)

	default:
		fmt.Println("Código inválido")
	}
}
