package main

import "fmt"

func main() {
	vetor1 := []int{}
	vetor2 := []int{}
	resultante1 := []int{}
	resultante2 := []int{}

	for i := 1; i <= 10; i++ {
		n := 0
		fmt.Print("Digite um número para o vetor 1:")
		fmt.Scan(&n)
		vetor1 = append(vetor1, n)
	}
	for i := 1; i <= 5; i++ {
		n := 0
		fmt.Print("Digite um número para o vetor 2:")
		fmt.Scan(&n)
		vetor2 = append(vetor2, n)
	}
	somaV2 := 0
	for i := range vetor2 {
		somaV2 += vetor2[i]
	}
	for i := range vetor1 {
		if vetor1[i]%2 == 0 {
			resultante1 = append(resultante1, vetor1[i]+somaV2)
		} else {
			resultante2 = append(resultante2, vetor1[i]+somaV2)
		}
	}
	fmt.Println(resultante1)
	fmt.Println(resultante2)
}
