package main

import "fmt"

func main() {
	vetor1 := []int{}
	vetor2 := []int{}
	vetorFinal := []int{}

	for i := 1; i <= 10; i++ {
		n := 0
		fmt.Printf("Informe o %dº número do vetor 1: ", i)
		fmt.Scan(&n)
		vetor1 = append(vetor1, n)
	}
	for i := 1; i <= 10; i++ {
		n := 0
		fmt.Printf("Informe o %dº número do vetor 2: ", i)
		fmt.Scan(&n)
		vetor2 = append(vetor2, n)
	}
	for i := 0; i <= 9; i++ {
		vetorFinal = append(vetorFinal, vetor1[i])
		vetorFinal = append(vetorFinal, vetor2[i])
	}
	fmt.Println(vetorFinal)
}
