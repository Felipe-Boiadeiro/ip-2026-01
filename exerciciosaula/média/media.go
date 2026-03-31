package main

import "fmt"

func main() {
	var nota [3]float64
	var soma float64
	for i := 0; i < 3; i++ {

		fmt.Printf("Informe a nota %d: ", i+1)
		fmt.Scan(&nota[i])

		soma += nota[i]
	}

	media := soma / 3

	fmt.Printf("A média é %.2f\n", media)

	for i := range nota {
		if nota[i] > media {
			fmt.Printf("A %v° nota está acima da média\n", i+1)

		}
	}
}
