package main

import "fmt"

func main() {

	n, soma, transformador := 0, 0, 0
	sequencia := []int{1000}

	fmt.Scan(&n)

	if n < 0 {
		return
	}

	for i := 1; i <= n-1; i++ {

		transformador = sequencia[i-1]

		if sequencia[i-1] < 0 {

			transformador = -transformador

		}

		soma = transformador - 3

		if i%2 != 0 {
			soma = -soma
		}

		sequencia = append(sequencia, soma)

	}
	fmt.Println(sequencia)

}
