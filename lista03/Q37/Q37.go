package main

import (
	"fmt"
	"math"
)

func main() {
	sequencia := []int{}
	n, resultado := 0, 0
	fmt.Scan(&n)

	for n > 0 {
		digito := n % 10
		sequencia = append(sequencia, digito)
		n = n / 10
	}

	for i := range sequencia {
		resultado += int(math.Pow(float64(8), float64(i))) * sequencia[i]
	}

	fmt.Println(resultado)
}
