package main

import "fmt"

func fatorial(x int) float64 {
	if x == 0 {
		return 1
	}
	return float64(x) * fatorial(x-1)
}

func main() {
	resultado := .0

	for i := 0; i <= 19; i++ {
		numerador := float64(100 - i)
		resultado += numerador / fatorial(i)
	}

	fmt.Println(resultado)
}
