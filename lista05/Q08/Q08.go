package main

import (
	"fmt"
	"math"
)

func main() {
	sequencia := []float64{}
	for i := 1; i <= 15; i++ {
		n := .0
		fmt.Printf("informe o %d° valor: ", i)
		fmt.Scan((&n))

		if n < 0 {
			sequencia = append(sequencia, -1)

		} else {
			sequencia = append(sequencia, math.Pow(n, 0.5))
		}
	}
	fmt.Println(sequencia)
}
