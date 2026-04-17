package main

import "fmt"

func main() {
	n := 0
	intervaloImpar := []int{}

	somatorio := []int{1}

	fmt.Scan(&n)

	for i := 1; i <= n; i++ {

		valor := 2*i + 1
		intervaloImpar = append(intervaloImpar, valor)

	}

	for i := 1; i <= n; i++ {

		novo := somatorio[i-1] + intervaloImpar[i-1]
		somatorio = append(somatorio, novo)

	}

	fmt.Println(somatorio)

}
