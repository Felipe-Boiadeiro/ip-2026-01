package main

import "fmt"

func main() {
	notas := []int{}
	frequenciaAbsoluta := map[int]int{}
	n := 0

	for i := 1; i <= 15; i++ {
		fmt.Scan(&n)
		notas = append(notas, n)
		frequenciaAbsoluta[n]++
	}

	for nota := 0; nota <= 10; nota++ {
		qtd := frequenciaAbsoluta[nota]
		fmt.Printf("Frequência absoluta da nota %d: %d\n", nota, qtd)
		fmt.Printf("Frequência relativa da nota %d: %.2f\n", nota, float64(qtd)/15.0)
	}
}
