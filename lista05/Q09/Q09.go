package main

import "fmt"

func main() {
	alturas := []float64{}

	for i := 1; i <= 10; i++ {
		altura := .0
		fmt.Printf("Informe a altura do %d° jogador: ", i)
		fmt.Scan(&altura)

		alturas = append(alturas, float64(altura))
	}
	media := .0
	soma := .0
	for i := range alturas {
		soma += alturas[i]
	}
	media = soma / 10
	for i := range alturas {
		if alturas[i] > media {
			fmt.Println(alturas[i])
		}

	}
}
