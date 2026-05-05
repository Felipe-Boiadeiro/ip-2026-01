package main

import "fmt"

func main() {
	sequencia := []int{}
	contador := 0

	for i := 1; i <= 10; i++ {
		a := 0

		fmt.Scan(&a)

		sequencia = append(sequencia, a)
	}

	for i := range sequencia {
		if sequencia[i] >= 50 {
			fmt.Printf("O número na posição %d vale: %d\n", i+1, sequencia[i])
			contador++
		}
	}
	if contador == 0 {
		fmt.Println("Não há números maiores que 50")
	}
}
