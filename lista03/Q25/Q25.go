package main

import "fmt"

func main() {
	soma, valor := .0, 1

	for i := 0; i <= 14; i++ {

		sinal := 1.0
		if i%2 != 0 {
			sinal = -1.0
		}

		divisor := (15 - i) * (15 - i)

		soma += sinal * float64(valor) / float64(divisor)

		valor *= 2
	}

	fmt.Println(soma)
}
