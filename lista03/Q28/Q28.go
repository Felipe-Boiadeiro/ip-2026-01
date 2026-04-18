package main

import (
	"fmt"
	"math"
)

func main() {

	soma, pi := .0, .0

	for i := 1; i <= 52; i++ {

		contador := float64(2*i) - 1

		sinal := math.Pow(-1, float64(i+1))

		soma += sinal / math.Pow(contador, 3)

	}

	pi = math.Pow(soma*32, 1./3.)
	fmt.Println(pi)

}
