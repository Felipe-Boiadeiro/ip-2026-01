package main

import "fmt"

func main() {
	var n1, n2, n int

	fmt.Print("Informe n1: ")
	fmt.Scan(&n1)

	fmt.Print("Informe n2: ")
	fmt.Scan(&n2)

	fmt.Print("Informe a quantidade de termos (>3): ")
	fmt.Scan(&n)

	serie := []int{n1, n2}

	for i := 2; i < n; i++ {
		valor := 0
		if i%2 == 0 {
			valor = serie[i-1] - serie[i-2]
		} else {
			valor = serie[i-1] + serie[i-2]
		}

		serie = append(serie, valor)
	}

	fmt.Println(serie)
}
