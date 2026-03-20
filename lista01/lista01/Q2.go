package main

import "fmt"

var n int

func main() {

	fmt.Println("Digite a quantidade de testes a ser realizado: ")

	fmt.Scan(&n)

	pes := make([]int, n)

	pop := make([]float64, n)

	ger := make([]float64, n)

	arq := make([]float64, n)

	cad := make([]float64, n)

	fmt.Println("Digite a quantidade de pessoas, seguido das porcentagens referentes as pessoas que compraram ingressos nas categorias popular, geral, arquibancada e cadeiras, respectivamente: ")

	for i := 0; i < n; i++ {

		fmt.Scan(&pes[i], &pop[i], &ger[i], &arq[i], &cad[i])

	}

	for i := 0; i < n; i++ {

		va := float64(pes[i]) * (pop[i] / 100)

		vb := float64(pes[i]) * (ger[i] / 100)

		vc := float64(pes[i]) * (arq[i] / 100)

		vd := float64(pes[i]) * (cad[i] / 100)

		vt := va + vb*5 + vc*10 + vd*20

		fmt.Printf("Renda do jogo N. %d= %.2f\n", i+1, vt)

	}

}
