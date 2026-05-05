package main

import (
	"fmt"
	"math"
)

func main() {
	sequencia := []float64{}
	valor := .0
	n := .0
	for i := 1; i <= 100; i++ {
		fmt.Scan(&n)
		sequencia = append(sequencia, n)
	}
	for i := range sequencia {
		valor += (math.Pow((sequencia[i] - sequencia[99-i]), 3))
	}
	fmt.Println(valor)

}
