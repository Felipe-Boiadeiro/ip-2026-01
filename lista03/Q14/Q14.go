package main

import "fmt"

func main() {
	n1 := 0
	n2 := 0
	lista := []int{}

	fmt.Scan(&n1)
	fmt.Scan(&n2)
	for i := n1; i <= n2; i++ {
		lista = append(lista, i)
	}

	for i := range lista {
		divisores := 0

		for y := 1; y <= lista[i]; y++ {
			if lista[i]%y == 0 {
				divisores++
			}
			if divisores > 2 {
				break
			}
		}

		if divisores == 2 {
			fmt.Println(lista[i], "é primo")

		}
	}
}
