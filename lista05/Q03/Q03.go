package main

import "fmt"

func main() {
	vetor := []int{}
	soma := 0
	pares := []int{}
	impares := []int{}
	contador := 0

	for i := 1; i <= 10; i++ {
		n := 0
		fmt.Printf("Informe o vetor %d: ", i)
		fmt.Scan(&n)

		vetor = append(vetor, n)
	}
	for i := range vetor {
		if vetor[i]%2 == 0 {
			pares = append(pares, vetor[i])
			soma += vetor[i]

		} else {
			impares = append(impares, vetor[i])
			contador++

		}

	}
	fmt.Println(pares)
	fmt.Println(soma)
	fmt.Println(impares)
	fmt.Println(contador)
}
