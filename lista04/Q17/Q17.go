package main

import "fmt"

func main() {
	sequencia := []int{}

	for i := 1; i <= 10; i++ {
		n := 0
		fmt.Printf("Informe o %dº número: ", i)
		fmt.Scan(&n)
		sequencia = append(sequencia, n)
	}
	for i := range sequencia {
		divisores := 0

		for y := 1; y <= sequencia[i]; y++ {
			if sequencia[i]%y == 0 {
				divisores++
			}
			if divisores > 2 {
				break
			}
		}

		if divisores == 2 {
			fmt.Printf("%d é primo e se encontra na posição %d do vetor\n", sequencia[i], i)

		}
	}
}
