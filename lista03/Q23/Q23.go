package main

import "fmt"

func main() {
	n, resultado := 0, 0.0
	sequencia := []int{1000}

	fmt.Scan(&n)

	if n < 0 {
		return
	}

	for i := 1; i <= n-1; i++ {
		soma := sequencia[i-1] - 3
		sequencia = append(sequencia, soma)
	}

	for i := 0; i <= n-1; i++ {
		termo := float64(sequencia[i]) / float64(i+1)
		if i%2 != 0 {
			termo = -termo
		}
		resultado += termo
	}

	fmt.Println(resultado)
}
