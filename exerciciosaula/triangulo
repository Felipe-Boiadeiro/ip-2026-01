package main

import "fmt"

var l1, l2, l3 float64

func main() {
	fmt.Println("Digite os lados do triângulo: ")
	fmt.Scan(&l1, &l2, &l3)

	if l1 < l2+l3 && l2 < l1+l3 && l3 < l1+l2 || l1 == l2 && l2 == l3 {
		fmt.Println("O triângulo está de acordo com a condição de existência")
		if l1 != l2 && l1 != l3 && l2 != l3 {
			fmt.Println("O triângulo é escaleno")
		} else if l1 == l2 && l2 == l3 {
			fmt.Println("O triângulo é equilátero")

		} else {
			fmt.Println("O triângulo é isósceles")
		}

	} else {
		fmt.Println("O triângulo não está de acordo com a condição de existência")

	}

}
